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
	route.GET("/domicilios", controllers.Domicilios)

	route.Run()
}
