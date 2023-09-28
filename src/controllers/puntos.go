package controllers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/myperri/copner/src/models"
)

func Puntos(c *gin.Context) {
    var puntos []models.Punto 

    // Ejecutar la consulta SQL
    result := db.Raw("call spcp_sil_puntos()").Scan(&puntos)
    if result.Error != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"reportes": puntos})
}