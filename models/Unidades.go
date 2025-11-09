package models

import (
	"gorm.io/gorm"
)

type Unidades struct {
	gorm.Model

	TypeUnidad   uint   `json:"type_Unidad"`
	DescUnidad   string `json:"description_Unidad"`
	StatusUnidad bool   `gorm:"default:false" json:"status_Unidad"`
	UnidadVenta  bool   `gorm:"default:false" json:"Venta_status"`
	UnidadCompra bool   `gorm:"default:false" json:"Compra_status"`
	ProductoID   uint   `json:"product_id"`
}
