package routes

import "net/http"

func Holamundo(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Conquistar el mundo con Go"))
}
