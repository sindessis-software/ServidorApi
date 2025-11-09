package models

import (
	"gorm.io/gorm"
)

type Producto struct {
	gorm.Model

	TipoProduct    uint   `json:"tipo_producto"`
	DescProducto   string `json:"descripcion_producto"`
	StatusProducto bool   `gorm:"default:false" json:"status_producto"`
	UnidadProducto []Unidades
}
