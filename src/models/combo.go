package models

type Combo struct {
	Codigo int    `gorm:"column:codigo"`
	Nombre string `gorm:"column:nombre"`
}
