package services

import (
	"os"

	db "github.com/ebarquero85/link-backend/src/database"
	"github.com/ebarquero85/link-backend/src/models/auth"
)

type Profiles struct {
	Rol_default string
}

var ProfilesServices Profiles = Profiles{}

func (profiles *Profiles) Start() {
	profiles.Rol_default = os.Getenv("ROL_DEFAULT")
}

func (Profiles *Profiles) Get_rol_default() auth.ProfileType {

	profile := auth.ProfileType{Name: Profiles.Rol_default}
	db.Databases.DBPostgresql.Instance.First(&profile)
	return profile
}

func (Profiles *Profiles) Rols() {}

func (Profiles *Profiles) CreateRol(rol auth.ProfileType) {}

func (Profiles *Profiles) Permissions() {}

func (Profiles *Profiles) CreatePermission(permission auth.Permission) {}

func (Profiles *Profiles) AssociateRolPermissions(rol auth.ProfileType, permission auth.Permission) {}

func (Profiles *Profiles) DuplicateRol(rol auth.ProfileType, rename string) {}
