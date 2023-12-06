package models

type Domicilios struct {
	IDOrdenGeneral     int     `gorm:"column:idOrdenGeneral"`
	IdOrdenNumero      string  `gorm:"column:idOrdenNumero"`
	NombreCliente      string  `gorm:"column:nombreCliente"`
	DireccionOrden     string  `gorm:"column:direccionOrden"`
	TotalOrden         float64 `gorm:"column:totalOrden"`
	FechaCrea          string  `gorm:"column:fechaCrea"`
	NombreTraza        string  `gorm:"column:nombreTraza"`
	NombreDomiciliario string  `gorm:"column:domiciliario"`
	IdPunto            int     `gorm:"column:idPunto"`
	PuntodeVenta       string  `gorm:"column:PuntoVenta"`
	Observaciones      string  `gorm:"column:observaciones"`
	Telefono           string  `gorm:"column:telefono"`
}
