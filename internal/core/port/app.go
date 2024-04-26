package port

import (
	"context"

	"github.com/ksckaan1/apiredator/internal/core/domain"
)

type AppService interface {
	Startup(ctx context.Context)
	SendRequest(data domain.Data) error
	SelectFiles(isMultiple bool) []string
	GetStats() (*domain.Stat, error)
}
