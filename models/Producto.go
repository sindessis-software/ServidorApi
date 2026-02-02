package models

import (
	"gorm.io/gorm"
)

type Producto struct {
	gorm.Model

	NameProducto    string `gorm:"not null" json:"nombre_producto"`
	DescProducto    string `gorm:"not null" json:"descripcion_producto"`
	TipoProducto    string `gorm:"not null" json:"tipo_producto"`
	PrecioProducto  string `gorm:"not null" json:"precio_producto"`
	Unidadproducto  string `gorm:"not null" json:"unidad_producto"`
	NameImgproducto string `gorm:"not null" json:"nombre_imagen_producto"`
	StatusProducto  bool   `gorm:"default:false" json:"status_producto"`
}
