package models

type ReporteMensual struct {
	Mes            string `gorm:"column:mes_palabra"`
	VentaEfectiva  string `gorm:"column:ventaEfectiva"`
	VentaCancelada string `gorm:"column:ventaCancelada"`
	TotalVenta     string `gorm:"column:totalVenta"`
}
