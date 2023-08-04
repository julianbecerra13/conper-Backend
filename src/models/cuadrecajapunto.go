package models

type CuadreCajaPunto struct {
	NombrePunto  string `gorm:"column:nombrePunto"`
	NombreTipoPago  string `gorm:"column:nombreTipoPago"`
	TotalOrdenes int    `gorm:"column:TotalOrdenes"`
	TotalVenta   float64    `gorm:"column:TotalVenta"`
}