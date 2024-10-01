package auth

import (
	"time"

	db "github.com/ebarquero85/link-backend/src/database"
	"gorm.io/gorm"
)

type LoginModel struct {
	gorm.Model
	UserID    uint   `gorm:"column:userId"`
	Token     string `gorm:"column:token;type:varchar(1024);not null"`
	IP        string `gorm:"column:ip;type:varchar(30);null"`
	Active    bool   `gorm:"column:active;type:boolean; not null; default:true"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	User      UserModel
}

func (LoginModel) TableName() string {
	return "auth.logins"
}

func (Login *LoginModel) Create() error {

	if err := db.Databases.DBPostgresql.Instance.Create(Login).Error; err != nil {
		return err
	}

	return nil

}
