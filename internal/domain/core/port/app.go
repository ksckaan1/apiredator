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
	GetBookmarkByID(id string) (*models.Bookmark, error)
	GetAllBookmarks(searchTerm, tag string, limit, offset int) (*models.BookmarkList, error)
	GetAllTags() ([]string, error)
	UpdateBookmark(d *models.UpdateBookmark) error
	DeleteBookmarks(ids []string) error
}

type Repository interface {
	CreateBookmark(ctx context.Context, d *models.Bookmark) (string, error)
	GetBookmarkByID(ctx context.Context, id string) (*models.Bookmark, error)
	GetAllBookmarks(ctx context.Context, searchTerm, tag string, limit, offset int) (*models.BookmarkList, error)
	GetAllTags(ctx context.Context) ([]string, error)
	UpdateBookmark(ctx context.Context, d *models.UpdateBookmark) error
	DeleteBookmark(ctx context.Context, id string) error
}
