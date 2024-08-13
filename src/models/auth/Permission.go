package auth

import (
	"time"

	"gorm.io/gorm"
)

type Permission struct {
	ID          uint           `gorm:"primaryKey"`
	Description string         `gorm:"column:description;type:varchar(1024);not null"`
	KeySecret   string         `gorm:"column:keySecret;type:varchar(1024);not null"`
	Accounts    []*UserModel   `gorm:"many2many:rel_accounts_permissions;"`
	ProfileType []*ProfileType `gorm:"many2many:rel_profileType_permissions;"`
	Active      bool           `gorm:"column:active;type:boolean;not null;default:true"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (Permission) TableName() string {
	return "public.permissions"
}
