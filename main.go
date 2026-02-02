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
	db.DB.AutoMigrate(models.Carritos{})
	db.DB.AutoMigrate(models.Cliente{})
	db.DB.AutoMigrate(models.Compras{})
	db.DB.AutoMigrate(models.Usuarios{})
	db.DB.AutoMigrate(models.Producto{})
	db.DB.AutoMigrate(models.Factura{})

	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)

	//Seccion Usuarios

	r.HandleFunc("/usuarios", routes.GetUsuariosHandler).Methods("GET")
	r.HandleFunc("/usuarios", routes.CreateUsuariosHandler).Methods("POST")
	r.HandleFunc("/usuariosValida", routes.ValidateUsuariosHandler).Methods("POST")
	r.HandleFunc("/usuarios/{id}", routes.GetUsuarioHandler).Methods("GET")
	r.HandleFunc("/usuarios", routes.DeleteUsuariosHandler).Methods("DELETE")

	//Seccion Clientes

	r.HandleFunc("/cliente", routes.GetClientesHandler).Methods("GET")
	r.HandleFunc("/cliente", routes.CreateClientesHandler).Methods("POST")
	r.HandleFunc("/cliente/{id}", routes.GetClienteHandler).Methods("GET")
	r.HandleFunc("/cliente/{id}", routes.DeleteClientesHandler).Methods("DELETE")

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

	http.ListenAndServe(":3000", r)

}
