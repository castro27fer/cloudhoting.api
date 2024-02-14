package models

import (
	"time"

	db "github.com/ebarquero85/link-backend/src/database"
	"gorm.io/gorm"
)

type UserModel struct {
	ID                uint
	UserId            int            `gorm:"column:user_id;type:integer;primaryKey"`
	Name              string         `gorm:"column:name;type:varchar(100);not null"`
	LastName          string         `gorm:"column:lastName;type:varchar(100);not null"`
	NumberPhone       string         `gorm:"column:numberPhone;type:varchar(30);null"`
	Address           string         `gorm:"column:address;type:varchar(400);null"`
	CountryId         uint           `gorm:"column:countryId;type:integer;null"`
	City              string         `gorm:"column:City;type:varchar(100);null"`
	Province          string         `gorm:"column:province;type:varchar(100);null"`
	CompanyName       string         `gorm:"column:companyName;type:varchar(100);null"`
	PostalCode        string         `gorm:"column:postalCode;type:varchar(10);null"`
	UserName          string         `gorm:"column:username;type:varchar(100);null"`
	Password          string         `gorm:"column:password;type:varchar(100);null"`
	Email             string         `gorm:"column:email;type:varchar(100);null"`
	Language          string         `gorm:"column:language;type:varchar(2);default:en;not null"`
	CollectionDefault int            `gorm:"column:collection_default;type:integer;default:0"`
	Confirmed         bool           `gorm:"column:confirmed;type:boolean;default:true"`
	CreatedAt         time.Time      `gorm:"column:created_at;type:timestamp;not null"`
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
