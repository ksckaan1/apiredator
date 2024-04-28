package port

import (
	"context"

	"github.com/ksckaan1/apiredator/internal/core/domain"
)

type AppService interface {
	Startup(ctx context.Context)
	SetCurrentRequest(data domain.Data) error
	GetCurrentRequest() (*domain.Data, error)
	StartCurrentRequest() error
	ResetCurrentRequest() error
	SelectFiles(isMultiple bool) []string
	GetStats() (*domain.Stat, error)
}
