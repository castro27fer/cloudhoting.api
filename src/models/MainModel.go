package models

import (
	"time"

	db "github.com/ebarquero85/link-backend/src/database"
	"gorm.io/gorm"
)

type MainModel struct {
	Id           int            `gorm:"column:id;type:integer;primaryKey"`
	CreatedAt    time.Time      `gorm:"column:created_at;not null"`
	UserId       int            `gorm:"column:user_id;type:integer;not null"`
	CollectionId int            `gorm:"column:collection_id;type:integer;not null"`
	BookmarkId   int            `gorm:"column:bookmark_id;type:integer;not null"`
	TitleId      int            `gorm:"column:title_id;type:integer;not null"`
	IconId       int            `gorm:"column:icon_id;type:integer" sql:"default: null"`
	NoteId       int            `gorm:"column:note_id;type:integer" sql:"default: null"` //sql.NullInt32
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp;null" json:"-"`
}

func (MainModel) TableName() string {
	return "public.main"
}

func (Main *MainModel) Create() error {

	if err := db.Databases.DBPostgresql.Instance.Create(Main).Error; err != nil {
		return err
	}

	return nil

}

func (Main *MainModel) Update(id int, user_id int) error {

	if err := db.Databases.DBPostgresql.Instance.Where("id = ? AND user_id = ?", id, user_id).Updates(&Main); err != nil {
		return err.Error
	}

	return nil

}
