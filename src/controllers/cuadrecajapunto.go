package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/myperri/copner/src/models"
)

type CuadreCajaPuntoRequest struct {
	IDusuario string `json:"idDomiciliario"`
	IDpunto        string `json:"idPunto"`
	FechaInicio    string `json:"fechaInicio"`
	FechaFin       string `json:"fechaFin"`
}

func CuadreCajaPunto(c *gin.Context) {

	var cuadreCajaPuntoRequest CuadreCajaPuntoRequest
	var cuadreCajaPunto []models.CuadreCajaPunto

	if err := c.ShouldBindJSON(&cuadreCajaPuntoRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := db.Raw("CALL spcp_sil_cuadrecajapunto(?,?,?,?)", cuadreCajaPuntoRequest.IDusuario, cuadreCajaPuntoRequest.IDpunto, cuadreCajaPuntoRequest.FechaInicio, cuadreCajaPuntoRequest.FechaFin).Scan(&cuadreCajaPunto)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"cuadreCaja": cuadreCajaPunto})
}
