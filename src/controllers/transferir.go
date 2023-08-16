package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransferirRequest struct {
	IdPedido       int `json:"idPedido"`
	IdDomiciliario int `json:"idDomiciliario"`
}

func Transferir(c *gin.Context) {
	var transferirRequest TransferirRequest

	if err := c.ShouldBindJSON(&transferirRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := db.Exec("call spcp_sil_cambiamovil(?,?)", transferirRequest.IdPedido, transferirRequest.IdDomiciliario)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"transferido": true})
}
