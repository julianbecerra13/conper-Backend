package models

type PedidosDomi struct {
	IdPedido        int    `gorm:"column: idOrdenGeneral"`
	NombreCliente   string `gorm:"column:nombreCliente"`
	TelefonoCelular string `gorm:"column:telefonoCelular"`
	Direccion       string `gorm:"column:direccion"`
	TotalOrden      string `gorm:"column:totalOrden"`
	FechaCrea       string `gorm:"column:fechaCrea"`
}
 