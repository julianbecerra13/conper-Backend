package models

type Novedades struct {
	ConcecutivoNovedad int    `gorm:"column:concecutivoNovedad"`
	Marca              string `gorm:"column:marca"`
	NombrePuntoVenta   string `gorm:"column:nombrePuntoVenta"`
	Novedad            string `gorm:"column:novedad"`
	FechaCrea          string `gorm:"column:fechaCrea"`
}
