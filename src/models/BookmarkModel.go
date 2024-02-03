package models

import (
	"time"

	db "github.com/ebarquero85/link-backend/src/database"
)

type BookmarkModel struct {
	BookmarkId int       `gorm:"column:bookmark_id;type:integer;primaryKey"`
	CreatedAt  time.Time `gorm:"column:created_at;not null"`
	Hash       string    `gorm:"column:hash;type:char(32);not null"`
	Url        string    `gorm:"column:url;type:varchar(2048);not null"`
}

func (BookmarkModel) TableName() string {
	return "public.bookmarks"
}

func (Bookmark *BookmarkModel) Create() error {

	if err := db.Databases.DBPostgresql.Instance.Create(Bookmark).Error; err != nil {
		return err
	}

	return nil

}
