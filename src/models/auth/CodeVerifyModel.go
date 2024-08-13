package auth

import (
	"time"

	db "github.com/ebarquero85/link-backend/src/database"
	"gorm.io/gorm"
)

type CodeVerifyModel struct {
	ID        uint   `gorm:"primaryKey"`
	Email     string `gorm:"column:email;type:varchar(100);not null"`
	Code      string `gorm:"column:code;type:varchar(6);not null"`
	Status    string `gorm:"column:status;type:varchar(100);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (CodeVerifyModel) TableName() string {
	return "public.codeVerify"
}

func (CodeVerify *CodeVerifyModel) Create() error {

	// config.Lock.Lock()
	// defer config.Lock.Unlock()

	if err := db.Databases.DBPostgresql.Instance.Create(CodeVerify).Error; err != nil {
		return err
	}

	return nil

}

func (CodeVerify *CodeVerifyModel) Update() error {
	if err := db.Databases.DBPostgresql.Instance.Where("id = ?", CodeVerify.ID).Updates(&CodeVerify); err != nil {
		return err.Error
	}

	return nil
}
