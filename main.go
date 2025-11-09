package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sindessis-software/ApiRestPuntoVenta/db"
	"github.com/sindessis-software/ApiRestPuntoVenta/models"
	"github.com/sindessis-software/ApiRestPuntoVenta/routes"
)

func main() {

	db.DBConnection()

	db.DB.AutoMigrate(models.Acceso{})
	db.DB.AutoMigrate(models.User{})

	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)

	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users", routes.CreateUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")

	//Seccion access

	r.HandleFunc("/access", routes.GetTasksHandler).Methods("GET")
	r.HandleFunc("/access", routes.CreateTaskHandler).Methods("POST")
	r.HandleFunc("/access/{id}", routes.GetTaskHandler).Methods("GET")
	r.HandleFunc("/access/{id}", routes.DeleteTasksHandler).Methods("DELETE")

	http.ListenAndServe(":3000", r)

}
