package work

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/ksckaan1/apiredator/internal/domain/core/models"
	"github.com/ksckaan1/apiredator/pkg/logger"
)

func TestStartWork(t *testing.T) {
	lg := logger.NewZerolog().
		WithWriter(os.Stdout)

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

	err := w.Start(ctx)
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
