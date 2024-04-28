package port

import (
	"context"

	"github.com/ksckaan1/apiredator/internal/core/domain"
)

type Work interface {
	Start(ctx context.Context) error
	Stop()
	IsActive() bool
	Wait()
	GetStats() domain.Stat
	GetDetails() *domain.Data
}
