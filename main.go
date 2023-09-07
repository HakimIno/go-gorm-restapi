package main

import (
	"net/http"

	"github.com/App/go-gorm-restapi/db"
	"github.com/App/go-gorm-restapi/models"
	"github.com/App/go-gorm-restapi/routes"
	"github.com/gorilla/mux"
)

func main() {

	db.DBConnection()

	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})

	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)

	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/user/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/user", routes.PostUserHandler).Methods("POST")
	r.HandleFunc("/user/{id}", routes.DeleteUsersHandler).Methods("DELETE")

	http.ListenAndServe(":3000", r)
}
