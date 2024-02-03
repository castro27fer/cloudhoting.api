package models

import (
	"time"

	db "github.com/ebarquero85/link-backend/src/database"
)

type LoginModel struct {
	LoginId   int       `gorm:"column:login_id;type:integer;primaryKey"`
	CreatedAt time.Time `gorm:"column:created_at;not null"`
	UserId    int       `gorm:"column:user_id;type:integer;not null"`
}

func (LoginModel) TableName() string {
	return "public.logins"
}

func (Login *LoginModel) Create() error {

	if err := db.Databases.DBPostgresql.Instance.Create(Login).Error; err != nil {
		return err
	}

	return nil

}
