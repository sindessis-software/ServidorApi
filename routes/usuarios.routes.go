package routes

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sindessis-software/ApiRestPuntoVenta/db"
	"github.com/sindessis-software/ApiRestPuntoVenta/models"
	"gorm.io/gorm"
)

func GetUsuariosHandler(w http.ResponseWriter, r *http.Request) {
	var Usuarios []models.Usuarios
	db.DB.Find(&Usuarios)
	json.NewEncoder(w).Encode(&Usuarios)
}

func GetUsuarioHandler(w http.ResponseWriter, r *http.Request) {
	var Usuarios models.Usuarios
	params := mux.Vars(r)
	db.DB.First(&Usuarios, params["id"])
	if Usuarios.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Usuarios not Found"))
		return
	}
	json.NewEncoder(w).Encode(&Usuarios)
}

func CreateUsuariosHandler(w http.ResponseWriter, r *http.Request) {
	var Usuarios models.Usuarios
	json.NewDecoder(r.Body).Decode(&Usuarios)
	createdUsuarios := db.DB.Create(&Usuarios)
	err := createdUsuarios.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusCreated)
	}
	json.NewEncoder(w).Encode(&Usuarios)
}

func ValidateUsuariosHandler(w http.ResponseWriter, r *http.Request) {
	var usuarios models.Usuarios
	err := json.NewDecoder(r.Body).Decode(&usuarios)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	result := db.DB.First(&models.Usuarios{}, "Usuario_email = ? AND Usuario_password = ?", usuarios.Usuario_email, usuarios.Usuario_password)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Usuarios not Found"))
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error no conocido"))
		return
	}
	json.NewEncoder(w).Encode(&usuarios)
}

func DeleteUsuariosHandler(w http.ResponseWriter, r *http.Request) {
	var usuarios models.Usuarios
	err := json.NewDecoder(r.Body).Decode(&usuarios)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	result := db.DB.Unscoped().Delete(&models.Usuarios{}, "Usuario_email = ? AND Usuario_password = ?", usuarios.Usuario_email, usuarios.Usuario_password)
	if result.Error == nil && result.RowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Usuarios not Found"))
		return
	}
	w.WriteHeader(http.StatusOK)
}
