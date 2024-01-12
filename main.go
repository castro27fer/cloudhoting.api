package main

import (
	"net/http"

	"github.com/castro27fer/cloudhosting.api/db"
	"github.com/castro27fer/cloudhosting.api/models"
	"github.com/castro27fer/cloudhosting.api/routes"
	"github.com/gorilla/mux"
)

func main() {

	//Connection DB
	db.ConnectionDB()

	//create shema
	db.CreateSchema("auth")
	// Migrate
	db.DB.Table("auth.users").AutoMigrate(models.User{})

	//Routes
	router := mux.NewRouter()
	router.HandleFunc("/", routes.Holamundo)
	router.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	router.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	router.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	router.HandleFunc("/users/{id}", routes.PutUserHandler).Methods("PUT")
	router.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")

	http.ListenAndServe(":3000", router)
}
