package port

import (
	"context"

	"github.com/ksckaan1/apiredator/internal/domain/core/models"
)

type Work interface {
	Start(ctx context.Context) error
	Stop()
	IsActive() bool
	Wait()
	GetStats() models.Stat
	GetDetails() *models.Data
}
