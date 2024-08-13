package auth

import (
	"time"

	"gorm.io/gorm"
)

type ProfileType struct {
	ID          uint          `gorm:"primaryKey"`
	Name        string        `gorm:"column:name;type:varchar(150); not null"`
	Active      bool          `gorm:"column:active;type:boolean;not null;default:true"`
	Permissions []*Permission `gorm:"many2many:rel_profileType_permissions;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (ProfileType) TableName() string {
	return "public.profileTypes"
}
