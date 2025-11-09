package models

import (
	"gorm.io/gorm"
)

type Cliente struct {
	gorm.Model

	TipoCliente      uint   `json:"tipo_cliente"`
	NombreCliente    string `json:"nombre_cliente"`
	DescCliente      string `json:"descripcion_cliente"`
	DireccionCliente string `json:"direccion_cliente"`
	ColoniaCliente   string `json:"colonia_cliente"`
	MunicipioCliente string `json:"municipio_cliente"`
	CiudadCliente    string `json:"ciudad_cliente"`
	PaisCliente      string `json:"Pais_cliente"`
	StatusCliente    bool   `gorm:"default:false" json:"status_cliente"`
	Facturacion      []Factura
}
