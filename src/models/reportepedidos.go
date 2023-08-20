package models

type ReportePedido struct {
	NombreMovil string `gorm:"column:nombreMovil"`
	NombrePunto string `gorm:"column:nombrePunto"`
	TotalOP     int    `gorm:"column:TotalOP"`
	ValorFP     string `gorm:"column:ValorFP"`
}
