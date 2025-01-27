package handlers

import (
	"net/http"

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

type DataResponse struct {
	Timestamp string  `json:"timestamp"`
	IDSensor  int     `json:"id_sensor"`
	Valor     float64 `json:"valor"`
}

