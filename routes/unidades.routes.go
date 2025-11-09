package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sindessis-software/ApiRestPuntoVenta/db"
	"github.com/sindessis-software/ApiRestPuntoVenta/models"
)

func GetUnidadesHandler(w http.ResponseWriter, r *http.Request) {
	var unidades []models.Unidades
	db.DB.Find(&unidades)
	json.NewEncoder(w).Encode(&unidades)
}

func CreateUnidadHandler(w http.ResponseWriter, r *http.Request) {
	var unidad models.Unidades
	json.NewDecoder(r.Body).Decode(&unidad)
	createunidad := db.DB.Create(&unidad)
	err := createunidad.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(&unidad)
}

func GetUnidadHandler(w http.ResponseWriter, r *http.Request) {
	var unidad models.Unidades
	params := mux.Vars(r)
	db.DB.First(&unidad, params["id"])
	if unidad.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Unidad Compra not Found"))
		return
	}
	json.NewEncoder(w).Encode(unidad)
}

func DeleteUnidadHandler(w http.ResponseWriter, r *http.Request) {
	var unidad models.Unidades
	params := mux.Vars(r)
	db.DB.First(&unidad, params["id"])
	if unidad.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Unidad Compra not Found"))
		return
	}
	db.DB.Unscoped().Delete(&unidad)
	w.WriteHeader(http.StatusNoContent)
}
