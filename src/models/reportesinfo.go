package models

type Reportesinfo struct {
	IdReporte int `gorm:"column:idReporte"`
	Nombre string `gorm:"column:nombre"`
	Descripcion string `gorm:"column:descripcion"`
	Call string `gorm:"column:nameQueryProcedureSQL"`
	Parametros string `gorm:"column:parametros"`

}
