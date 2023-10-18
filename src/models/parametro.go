package models

type Parametro struct {
    Parametro string `gorm:"column:nombreParametro"`
    QueryParametro string `gorm:"column:queryParametro"`
    
}