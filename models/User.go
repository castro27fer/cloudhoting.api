package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Id          uint   `gorm:"primaryKey;autoIncrement"`
	Name        string `gorm:"not null"`
	LastName    string `gorm:"not null"`
	numberPhone string
	Email       string `gorm:"not null;unique_index"`
	Active      bool
}
