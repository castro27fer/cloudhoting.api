package auth

import (
	"time"

	"gorm.io/gorm"
)

type Country struct {
	gorm.Model
	ID        string `gorm:"primaryKey"`
	Name      string `gorm:"column:name;type:varchar(250);not null"`
	Active    bool   `gorm:"column:active;type:boolean;default:true"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (Country) TableName() string {
	return "auth.country"
}
