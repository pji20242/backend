package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pji20242/backend/api/database"
	"github.com/pji20242/backend/api/models"
)

// @Summary Lista dispositivos
// @Description Obtém a lista de todos os dispositivos
// @Tags devices
// @Produce json
// @Success 200 {array} models.Device
// @Router /devices [get]
func ListDevices(c *gin.Context) {
	var devices []models.Device
	result := database.GetDB().Find(&devices)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, devices)
}
