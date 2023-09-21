package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/myperri/copner/src/config"
)

type ReportRequest struct {
	Cadena string `json:"Cadena"`
	Param1 string `json:"Param1"`
	Param2 string `json:"Param2"`
	Param3 string `json:"Param3"`
	Param4 string `json:"Param4"`
	Param5 string `json:"Param5"`
}

func Inforeportes(c *gin.Context) {
	var reportRequest ReportRequest
	
	// Decodifica el JSON del cuerpo de la solicitud en la estructura ReportRequest
	if err := c.ShouldBindJSON(&reportRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Obtiene la conexión a la base de datos utilizando ConnectDB() de config.go
	db := config.ConnectDB()

	// Importante: No cierres manualmente la conexión aquí, GORM lo manejará automáticamente

	// Construye el procedimiento almacenado con los parámetros adecuados
	procedureCall := "CALL " + reportRequest.Cadena + "(?, ?, ?, ?, ?)"

	// Ejecuta el procedimiento almacenado con los parámetros proporcionados
	rows, err := db.Raw(procedureCall, reportRequest.Param1, reportRequest.Param2, reportRequest.Param3, reportRequest.Param4, reportRequest.Param5).Rows()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	// Construye una matriz de interfaz para almacenar los resultados
	results := make([]map[string]interface{}, 0)

	// Obtiene las columnas de los resultados
	columns, err := rows.Columns()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Crea una matriz para almacenar los valores escaneados de cada fila
	values := make([]interface{}, len(columns))
	valuePtrs := make([]interface{}, len(columns))

	// Inicializa los valuePtrs con las direcciones de los valores correspondientes en values
	for i := range values {
		valuePtrs[i] = &values[i]
	}

	// Recorre las filas
	for rows.Next() {
		// Escanea los valores en values
		if err := rows.Scan(valuePtrs...); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Crea un mapa para almacenar los valores de esta fila
		rowData := make(map[string]interface{})

		// Recorre los valores escaneados y los convierte al tipo adecuado
		for i, col := range columns {
			val := values[i]
			if val == nil {
				rowData[col] = nil
			} else {
				rowData[col] = string(val.([]byte))
			}
		}

		// Agrega el mapa de esta fila a los resultados
		results = append(results, rowData)
	}

	// Convierte los resultados a formato JSON y responde con ellos
	c.JSON(http.StatusOK, gin.H{"data": results})
}
