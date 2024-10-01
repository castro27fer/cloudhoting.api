package auth

import (
	"time"

	"gorm.io/gorm"
)

type ResetPassword struct {
	gorm.Model
	Token       string `gorm:"column:token;type:varchar(1024);not null"`
	PasswordOld string `gorm:"column:passwordOld;type:varchar(1024);not null"`
	UserID      uint   `gorm:"column:passwordOld"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	User        UserModel
}

func (ResetPassword) TableName() string {
	return "auth.resetPasswords"
}
