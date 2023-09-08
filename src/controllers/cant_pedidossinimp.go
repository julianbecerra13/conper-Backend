package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/myperri/copner/src/models"// Reemplaza "your-package-name" con el nombre de tu paquete
)



func ObtenerNumero(c *gin.Context) {
	var numero []models.Numero

	// Realizar la consulta SELECT
	result := db.Table("tcpr_ordengeneral").
		Select("count(1) as cantidad").
		Where("impresion = ? AND idTipoVenta = ? AND idEstadoOrden = ?", 0, 2, 2).
		Scan(&numero)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"numero": numero})
}
