package models

import (
	"time"

	db "github.com/ebarquero85/link-backend/src/database"
)

type TitleModel struct {
	TitleId   int       `gorm:"column:title_id;type:integer;primaryKey"`
	CreatedAt time.Time `gorm:"column:created_at;not null"`
	Hash      string    `gorm:"column:hash;type:char(32);not null"`
	Title     string    `gorm:"column:title;type:varchar(512);not null"`
}

func (TitleModel) TableName() string {
	return "public.titles"
}

func (Title *TitleModel) Create() error {

	if err := db.Databases.DBPostgresql.Instance.Create(Title).Error; err != nil {
		return err
	}

	return nil

}
