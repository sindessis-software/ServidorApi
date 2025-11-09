package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	User      string `gorm:"not null"`
	UserName  string `gorm:"not null"`
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Email     string `gorm:"not null;unique_index"`
	Acceso    []Acceso
}
