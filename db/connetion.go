package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dsn = "host=localhost user=postgres password=postgresql dbname=cloudHostingDB port=5432 sslmode=disable TimeZone=Asia/Shanghai"
var DB *gorm.DB

func ConnectionDB() {

	var errorConnection error
	DB, errorConnection = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if errorConnection != nil {
		log.Fatal(errorConnection)
	} else {
		log.Println("BD connected...")
	}
}

func CreateSchema(schema string) {

	err := DB.Exec("CREATE SCHEMA IF NOT EXISTS " + schema).Error
	if err != nil {
		panic("Error al crear el esquema: " + err.Error())
	}
}
