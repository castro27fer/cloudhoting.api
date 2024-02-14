package models

import (
	"time"

	db "github.com/ebarquero85/link-backend/src/database"
	"gorm.io/gorm"
)

type AccountModel struct {
	ID        uint
	UserId    uint           `gorm:"column:userId;type:integer;not null"`
	Password  string         `gorm:"column:password;type:varchar(100);null"`
	Email     string         `gorm:"column:email;type:varchar(100);null"`
	Language  string         `gorm:"column:language;type:varchar(2);default:en;not null"`
	Token     string         `gorm:"column:token;type:varchar(300);null"`
	Confirmed bool           `gorm:"column:confirmed;type:boolean;default:true"`
	CreatedAt time.Time      `gorm:"column:created_at;type:timestamp;not null"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp;null"`
}

func (AccountModel) TableName() string {
	return "public.accounts"
}

func (Account *AccountModel) Create() error {

	if err := db.Databases.DBPostgresql.Instance.Create(Account).Error; err != nil {
		return err
	}

	return nil

}
