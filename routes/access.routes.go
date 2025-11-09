package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sindessis-software/ApiRestPuntoVenta/db"
	"github.com/sindessis-software/ApiRestPuntoVenta/models"
)

func GetAccesosHandler(w http.ResponseWriter, r *http.Request) {
	var acceso []models.Acceso
	db.DB.Find(&acceso)
	json.NewEncoder(w).Encode(&acceso)
}

func CreateAccesoHandler(w http.ResponseWriter, r *http.Request) {
	var acceso models.Acceso
	json.NewDecoder(r.Body).Decode(&acceso)
	createAcceso := db.DB.Create(&acceso)
	err := createAcceso.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(&acceso)
}

func GetAccesoHandler(w http.ResponseWriter, r *http.Request) {
	var acceso models.Acceso
	params := mux.Vars(r)
	db.DB.First(&acceso, params["id"])
	if acceso.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Acceso not Found"))
		return
	}
	json.NewEncoder(w).Encode(acceso)
}

func DeleteAccesosHandler(w http.ResponseWriter, r *http.Request) {
	var acceso models.Acceso
	params := mux.Vars(r)
	db.DB.First(&acceso, params["id"])
	if acceso.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Accesos not Found"))
		return
	}
	db.DB.Unscoped().Delete(&acceso)
	w.WriteHeader(http.StatusNoContent)
}
