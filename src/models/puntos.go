package models

type Punto struct {
	IdPunto int `gorm:"column:idPunto"`
	Nombre string `gorm:"column:nombre"`
}