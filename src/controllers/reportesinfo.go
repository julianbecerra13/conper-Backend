package controllers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/myperri/copner/src/models"
)

func Reportesinfo(c *gin.Context) {
    var reportes []models.Reportesinfo 

    // Ejecutar la consulta SQL
    result := db.Raw("SELECT * FROM tcpc_reportes").Scan(&reportes)
    if result.Error != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"reportes": reportes})
}
