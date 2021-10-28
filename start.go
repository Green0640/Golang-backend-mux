package main

import (
	"backend/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/users", controllers.GetAllUsers).Methods("GET", "OPTION")
	r.HandleFunc("/user/{id}", controllers.GetUser).Methods("GET")
	r.HandleFunc("/user", controllers.AddUser).Methods("POST")
	r.HandleFunc("/user/{id}", controllers.UpdateUser).Methods("PUT")
	r.HandleFunc("/user/{id}", controllers.DeleteUser).Methods("DELETE")
	r.Use(mux.CORSMethodMiddleware(r))
	http.ListenAndServe(":8080", r)
}
