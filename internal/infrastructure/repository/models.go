package repository

import (
	"time"

	"github.com/ksckaan1/apiredator/internal/domain/core/models"
	"gorm.io/datatypes"
)

type Bookmark struct {
	ID        string                             `json:"id"`
	CreatedAt time.Time                          `json:"created_at"`
	Title     string                             `json:"title"`
	Request   datatypes.JSONType[models.Request] `json:"request"`
	Options   datatypes.JSONType[models.Options] `json:"options"`
	Stat      datatypes.JSONType[models.Stat]    `json:"stat"`
	Tags      datatypes.JSONSlice[string]        `json:"tags"`
}
