package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sindessis-software/ApiRestPuntoVenta/db"
	"github.com/sindessis-software/ApiRestPuntoVenta/models"
)

func GetComprasHandler(w http.ResponseWriter, r *http.Request) {
	var Compras []models.Compras
	db.DB.Find(&Compras)
	json.NewEncoder(w).Encode(&Compras)
}

func CreateComprasHandler(w http.ResponseWriter, r *http.Request) {
	var compras []models.Compras
	json.NewDecoder(r.Body).Decode(&compras)
	createCompras := db.DB.Create(&compras)
	err := createCompras.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(&compras)
}

func GetCompraHandler(w http.ResponseWriter, r *http.Request) {
	var compras []models.Compras
	params := mux.Vars(r)
	db.DB.First(&compras, params["id"])
	/*if compras.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Compras not Found"))
		return
	}*/
	json.NewEncoder(w).Encode(compras)
}

func DeleteComprasHandler(w http.ResponseWriter, r *http.Request) {
	var compras []models.Compras
	params := mux.Vars(r)
	db.DB.First(&compras, params["id"])
	/*if compras.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Accesos not Found"))
		return
	}*/
	db.DB.Unscoped().Delete(&compras)
	w.WriteHeader(http.StatusNoContent)
}
