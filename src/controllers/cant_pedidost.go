package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/myperri/copner/src/models" // Reemplaza "your-package-name" con el nombre de tu paquete
	"gorm.io/gorm"
)

func ObtenerNumeroTotal(c *gin.Context) {
	var numero []models.Numero

	// Realizar la nueva consulta SELECT
	result := db.Table("tcpr_ordengeneral").
		Select("count(1) as cantidad").
		Where("idTipoVenta = ? AND idEstadoOrden = ? AND date(fechaCrea) = CURDATE()", 2, 2).
		Scan(&numero)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			// Si no se encuentra ningún registro, devolver un valor predeterminado
			c.JSON(http.StatusOK, gin.H{"numero": 0})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"numero": numero})
}
