package models

type AggDomiciliarios struct {
	Nombre         string `gorm:"column:nombre"`
	Identificacion string `gorm:"column:identificacion"`
	Telefono       string `gorm:"column:telefono"`
	IdPunto        int    `gorm:"column:idPunto"`
}
