package models

import (
	"gorm.io/gorm"
)

type Acceso struct {
	gorm.Model

	TypeAccess uint   `json:"type_access"`
	DescAccess string `json:"description_access"`
	RollAccess uint   `json:"roll_access"`
	Status     bool   `gorm:"default:false" json:"status"`
	UserID     uint   `json:"user_id"`
}
