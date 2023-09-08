package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/myperri/copner/src/models" 
	"gorm.io/gorm"
)

func ObtenerNumeroMovil(c *gin.Context) {
	var numero []models.Numero

	// Realizar la consulta SELECT con JOIN
	result := db.Table("tcpr_ordengeneral og").
		Select("count(1) as cantidad").
		Joins("LEFT JOIN tcpr_trazasdia ttd ON og.idOrdenGeneral = ttd.idOrdenGeneral").
		Where("og.impresion = ? AND og.idTipoVenta = ? AND og.idEstadoOrden = ? AND og.idDomiciliario != ? AND date(og.FechaCrea) = CURDATE() AND ttd.idTraza = ? AND ttd.fechaFin = ?", 1, 2, 2, 0, 5, "1981-01-07 00:00:00").
		Limit(100).
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
