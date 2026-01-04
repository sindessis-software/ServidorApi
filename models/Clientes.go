package models

import (
	"gorm.io/gorm"
)

type Cliente struct {
	gorm.Model

	Cliente_Tipo_Cliente string `gorm:"not null" json:"tipo_cliente"`
	Cliente_Nombre       string `gorm:"not null" json:"nombre_cliente"`
	Cliente_Descripcion  string `gorm:"not null" json:"descripcion_cliente"`
	Cliente_Rfc          string `gorm:"not null" json:"rfc_cliente"`
	Cliente_Direccion    string `gorm:"not null" json:"direccion_cliente"`
	Cliente_Colonia      string `gorm:"not null" json:"colonia_cliente"`
	Cliente_Municipio    string `gorm:"not null" json:"municipio_cliente"`
	Cliente_Ciudad       string `gorm:"not null" json:"ciudad_cliente"`
	Cliente_Pais         string `gorm:"not null" json:"Pais_cliente"`
	Cliente_Status       bool   `gorm:"default:false" json:"status_cliente"`
}
