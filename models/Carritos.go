package models

import "gorm.io/gorm"

type Carritos struct {
	gorm.Model

	Carrito_User         string `gorm:"not null" json:"Carrito_User"`
	Carrito_Costumer     string `gorm:"not null" json:"Carrito_Costumer"`
	Carrito_Status       string `gorm:"not null" json:"Carrito_Status"`
	Carrito_Fecha_compra string `gorm:"not null" json:"Fecha_Compra"`
	Carrito_Fecha_Update string `gorm:"not null" json:"Fecha_Update"`
	//cls	Carrito_Productos    []Producto
}
