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
	db.DB.AutoMigrate(models.Producto{})
	//	db.DB.AutoMigrate(models.Unidades{})

	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)

	//Seccion Users

	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users", routes.CreateUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")

	//Seccion Access

	r.HandleFunc("/access", routes.GetAccesosHandler).Methods("GET")
	r.HandleFunc("/access", routes.CreateAccesoHandler).Methods("POST")
	r.HandleFunc("/access/{id}", routes.GetAccesoHandler).Methods("GET")
	r.HandleFunc("/access/{id}", routes.DeleteAccesosHandler).Methods("DELETE")

	//Seccion Productos

	r.HandleFunc("/productos", routes.GetProductosHandler).Methods("GET")
	r.HandleFunc("/productos", routes.CreateProductosHandler).Methods("POST")
	r.HandleFunc("/productos/{id}", routes.GetProductoHandler).Methods("GET")
	r.HandleFunc("/productos/{id}", routes.DeleteProductoHandler).Methods("DELETE")

	//Seccion Unidades

	r.HandleFunc("/unidades", routes.GetUnidadesHandler).Methods("GET")
	r.HandleFunc("/unidades", routes.CreateUnidadHandler).Methods("POST")
	r.HandleFunc("/unidades/{id}", routes.GetUnidadHandler).Methods("GET")
	r.HandleFunc("/unidades/{id}", routes.DeleteUnidadHandler).Methods("DELETE")

	http.ListenAndServe(":3000", r)

}
