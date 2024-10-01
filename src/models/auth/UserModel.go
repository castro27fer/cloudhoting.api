package auth

import (
	"time"

	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	Name           string          `gorm:"column:name;type:varchar(100);not null"`
	LastName       string          `gorm:"column:lastName;type:varchar(100);not null"`
	NumberPhone    string          `gorm:"column:numberPhone;type:varchar(30);null"`
	Address        string          `gorm:"column:address;type:varchar(400);null"`
	CountryID      string          `gorm:"column:countryId"`
	City           string          `gorm:"column:City;type:varchar(100);null"`
	Province       string          `gorm:"column:province;type:varchar(100);null"`
	CompanyName    string          `gorm:"column:companyName;type:varchar(100);null"`
	PostalCode     string          `gorm:"column:postalCode;type:varchar(10);null"`
	Password       string          `gorm:"column:password;type:varchar(100);null"`
	Email          string          `gorm:"column:email;type:varchar(100);null"`
	Language       string          `gorm:"column:language;type:varchar(2);default:en;not null"`
	Token          string          `gorm:"column:token;type:varchar(300);null"`
	Confirmed      bool            `gorm:"column:confirmed;type:boolean;default:true"`
	Logins         []LoginModel    `gorm:"foreignKey:UserID;references:ID"`
	ResetPasswords []ResetPassword `gorm:"foreignKey:UserID;references:ID"`
	ProfileTypeId  uint            `gorm:"column:ProfileTypeId"`
	Country        Country
	ProfileType    ProfileType
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}

func (UserModel) TableName() string {
	return "auth.user"
}
