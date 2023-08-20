package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/myperri/copner/src/models"
)

type ReportePedidosRequest struct {
	FechaInicio    string `json:"fechaInicio"`
	FechaFin       string `json:"fechaFin"`
	IdDomiciliario string `json:"idDomiciliario"`
}

func ReportePedidos(c *gin.Context) {
	var reportePedidosRequest ReportePedidosRequest
	var reportePedidos []models.ReportePedido

	reportePedidosRequest.FechaInicio = c.Query("fechaInicio")
	reportePedidosRequest.FechaFin = c.Query("fechaFin")
	idDomiciliarioStr := c.Query("idDomiciliario")

	idDomiciliario, err := strconv.Atoi(idDomiciliarioStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "idDomiciliario debe ser un número válido"})
		return
	}

	result := db.Raw("call tcpc_sil_app_reportePedidos(?, ?, ?)", reportePedidosRequest.FechaInicio, reportePedidosRequest.FechaFin, idDomiciliario).Scan(&reportePedidos)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"reportePedidos": reportePedidos})
}
