package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/wailsapp/wails/v2/pkg/runtime"

	"github.com/ksckaan1/apiredator/internal/core/domain"
	"github.com/ksckaan1/apiredator/internal/core/port"
	"github.com/ksckaan1/apiredator/pkg/work"
)

var _ port.AppService = (*AppService)(nil)

// App struct
type AppService struct {
	ctx         context.Context
	currentWork port.Work
	logger      port.Logger
}

// NewApp creates a new App application struct
func NewAppService(lg port.Logger) *AppService {
	return &AppService{
		logger: lg,
	}
}

func (a *AppService) Startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *AppService) SendRequest(data domain.Data) error {
	a.currentWork = work.New(a.logger, &data)

	err := a.currentWork.Start(context.Background())
	if err != nil {
		return fmt.Errorf("start work: %w", err)
	}

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
	if !a.currentWork.IsActive() {
		return errors.New("already stopped")
	}
	a.currentWork.Stop()
	return nil
}

func (a *AppService) IsWorkActive() bool {
	return !(a.currentWork == nil || !a.currentWork.IsActive())
}

func (a *AppService) WaitWork() error {
	if a.currentWork == nil {
		return errors.New("no work started")
	}
	a.currentWork.Wait()
	return nil
}
