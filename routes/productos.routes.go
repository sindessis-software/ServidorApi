package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sindessis-software/ApiRestPuntoVenta/db"
	"github.com/sindessis-software/ApiRestPuntoVenta/models"
)

func GetProductosHandler(w http.ResponseWriter, r *http.Request) {
	var products []models.Producto
	db.DB.Find(&products)
	json.NewEncoder(w).Encode(&products)
}

func CreateProductosHandler(w http.ResponseWriter, r *http.Request) {
	var product models.Producto
	json.NewDecoder(r.Body).Decode(&product)
	createproduct := db.DB.Create(&product)
	err := createproduct.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(&product)
}

func GetProductoHandler(w http.ResponseWriter, r *http.Request) {
	var product models.Producto
	params := mux.Vars(r)
	db.DB.First(&product, params["id"])
	if product.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Product not Found"))
		return
	}
	//db.DB.Model(&product).Association("Unidades").Find(&product.Unidades)

	json.NewEncoder(w).Encode(product)
}

func DeleteProductoHandler(w http.ResponseWriter, r *http.Request) {
	var product models.Producto
	params := mux.Vars(r)
	db.DB.First(&product, params["id"])
	if product.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Producto not Found"))
		return
	}
	db.DB.Unscoped().Delete(&product)
	w.WriteHeader(http.StatusNoContent)
}
