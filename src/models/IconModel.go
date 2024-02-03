package models

import (
	"time"

	db "github.com/ebarquero85/link-backend/src/database"
)

type IconModel struct {
	IconId    int       `gorm:"column:icon_id;type:integer;primaryKey"`
	CreatedAt time.Time `gorm:"column:created_at;not null"`
	Hash      string    `gorm:"column:hash;type:char(32);not null"`
	Icon      string    `gorm:"column:icon;type:text;not null"`
}

func (IconModel) TableName() string {
	return "public.icons"
}

func (Icon *IconModel) Create() error {

	if err := db.Databases.DBPostgresql.Instance.Create(Icon).Error; err != nil {
		return err
	}

	return nil

}
