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

// @Summary Cria uma novo dispositivo
// @Description Cria um novo dispositivo no banco de dados
// @Tags devices
// @Accept json
// @Produce json
// @Param device body models.Device true "Novo Dispositivo"
// @Success 201 {object} models.Device
// @Failure 400 {object} map[string]interface{}
// @Router /devices [post]

func CreateDevice(c *gin.Context) {
	var device models.Device

	// Bind JSON body to device struct
	if err := c.ShouldBindJSON(&device); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Insere no banco de dados
	result := database.GetDB().Create(&device)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, device)
}

// @Summary Deleta um dispositivo
// @Description Deleta um dispositivo do banco de dados
// @Tags devices
// @Accept json
// @Produce json
// @Param id path string true "ID do dispositivo"
// @Success 204 {object} string
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /devices/{id} [delete]
func DeleteDevice(c *gin.Context) {
	uuid := c.Param("uuid")

	// Deleta o dispositivo no banco de dados
	result := database.GetDB().Where("uuid = ?", uuid).Delete(&models.Device{})
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Dispositivo não encontrado"})
		return
	}

	c.JSON(http.StatusNoContent, "Deletado com sucesso")
}