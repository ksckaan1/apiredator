package work

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"slices"
	"sync"
	"time"

	"github.com/ksckaan1/apiredator/internal/core/domain"
	"github.com/ksckaan1/apiredator/internal/core/port"
)

var _ port.Work = (*Work)(nil)

type Work struct {
	ctx       context.Context
	ctxCancel context.CancelFunc
	stat      *domain.Stat
	data      *domain.Data
	mut       *sync.RWMutex
	client    *http.Client
	err       error
	isActive  bool
	logger    port.Logger
}

func New(lg port.Logger, data *domain.Data) *Work {
	return &Work{
		data:   data,
		mut:    &sync.RWMutex{},
		client: &http.Client{},
		logger: lg,
		stat: &domain.Stat{
			StatusCodes: map[int]uint64{},
		},
	}
}

func (w *Work) Start(ctx context.Context) error {
	if w.data.Options.TestType == domain.TTDuration {
		seconds := time.Second * time.Duration(w.data.Options.Duration.Seconds)
		minutes := time.Minute * time.Duration(w.data.Options.Duration.Minutes)
		hours := time.Hour * time.Duration(w.data.Options.Duration.Hours)
		dur := hours + minutes + seconds
		w.ctx, w.ctxCancel = context.WithTimeout(ctx, dur)
		fmt.Println("dur", dur.String())
	} else {
		w.ctx, w.ctxCancel = context.WithCancel(ctx)
	}

	w.reset()

	err := w.run()
	if err != nil {
		return fmt.Errorf("run: %w", err)
	}

	w.logger.Info("work started")

	return nil
}

func (w *Work) Stop() {
	w.ctxCancel()
	w.logger.Info("work stopped")
}

func (w *Work) IsActive() bool {
	return w.isActive
}

func (w *Work) Wait() {
	<-w.ctx.Done()
}

func (w *Work) GetStats() domain.Stat {
	return *w.stat
}

func (w *Work) run() error {
	go func() {
		w.stat.StartedAt = time.Now()
		w.isActive = true
		defer func() {
			w.stat.EndedAt = time.Now()
			w.isActive = false
			w.stat.PassedDuration = w.stat.EndedAt.Sub(w.stat.StartedAt)
			w.ctxCancel()
			w.logger.Info("work finished")
		}()

		go w.checkPeriodically()

		var wg sync.WaitGroup

		for i := range w.data.Options.NumberOfClients {
			wg.Add(1)
			go func() {
				w.logger.Debug("client started",
					"num", i,
				)
				if w.data.Options.TestType == domain.TTDuration {
					for {
						err := w.makeRequest()
						if err != nil {
							break
						}
					}
				} else {
					for range w.data.Options.NumberOfRequests {
						err := w.makeRequest()
						if err != nil {
							break
						}
					}
				}
				w.logger.Debug("client finished",
					"num", i,
				)
				wg.Done()
			}()
		}
		wg.Wait()
	}()
	return nil
}

func (w *Work) makeRequest() error {
	req, err := w.generateRequest(w.data.Request)
	if err != nil {
		w.err = fmt.Errorf("generate request: %w", err)
		return w.err
	}

	resp, err := w.client.Do(req)
	if err != nil {
		w.err = fmt.Errorf("client do: %w", err)
		return w.err
	}

	err = resp.Body.Close()
	if err != nil {
		w.err = fmt.Errorf("resp body close: %w", err)
		return w.err
	}

	w.addResponse(resp.StatusCode)

	return nil
}

func (w *Work) addResponse(statusCode int) {
	w.mut.Lock()
	defer w.mut.Unlock()
	w.stat.SentCount++
	_, ok := w.stat.StatusCodes[statusCode]
	if !ok {
		w.stat.StatusCodes[statusCode] = 0
	}
	w.stat.StatusCodes[statusCode] += 1
}

func (w *Work) generateRequest(requestData domain.Request) (*http.Request, error) {
	body, err := w.generateBody(requestData.Body)
	if err != nil {
		return nil, fmt.Errorf("generate body: %w", err)
	}
	req, err := http.NewRequestWithContext(w.ctx, requestData.Method, requestData.URL, body)
	if err != nil {
		return nil, fmt.Errorf("new request with context: %w", err)
	}
	w.addHeaders(req, requestData.Header)

	return req, nil
}

func (w *Work) generateBody(b domain.Body) (io.Reader, error) {
	var body io.Reader
	switch b.Type {
	case "raw":
		body = bytes.NewBufferString(b.RawValue)
	case "formdata":

	}
	return body, nil
}

func (w *Work) generateFormData(fd []domain.FormData) (io.Reader, string, error) {
	body := &bytes.Buffer{}

	mw := multipart.NewWriter(body)
	defer mw.Close()

	for _, f := range fd {
		if !f.IsActive {
			continue
		}

		if f.RowType == "text" {
			err := mw.WriteField(f.Key, f.TextValue)
			if err != nil {
				return nil, "", fmt.Errorf("mw write field: %w", err)
			}
		}
	}

	return nil, mw.FormDataContentType(), nil
}

func (w *Work) addHeaders(req *http.Request, header []domain.KeyValueData) {
	for _, item := range header {
		if !item.IsActive {
			continue
		}
		req.Header.Add(item.Key, item.Value)
	}
}

func (w *Work) GetDetails() *domain.Data {
	return w.data
}

func (w *Work) reset() {
	w.stat.StartedAt = time.Time{}
	w.stat.EndedAt = time.Time{}
	w.stat.PassedDuration = 0
	w.stat.SentCount = 0
	w.stat.StatusCodes = map[int]uint64{}
	w.stat.RPS.List = []uint64{}
	w.stat.RPS.Latest = 0
	w.stat.RPS.Min = 0
	w.stat.RPS.Avg = 0
	w.stat.RPS.Max = 0
}

func (w *Work) checkPeriodically() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	var lastSentCount uint64
	for {
		select {
		case <-w.ctx.Done():
			return
		case <-ticker.C:
			w.mut.Lock()
			w.stat.RPS.Latest = w.stat.SentCount - lastSentCount
			lastSentCount = w.stat.SentCount
			w.stat.RPS.List = append(w.stat.RPS.List, w.stat.RPS.Latest)

			w.stat.RPS.Min = slices.Min(w.stat.RPS.List)
			w.stat.RPS.Max = slices.Max(w.stat.RPS.List)
			w.stat.RPS.Avg = getAvg(w.stat.RPS.List)

			w.stat.PassedDuration = time.Since(w.stat.StartedAt)
			w.mut.Unlock()
		}
	}
}

func getAvg(list []uint64) float64 {
	if len(list) == 0 {
		return 0
	}

	var (
		sum   uint64
		count uint64
	)

	for i := range list {
		sum += list[i]
		count++
	}

	return float64(sum) / float64(count)
}
