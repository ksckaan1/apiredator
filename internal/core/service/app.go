package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"

	"github.com/ksckaan1/apiredator/internal/core/domain"
	"github.com/ksckaan1/apiredator/internal/core/port"
	"github.com/ksckaan1/apiredator/pkg/work"
)

var _ port.AppService = (*AppService)(nil)

type AppService struct {
	ctx         context.Context
	currentWork port.Work
	logger      port.Logger
}

func NewAppService(lg port.Logger) *AppService {
	return &AppService{
		logger: lg,
	}
}

func (a *AppService) Startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *AppService) SetCurrentRequest(data domain.Data) error {
	var reqTimeout time.Duration

	if data.Options.RequestTimeout != "" {
		v, err := time.ParseDuration(data.Options.RequestTimeout)
		if err != nil {
			return fmt.Errorf("time: parse duration: %w", err)
		}
		reqTimeout = v
	}

	a.currentWork = work.New(a.logger, &data, reqTimeout)
	a.logger.Info("set current request",
		"data", data,
	)
	return nil
}

func (a *AppService) GetCurrentRequest() (*domain.Data, error) {
	if a.currentWork == nil {
		a.logger.Error("current work not found")
		return nil, nil
	}
	data := a.currentWork.GetDetails()
	a.logger.Info("get current work",
		"data", data,
	)
	return data, nil
}

func (a *AppService) StartCurrentRequest() error {
	if a.currentWork == nil {
		a.logger.Error("current request not found")
		return errors.New("current work is not set")
	}

	err := a.currentWork.Start(context.Background())
	if err != nil {
		a.logger.Error("current request can not started",
			"error", err,
		)
		return fmt.Errorf("start work: %w", err)
	}
	a.logger.Info("current request started")

	return nil
}

func (a *AppService) ResetCurrentRequest() error {
	a.currentWork = nil
	return nil
}

func (a *AppService) SelectFiles(isMultiple bool) []string {
	if isMultiple {
		files, err := runtime.OpenMultipleFilesDialog(a.ctx, runtime.OpenDialogOptions{
			ShowHiddenFiles: true,
		})
		if err != nil {
			return nil
		}
		return files
	}

	file, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		ShowHiddenFiles: true,
	})
	if err != nil {
		return nil
	}
	return []string{file}
}

func (a *AppService) GetStats() (*domain.Stat, error) {
	if a.currentWork == nil {
		return nil, errors.New("no work started")
	}
	stats := a.currentWork.GetStats()
	return &stats, nil
}

func (a *AppService) StopWork() error {
	if a.currentWork == nil {
		return errors.New("current work not found")
	}

	if !a.currentWork.IsActive() {
		return errors.New("already stopped")
	}

	a.currentWork.Stop()
	a.logger.Info("current work stopped")

	return nil
}

func (a *AppService) IsWorkActive() bool {
	return !(a.currentWork == nil || !a.currentWork.IsActive())
}

func (a *AppService) WaitWork() error {
	if a.currentWork == nil {
		a.logger.Error("current work not found")
		return errors.New("current work not found")
	}

	a.logger.Info("waiting for current work")
	a.currentWork.Wait()
	a.logger.Info("finished current work")
	return nil
}
