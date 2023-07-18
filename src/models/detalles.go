package models

type Detalles struct {
	Cantidad int `gorm:"column:cantidad"`
	ItemNombre string `gorm:"column:itemNombre"`
	ValorBaseUni string `gorm:"column:valorBaseUni"`
	ValorTotal int `gorm:"column:valorTotal"`
}