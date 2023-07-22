package models

type CuadreCaja struct {
	NombreMovil    string `gorm:"column:nombreMovil"`
	NombrePunto    string `gorm:"column:nombrePunto"`
	NombreTipoPago string `gorm:"column:nomTipoPago"`
	TotalFp        string    `gorm:"column:TotalFP"`
	TotalOrdenes   int    `gorm:"column:TotalOrdenes"`
	TotalVenta     float64    `gorm:"column:ValorFP"`
}
