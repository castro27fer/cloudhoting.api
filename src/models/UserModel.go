package models

import (
	"time"

	db "github.com/ebarquero85/link-backend/src/database"
	"gorm.io/gorm"
)

type UserModel struct {
	UserId            int            `gorm:"column:user_id;type:integer;primaryKey"`
	CreatedAt         time.Time      `gorm:"column:created_at;type:timestamp;not null"`
	UserName          string         `gorm:"column:username;type:varchar(100);not null"`
	Password          string         `gorm:"column:password;type:varchar(100);not null"`
	Email             string         `gorm:"column:email;type:varchar(100);not null"`
	Language          string         `gorm:"column:language;type:varchar(2);default:en;not null"`
	CollectionDefault int            `gorm:"column:collection_default;type:integer;default:0"`
	Confirmed         bool           `gorm:"column:confirmed;type:boolean;default:true"`
	DeletedAt         gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp;null"`
}

func (UserModel) TableName() string {
	return "public.users"
}

func (User *UserModel) Create() error {

	// config.Lock.Lock()
	// defer config.Lock.Unlock()

	if err := db.Databases.DBPostgresql.Instance.Create(User).Error; err != nil {
		return err
	}

	return nil

}
