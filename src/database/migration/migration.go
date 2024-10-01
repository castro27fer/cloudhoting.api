package migration

import (
	"github.com/ebarquero85/link-backend/src/database"
	auth "github.com/ebarquero85/link-backend/src/models/auth"
)

func Init(BD2 database.BD, load_data bool) error {

	// create schema if not exists
	BD2.DBPostgresql.Instance.Exec("CREATE SCHEMA IF NOT EXISTS auth")

	err := BD2.DBPostgresql.Instance.Migrator().AutoMigrate(
		&auth.ProfileType{},
		&auth.UserModel{},
		&auth.CodeVerifyModel{},
		&auth.LoginModel{},
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

	countries := []auth.Country{
		{Name: "Nicaragua", ID: "NI"},
		{Name: "Estados Unidos", ID: "EU"},
		{Name: "Rusia", ID: "RU"},
		{Name: "China", ID: "CH"},
		{Name: "Argentina", ID: "AR"},
	}

	if err := create(profile_types); err != nil {
		return err
	}

	if err := create(permissions); err != nil {
		return err
	}

	if err := create(countries); err != nil {
		return err
	}

	return nil
}
