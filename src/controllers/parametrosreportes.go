package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/myperri/copner/src/models"
)

type parametrosRequest struct {
	Id int `json:"id"`
}

func ParametrosReportes(c *gin.Context) {
	var dato parametrosRequest

	if err := c.ShouldBindJSON(&dato); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Obtener los parámetros
	var parametros []models.Parametro
	db.Raw("Call spcp_reportes(?)", dato.Id).Select("nombreParametro, queryParametro").Scan(&parametros)

	// Mapear los resultados al formato deseado
	parametrosResponse := make([]models.Parametro, len(parametros))
	for i, r := range parametros {
		parametrosResponse[i].Parametro = r.Parametro
	}

	// Recolectar datos de parámetros combo
	parametrosComboData := make(map[string]interface{})

	for _, r := range parametros {
		if r.QueryParametro != "" {
			var datoscombo []models.Combo
			db.Raw(r.QueryParametro).Scan(&datoscombo)
			parametrosComboData[r.Parametro] = datoscombo
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"parametros":          parametrosResponse,
		"parametrosComboData": parametrosComboData,
	})
}
