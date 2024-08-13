package auth

import (
	"time"

	db "github.com/ebarquero85/link-backend/src/database"
	"gorm.io/gorm"
)

type UserModel struct {
	ID          uint           `gorm:"primaryKey"`
	Name        string         `gorm:"column:name;type:varchar(100);not null"`
	LastName    string         `gorm:"column:lastName;type:varchar(100);not null"`
	NumberPhone string         `gorm:"column:numberPhone;type:varchar(30);null"`
	Address     string         `gorm:"column:address;type:varchar(400);null"`
	CountryId   uint           `gorm:"column:countryId;type:integer;null"`
	City        string         `gorm:"column:City;type:varchar(100);null"`
	Province    string         `gorm:"column:province;type:varchar(100);null"`
	CompanyName string         `gorm:"column:companyName;type:varchar(100);null"`
	PostalCode  string         `gorm:"column:postalCode;type:varchar(10);null"`
	Accounts    []AccountModel `gorm:"foreignKey:UserId"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (UserModel) TableName() string {
	return "public.users"
}

func (User *UserModel) Create() error {

	if err := db.Databases.DBPostgresql.Instance.Create(User).Error; err != nil {
		return err
	}

	return nil

}

func (user *UserModel) Account_create(account *AccountModel) error {

	db := db.Databases.DBPostgresql.Instance

	profile := ProfileType{Name: "customer"}
	db.First(&profile)

	account.UserId = user.ID
	account.Permissions = profile.Permissions

	if err := db.Create(account).Error; err != nil {
		return err
	}

	return nil
}
