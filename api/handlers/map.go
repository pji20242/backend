package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pji20242/backend/api/database"
	"github.com/pji20242/backend/api/models"
)

type DeviceMapResponse struct {
	UUID      string  `json:"uuid"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Peso      float64 `json:"peso"`
}

// @Summary Lista todos os dispositivos com localização e peso
// @Description Retorna os dispositivos com as colunas UUID, LATITUDE, LONGITUDE e PESO
// @Tags devices
// @Produce json
// @Success 200 {array} DeviceMapResponse
// @Router /map [get]
func GetDeviceMap(c *gin.Context) {
	var devices []models.Device
	result := database.GetDB().Find(&devices)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	// Mapear para a estrutura de resposta personalizada
	var response []DeviceMapResponse
	for _, device := range devices {
		response = append(response, DeviceMapResponse{
			UUID:      device.UUID,
			Latitude:  device.Latitude,
			Longitude: device.Longitude,
			Peso:      device.Peso,
		})
	}

	c.JSON(http.StatusOK, response)
}
