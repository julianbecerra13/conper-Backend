package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/myperri/copner/src/controllers"
	"time"
)

func Routes() {
	route := gin.Default()

    config := cors.DefaultConfig()
    config.AllowAllOrigins = true
    config.AllowMethods = []string{"GET", "POST", "PUT", "OPTIONS"}
    config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type"}
    config.ExposeHeaders = []string{"Content-Length"}
    config.AllowCredentials = true
    config.MaxAge = 12 * time.Hour

    route.Use(cors.New(config))

	route.POST("/login", controllers.Login)
	route.PUT("/actualizar", controllers.Actualizar)
	route.PUT("/actualizarT", controllers.ActualizarT)
	route.GET("/pedidos", controllers.Pedidos)
	route.GET("/domicilios", controllers.Domicilios)
	route.GET("/domiciliarios", controllers.Domiciliarios)
	route.GET("/novedades", controllers.Novedades)
	route.PUT("/aggdomiciliarios", controllers.AggDomiciliarios)
	route.PUT("/aggdomiciliariosn2", controllers.AggDomiciliariosn2)
	route.PUT("/detalles", controllers.Detalles)
	route.PUT("/impresora", controllers.Impresora)
	route.PUT("/aggnovedad1", controllers.AggNovedad)
	route.PUT("/archivopost", controllers.GenerarArchivoPO1)
	route.GET("/pqrs", controllers.Pqrs)
	route.PUT("/respuestaPqrs", controllers.RespuestaPqrs)
	route.GET("/cuadrecajadomi", controllers.CuadreCajaDomi)
	route.GET("/cuadrecajapunto", controllers.CuadreCajaPunto)
	route.GET("/PedidosDomi", controllers.PedidosDomi)
	route.PUT("/transferir", controllers.Transferir)
	route.GET("/reportepedidos", controllers.ReportePedidos)
	route.GET("/cantidadsinimp", controllers.ObtenerNumero)
	route.GET("/cantidadTotal", controllers.ObtenerNumeroTotal)
	route.GET("/cantidadEnGestion", controllers.ObtenerNumeroEnGestion)
	route.GET("/cantidadMovil", controllers.ObtenerNumeroMovil)
	route.GET("/datosgraficobarra", controllers.ObtenerReporteMensual)
	route.GET("/reporteinfo", controllers.Reportesinfo)
	route.POST("/procesar", controllers.HandleDatos)
	route.GET("/puntos", controllers.Puntos)
	route.GET("/diarioventas", controllers.Diario)
	route.GET("/mensualventas", controllers.Mensuales)
	route.GET("/anualventas", controllers.Anual)
	route.POST("/parametrosreportes", controllers.ParametrosReportes)


	route.Run()
}
