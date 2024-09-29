package port

import (
	"context"

	"github.com/ksckaan1/apiredator/internal/domain/core/models"
)

type AppService interface {
	Startup(ctx context.Context)
	SetCurrentRequest(data models.Data) error
	GetCurrentRequest() (*models.Data, error)
	StartCurrentRequest() error
	ResetCurrentRequest() error
	SelectFiles(isMultiple bool) []string
	GetStats() (*models.Stat, error)
	AddToBookmark(title string, tags []string) error
	GetAllBookmarks(searchTerm, tag string, limit, offset int) (*models.BookmarkList, error)
	GetAllTags() ([]string, error)
	DeleteBookmarks(ids []string) error
}

type Repository interface {
	CreateBookmark(ctx context.Context, d *models.Bookmark) (string, error)
	GetAllBookmarks(ctx context.Context, searchTerm, tag string, limit, offset int) (*models.BookmarkList, error)
	GetAllTags(ctx context.Context) ([]string, error)
	DeleteBookmark(ctx context.Context, id string) error
}
