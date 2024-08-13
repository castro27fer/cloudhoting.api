package migration

import (
	"github.com/ebarquero85/link-backend/src/database"
	auth "github.com/ebarquero85/link-backend/src/models/auth"
)

func Init(BD2 database.BD, load_data bool) error {

	err := BD2.DBPostgresql.Instance.Migrator().AutoMigrate(
		&auth.UserModel{},
		&auth.AccountModel{},
		&auth.CodeVerifyModel{},
		&auth.LoginModel{},
		&auth.ProfileType{},
	)

	if err != nil {
		return err
	}

	if load_data {
		if err = load_data_base(); err != nil {
			return err
		}
	}

	return nil
}

func create(i interface{}) error {

	result := database.Databases.DBPostgresql.Instance.Create(i)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func load_data_base() error {

	profile_types := []*auth.ProfileType{
		{Name: "customer"},
		{Name: "employee"},
	}

	permissions := []*auth.Permission{
		{Description: "Access profile", KeySecret: "ACCESS_PROFILE"},
		{Description: "Edict Profile", KeySecret: "EDICT_PROFILE"},
		{Description: "Change Password", KeySecret: "CHANGE_PASSWORD"},
	}

	if err := create(profile_types); err != nil {
		return err
	}

	if err := create(permissions); err != nil {
		return err
	}

	return nil
}
