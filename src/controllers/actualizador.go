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
	IdAdicional    string `json:"idAdicional"`
	Campo          string `json:"campo"`
	IdDomiciliario string `json:"idDomiciliario"`
}

func Actualizar(c *gin.Context) {
	var actualizadorRequest ActualizadorRequest

	actualizadorRequest.IdPunto = c.Query("idPunto")
	actualizadorRequest.IdPedido = c.Query("idPedido")
	actualizadorRequest.IdUsuario = c.Query("idUsuario")
	actualizadorRequest.IdAdicional = c.Query("idAdicional")
	actualizadorRequest.IdTraza = c.Query("idTraza")
	actualizadorRequest.Campo = c.Query("campo")
	actualizadorRequest.IdDomiciliario = c.Query("idDomiciliario")

	result := db.Raw("CALL spcp_sil_creatraza_motorizado(?, ?, ?, ?, ?, ?, ?)", actualizadorRequest.IdPunto, actualizadorRequest.IdPedido, actualizadorRequest.IdUsuario, actualizadorRequest.IdAdicional, actualizadorRequest.IdTraza, actualizadorRequest.Campo, actualizadorRequest.IdDomiciliario)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"actualizado": true})
}
