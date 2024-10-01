package services

import (
	db "github.com/ebarquero85/link-backend/src/database"
	models "github.com/ebarquero85/link-backend/src/models/auth"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	models.UserModel
}

func (user *User) FullName() string {
	return user.Name + " " + user.LastName
}

func (user *User) Create() error {

	//crypt password the of user
	password, err2 := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err2 != nil {
		return err2
	}

	user.Password = string(password)

	if err := db.Databases.DBPostgresql.Instance.Create(user).Error; err != nil {
		return err
	}

	return nil

}

func (user *User) ActiveAccount() {}

func (user *User) addLoginRecord() {}

func (user *User) LogOut() {}

func (user *User) LogIn(Email string, Password string) {}

func (user *User) ResetPassword(Email string, Token string, NewPassword string) {}
