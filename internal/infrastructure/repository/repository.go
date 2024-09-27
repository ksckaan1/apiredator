package repository

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/ksckaan1/apiredator/internal/domain/core/models"
	"github.com/ksckaan1/apiredator/internal/domain/core/port"
	"github.com/samber/lo"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

var _ port.Repository = (*Repository)(nil)

func NewRepository(db *gorm.DB) (*Repository, error) {
	return &Repository{
		db: db,
	}, nil
}

func (r *Repository) Migrate(ctx context.Context) error {
	err := r.db.WithContext(ctx).AutoMigrate(&Bookmark{})
	if err != nil {
		return fmt.Errorf("db: auto migrate: %w", err)
	}
	return nil
}

func (r *Repository) CreateBookmark(ctx context.Context, d *models.Bookmark) (string, error) {
	id := uuid.NewString()
	err := r.db.WithContext(ctx).
		Create(&Bookmark{
			ID:        id,
			CreatedAt: time.Now(),
			Title:     d.Title,
			Request:   datatypes.NewJSONType(d.Request),
			Options:   datatypes.NewJSONType(d.Options),
			Stat:      datatypes.NewJSONType(d.Stat),
			Tags:      d.Tags,
		}).Error
	if err != nil {
		return "", fmt.Errorf("db: create: %w", err)
	}
	return id, nil
}

func (r *Repository) GetAllBookmarks(ctx context.Context, searchTerm, tag string, limit int, offset int) (*models.BookmarkList, error) {
	var (
		bookmarks []Bookmark
		count     int64
	)

	// SELECT bookmarks.id, bookmarks.*
	// FROM bookmarks, json_each(tags)
	// WHERE title LIKE "%%"
	// AND EXISTS (SELECT 1 FROM json_each(tags) WHERE json_each.value = 'custom api') GROUP By bookmarks.id;
	tx := r.db.WithContext(ctx).
		Select("bookmarks.id, bookmarks.*").
		Table("bookmarks, json_each(tags)")

	if strings.TrimSpace(searchTerm) != "" {
		tx = tx.Where("title LIKE ?", "%"+searchTerm+"%")
	}

	if strings.TrimSpace(tag) != "" {
		tx = tx.Where("EXISTS (SELECT 1 FROM json_each(tags) WHERE json_each.value = ?)", tag)
	}

	err := tx.
		Group("bookmarks.id").
		Count(&count).
		Offset(offset).
		Limit(limit).
		Find(&bookmarks).Error
	if err != nil {
		return nil, fmt.Errorf("db: find: %w", err)
	}

	return &models.BookmarkList{
		Bookmarks: lo.Map(bookmarks, func(b Bookmark, _ int) models.Bookmark {
			return models.Bookmark{
				ID:       b.ID,
				CreateAt: b.CreatedAt,
				Title:    b.Title,
				Request:  b.Request.Data(),
				Options:  b.Options.Data(),
				Stat:     b.Stat.Data(),
				Tags:     b.Tags,
			}
		}),
		Limit:  limit,
		Offset: offset,
		Count:  count,
	}, nil
}

func (r *Repository) GetAllTags(ctx context.Context) ([]string, error) {
	var results []struct {
		Tag string
	}
	err := r.db.WithContext(ctx).
		Select("json_each.value as tag").
		Table("bookmarks, json_each(tags)").
		Group("tag").
		Find(&results).
		Error
	if err != nil {
		return nil, fmt.Errorf("db: find: %w", err)
	}
	return lo.Map(results, func(t struct{ Tag string }, _ int) string {
		return t.Tag
	}), nil
}
