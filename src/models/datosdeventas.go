package models

type Datosventas struct {
	TotalOrdenes int `gorm:"column:totalOrdenes"`
	TotalVenta string `gorm:"column:totalVenta"`
}