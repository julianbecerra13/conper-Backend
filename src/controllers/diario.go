package controllers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/myperri/copner/src/models"
)

func Diario(c *gin.Context) {
    var datosventa []models.Datosventas

    // Ejecutar la consulta SQL
    result := db.Raw("call spcp_consulta_venta_diaria()").Scan(&datosventa)
    if result.Error != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"datos": datosventa})
}