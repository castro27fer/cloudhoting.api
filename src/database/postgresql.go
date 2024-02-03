package database

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Postgres struct {
	Instance *gorm.DB
}

func (pg *Postgres) Connect() *Postgres {

	var (
		dbhost     = os.Getenv("DB_HOST")
		dbname     = os.Getenv("DB_NAME")
		dbuser     = os.Getenv("DB_USER")
		dbpassword = os.Getenv("DB_PASSWORD")
		dbport     = os.Getenv("DB_PORT")
		dns        = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Managua", dbhost, dbuser, dbpassword, dbname, dbport)
	)

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{
		SkipDefaultTransaction: true,
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",    // table name prefix, table for `User` would be `t_users`
			SingularTable: false, // use singular table name, table for `User` would be `user` with this option enabled
			NoLowerCase:   false, // skip the snake_casing of names
			//NameReplacer:  strings.NewReplacer("CID", "Cid"), // use name replacer to change struct/field name before convert it to db name
		},
	})

	if err != nil {
		panic(err.Error())
	}

	pg.Instance = db

	return pg

}

// func (pg *Postgres) Delete(model *interface{}, id int) {

// 	pg.Instance.Delete(model, id)

// }
