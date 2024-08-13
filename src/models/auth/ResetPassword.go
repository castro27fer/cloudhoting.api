package auth

import (
	"time"

	"gorm.io/gorm"
)

type ResetPassword struct {
	ID          uint   `gorm:"primaryKey"`
	Token       string `gorm:"column:token;type:varchar(1024);not null"`
	PasswordOld string `gorm:"column:passwordOld;type:varchar(1024);not null"`
	AccountId   uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (ResetPassword) TableName() string {
	return "public.resetPasswords"
}
