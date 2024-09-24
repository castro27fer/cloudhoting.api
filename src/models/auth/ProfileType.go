package auth

import (
	"time"

	db "github.com/ebarquero85/link-backend/src/database"
	models "github.com/ebarquero85/link-backend/src/models/auth"
	"gorm.io/gorm"
)

type ProfileType struct {
	ID          uint          `gorm:"primaryKey"`
	Name        string        `gorm:"column:name;type:varchar(150); not null"`
	Description string        `gorm:"column:description;type:varchar(500); null"`
	Active      bool          `gorm:"column:active;type:boolean;not null;default:true"`
	Permissions []*Permission `gorm:"many2many:rel_profileType_permissions;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (ProfileType) TableName() string {
	return "public.profileTypes"
}

func (r *ProfileType) AddPermission(p Permission) {

	rol := models.ProfileType{ID: r.ID}
	db.Databases.DBPostgresql.Instance.First(&rol)

	return rol

}
