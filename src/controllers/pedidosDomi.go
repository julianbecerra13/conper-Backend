package controllers

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/myperri/copner/src/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type PedidosDomiRequest struct {
	IdDomiciliario string `json:"idDomiciliario"`
	IdPunto       string `json:"idPunto"`
}

func GetDB() (*gorm.DB, error) {
	// Cargar las variables de entorno desde el archivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return nil, err
	}

	// Obtener las variables de entorno necesarias para la conexión a la base de datos
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	// Establecer la conexión a la base de datos
	dsn := dbUser + ":" + dbPass + "@tcp(" + dbHost + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database")
		return nil, err
	}

	return db, nil
}

func PedidosDomi(c *gin.Context) {
	var pedidosDomiRequest PedidosDomiRequest
	var pedidosDomiResponse []models.PedidosDomi

	// Capturar los parametros de la url
	pedidosDomiRequest.IdDomiciliario = c.Query("idDomiciliario")
	pedidosDomiRequest.IdPunto = c.Query("idPunto")

	// Obtener la conexión a la base de datos utilizando GetDB
	db, err := GetDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error connecting to the database"})
		return
	}

	// Ejecutar la consulta directamente
	query := `SELECT idOrdenGeneral,idOrdenNumero,nombreCliente,telefonoCelular,direccion,CONCAT('$ ', FORMAT(totalOrden, 0)) AS totalOrden,fechaCrea
	FROM tcpr_ordengeneral
	WHERE date(fechaCrea) = curdate() AND idTipoVenta = 2 AND idEstadoOrden = 2 AND idPunto=? AND idDomiciliario=? LIMIT 0,100`

	result := db.Raw(query, pedidosDomiRequest.IdPunto, pedidosDomiRequest.IdDomiciliario).Scan(&pedidosDomiResponse)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"pedidos": pedidosDomiResponse})
	print(pedidosDomiResponse)
}
