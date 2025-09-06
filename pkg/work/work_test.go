package work

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/ksckaan1/logger"
	"github.com/stretchr/testify/require"

	"github.com/ksckaan1/apiredator/internal/domain/core/models"
)

func TestStartWork(t *testing.T) {
	lg, err := logger.New(logger.DefaultConfig())
	require.NoError(t, err)

	w := New(
		lg,
		&models.Data{
			Request: models.Request{
				Method: http.MethodPost,
				URL:    "https://jsonplaceholder.typicode.com/posts",
				Header: []models.KeyValueData{},
				Body: models.Body{
					Type:     "raw",
					RawValue: `{"title":"foo","body":"bar","userId":1}`,
				},
			},
			Options: models.Options{
				NumberOfRequests: 100,
			},
		},
		0,
	)

	ctx := context.Background()

	err = w.Start(ctx)
	require.NoError(t, err)
	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()

		for range ticker.C {
			stats := w.GetStats()
			fmt.Println(stats)
		}
	}()

	w.Wait()
}
