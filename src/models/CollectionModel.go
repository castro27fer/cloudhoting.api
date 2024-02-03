package models

import (
	"time"

	db "github.com/ebarquero85/link-backend/src/database"
	"gorm.io/gorm"
)

type CollectionModel struct {
	CollectionId int            `gorm:"column:collection_id;type:integer;primaryKey"`
	CreatedAt    time.Time      `gorm:"column:created_at;not null" json:"-"`
	UserId       int            `gorm:"column:user_id;type:integer;not null" json:"-"`
	Name         string         `gorm:"column:name;type:varchar(100);not null"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp;null" json:"-"`
}

func (CollectionModel) TableName() string {
	return "public.collections"
}

func (Collection *CollectionModel) Create() error {

	if err := db.Databases.DBPostgresql.Instance.Create(Collection).Error; err != nil {
		return err
	}

	return nil

}
