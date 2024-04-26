package work

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
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
	if w.data.Options.Duration.IsDurationActive {
		dur := time.Duration(w.data.Options.Duration.Hours + w.data.Options.Duration.Minutes + w.data.Options.Duration.Seconds)
		w.ctx, w.ctxCancel = context.WithTimeout(ctx, dur)
	} else {
		w.ctx, w.ctxCancel = context.WithCancel(ctx)
	}

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
			w.stat.Duration = w.stat.EndedAt.Sub(w.stat.StartedAt)
			w.ctxCancel()
			w.logger.Info("work finished")
		}()
		for range w.data.Options.NumberOfRequests {
			req, err := w.generateRequest(w.data.Request)
			if err != nil {
				w.err = fmt.Errorf("generate request: %w", err)
				break
			}

			resp, err := w.client.Do(req)
			if err != nil {
				w.err = fmt.Errorf("client do: %w", err)
				break
			}

			err = resp.Body.Close()
			if err != nil {
				w.err = fmt.Errorf("resp body close: %w", err)
				break
			}

			w.addResponse(resp.StatusCode)
		}
	}()
	return nil
}

func (w *Work) addResponse(statusCode int) {
	w.mut.Lock()
	defer w.mut.Unlock()
	w.stat.Completed++
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
