package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RequestData struct {
	FromData map[string]interface{} `json:"fromData"`
	Call     string                 `json:"Call"`
}

func HandleDatos(c *gin.Context) {
	var requestData RequestData

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("Ejecutando procedimiento almacenado: CALL %s\n", requestData.Call)
	fmt.Printf("Parámetros: %v\n", requestData.FromData)

	var rows *sql.Rows
	var err error

	// Verifica si FromData contiene valores
	if len(requestData.FromData) > 0 {
		// Si hay valores en FromData, crea una consulta con parámetros
		query := "CALL " + requestData.Call + "("
		params := []interface{}{}
		for _, value := range requestData.FromData {
			query += "?, "
			params = append(params, value)
		}
		query = query[:len(query)-2] + ")"

		rows, err = db.Raw(query, params...).Rows()
	} else {
		// Si no hay valores en FromData, ejecuta el procedimiento almacenado sin parámetros
		rows, err = db.Raw("CALL " + requestData.Call).Rows()
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	// Convierte los resultados en un mapa
	result := make(map[string]interface{})
	columns, err := rows.Columns()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("Columnas: %v\n", columns) // Punto de impresión

	data := []map[string]interface{}{} // Almacena los datos de filas

	// Get the column types
	columnTypes, err := rows.ColumnTypes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for rows.Next() {
		values := make([]interface{}, len(columns))
		scanArgs := make([]interface{}, len(columns))

		for i := range columns {
			// Use the column types to determine the appropriate scan type
			switch columnTypes[i].DatabaseTypeName() {
			case "INT", "TINYINT", "SMALLINT", "MEDIUMINT", "BIGINT":
				scanArgs[i] = new(int64)
			default:
				scanArgs[i] = new(string)
			}

			values[i] = scanArgs[i]
		}

		if err := rows.Scan(values...); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		fmt.Printf("Valores escaneados: %v\n", values) // Punto de impresión

		row := make(map[string]interface{})
		for i, col := range columns {
			// Cast the scanned values to the appropriate data types
			switch columnTypes[i].DatabaseTypeName() {
			case "INT", "TINYINT", "SMALLINT", "MEDIUMINT", "BIGINT":
				row[col] = *(scanArgs[i].(*int64))
			default:
				row[col] = *(scanArgs[i].(*string))
			}
		}

		fmt.Printf("Fila escaneada: %v\n", row) // Punto de impresión

		// Agrega esta fila al resultado
		data = append(data, row)
	}

	// Formatea los datos en el formato deseado
	formattedData := [][]interface{}{}
	for _, row := range data {
		formattedRow := []interface{}{row["idPunto"], row["nombre"]}
		formattedData = append(formattedData, formattedRow)
	}

	fmt.Printf("Datos finales: %v\n", formattedData) // Punto de impresión

	if len(formattedData) > 0 {
		result["columns"] = columns
		result["data"] = formattedData
	} else {
		// Si no hay datos, establece un mensaje personalizado
		result["message"] = "Sin datos"
	}

	// Convierte el resultado en JSON y devuelve la respuesta
	responseJSON, err := json.Marshal(result)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, string(responseJSON))
}
