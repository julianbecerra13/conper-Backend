package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ActualizadorRequest struct {
	IdPunto        string `json:"idPunto"`
	IdPedido       string `json:"idPedido"`
	IdUsuario      string `json:"idUsuario"`
	IdTraza        string `json:"idTraza"`
	IdDomiciliario string `json:"idDomiciliario"`
}

func Actualizar(c *gin.Context) {
	var actualizadorRequest ActualizadorRequest

	if err := c.ShouldBindJSON(&actualizadorRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := db.Exec("CALL spcp_sil_creatraza_motorizado(?, ?, ?, ?, ?, ?, ?)", actualizadorRequest.IdPunto, actualizadorRequest.IdPedido, actualizadorRequest.IdUsuario, "1", actualizadorRequest.IdTraza, "", actualizadorRequest.IdDomiciliario)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"actualizado": true})
}
