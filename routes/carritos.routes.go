package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sindessis-software/ApiRestPuntoVenta/db"
	"github.com/sindessis-software/ApiRestPuntoVenta/models"
)

func GetCarritosHandler(w http.ResponseWriter, r *http.Request) {
	var carritos []models.Carritos
	db.DB.Find(&carritos)
	json.NewEncoder(w).Encode(&carritos)
}

func CreateCarritosHandler(w http.ResponseWriter, r *http.Request) {
	var carritos []models.Carritos
	json.NewDecoder(r.Body).Decode(&carritos)
	createCarritos := db.DB.Create(&carritos)
	err := createCarritos.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(&carritos)
}

func GetCarritoHandler(w http.ResponseWriter, r *http.Request) {
	var carritos []models.Carritos
	params := mux.Vars(r)
	db.DB.First(&carritos, params["id"])
	/*if carritos.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Carritos not Found"))
		return
	}*/
	json.NewEncoder(w).Encode(carritos)
}

func DeleteCarritosHandler(w http.ResponseWriter, r *http.Request) {
	var carritos []models.Carritos
	params := mux.Vars(r)
	db.DB.First(&carritos, params["id"])
	/*if carritos.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Accesos not Found"))
		return
	}*/
	db.DB.Unscoped().Delete(&carritos)
	w.WriteHeader(http.StatusNoContent)
}
