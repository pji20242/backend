package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pji20242/backend/api/database"
	"github.com/pji20242/backend/api/models"
)

// @Summary Lista dados do dispositivo
// @Description Obtém os dados de um dispositivo específico pelo UUID
// @Tags devices
// @Produce json
// @Param uuid path string true "UUID do dispositivo"
// @Success 200 {array} models.Data
// @Router /devices/{uuid} [get]
func GetDeviceData(c *gin.Context) {
	uuid := c.Param("uuid")

	var data []models.Data
	result := database.GetDB().Where("uuid = ?", uuid).Find(&data)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	if len(data) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Nenhum dado encontrado para o dispositivo especificado",
		})
		return
	}

	c.JSON(http.StatusOK, data)
}

// @Summary Lista dados de um sensor específico de um dispositivo
// @Description Obtém os dados de um sensor específico de um dispositivo pelo UUID e idSensor
// @Tags devices
// @Produce json
// @Param uuid path string true "UUID do dispositivo"
// @Param idSensor path int true "ID do sensor"
// @Success 200 {array} models.Data
// @Router /devices/{uuid}/sensor/{idSensor} [get]
func GetSensorData(c *gin.Context) {
	uuid := c.Param("uuid")
	idSensorStr := c.Param("idSensor")

	idSensor, err := strconv.Atoi(idSensorStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "idSensor deve ser um número inteiro"})
		return
	}

	var data []models.Data
	result := database.GetDB().Where("uuid = ? AND idSensor = ?", uuid, idSensor).Find(&data)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	if len(data) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Nenhum dado encontrado para o sensor especificado do dispositivo",
		})
		return
	}

	c.JSON(http.StatusOK, data)
}




