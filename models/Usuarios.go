package models

import "gorm.io/gorm"

type Usuarios struct {
	gorm.Model
	Usuario_nombre    string `gorm:"not null" json:"nombre"`
	Usuario_apellido1 string `gorm:"not null" json:"apellido1"`
	Usuario_apellido2 string `gorm:"not null" json:"apellido2"`
	Usuario_password  string `gorm:"not null" json:"password"`
	Usuario_email     string `gorm:"not null;unique_index" json:"email"`
	Usuario_permisos  string `gorm:"not null" json:"permisos"`
	Usuario_activo    string `gorm:"not null" json:"activo"`
}
