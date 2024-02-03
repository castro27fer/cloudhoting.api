package models

import (
	"time"

	db "github.com/ebarquero85/link-backend/src/database"
	"gorm.io/gorm"
)

type CategoryModel struct {
	CategoryId   int            `gorm:"column:category_id;type:integer;primaryKey"`
	CreatedAt    time.Time      `gorm:"column:created_at;not null" json:"-"`
	UserId       int            `gorm:"column:user_id;type:integer;not null" json:"-"`
	CollectionId int            `gorm:"column:collection_id;type:integer;not null" json:"-"`
	Name         string         `gorm:"column:name;type:varchar(100);not null"`
	Color        string         `gorm:"column:color;type:char(3);not null"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp;null" json:"-"`
}

func (CategoryModel) TableName() string {
	return "public.categories"
}

func (Category *CategoryModel) Create() error {

	if err := db.Databases.DBPostgresql.Instance.Create(Category).Error; err != nil {
		return err
	}

	return nil

}
