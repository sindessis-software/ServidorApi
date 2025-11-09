package models

import (
	"gorm.io/gorm"
)

type Factura struct {
	gorm.Model

	TipoFactura   uint   `json:"tipo_Factura"`
	DescFactura   string `json:"descripcion_Factura"`
	StatusFactura bool   `gorm:"default:false" json:"status_Factura"`
	ClienteID     uint   `json:"cliente_id"`
}
