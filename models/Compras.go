package models

import "gorm.io/gorm"

type Compras struct {
	gorm.Model
	Compras_User          string `gorm:"not null" json:"Compras_User"`
	Compras_Costumer      string `gorm:"not null" json:"Compras_Costumer"`
	Compras_status        string `gorm:"not null" json:"Carrito_Status"`
	Compras_Fecha_compra  string `gorm:"not null" json:"Fecha_Compra"`
	Compras_Fecha_Update  string `gorm:"not null" json:"Fecha_Update"`
	Compras_Fecha_Entrega string `gorm:"not null" json:"Fecha_Entrega"`
	//	Compras_Producto      []Producto
}
