package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/myperri/copner/src/models" 
	"gorm.io/gorm"
)

func ObtenerReporteMensual(c *gin.Context) {
	var reporteMensual []models.ReporteMensual

	// Llama al procedimiento almacenado
	result := db.Raw("CALL spcp_reporte_graficoventames()").Scan(&reporteMensual)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			// Si no se encuentra ningún registro, devolver un valor predeterminado
			c.JSON(http.StatusOK, gin.H{"reporteMensual": []models.ReporteMensual{}})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reporteMensual": reporteMensual})
}
