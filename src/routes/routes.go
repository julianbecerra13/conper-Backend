package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/myperri/copner/src/controllers"
)

func Routes() {
	route := gin.Default()
	route.Use(cors.Default())

	route.POST("/login", controllers.Login)
	route.PUT("/actualizar", controllers.Actualizar)
	route.GET("/pedidos", controllers.Pedidos)
	route.GET("/domicilios", controllers.Domicilios)
	route.GET("/domiciliarios", controllers.Domiciliarios)

	route.Run()
}
