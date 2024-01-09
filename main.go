package main

import (
	"fmt"
	"net/http"

	"github.com/castro27fer/cloudhosting.api/routes"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", routes.Holamundo)

	http.ListenAndServe(":3000", router)
	fmt.Println("Conquistar el mundo con Go")
}
