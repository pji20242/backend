package handlers

import (
	"net/http"
	"strconv"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/pji20242/backend/api/database"
	"github.com/pji20242/backend/api/models"
)

// @Summary Lista sensores
// @Description Obtém a lista de todos os sensores
// @Tags sensores
// @Produce json
// @Success 200 {array} models.Sensors
// @Router /sensores [get]
func ListSensors(c *gin.Context) {
	var sensors []models.Sensors
	result := database.GetDB().Find(&sensors)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, sensors)
}

// @Summary Lista sensores de um dispositivo
// @Description Retorna a lista de sensores do dispositivo com base no UUID fornecido
// @Tags sensores
// @Accept json
// @Produce json
// @Param uuid path string true "UUID do dispositivo"
// @Success 200 {array} models.Sensors
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/sensores/{uuid} [get]
func GetSensorsByUUID(c *gin.Context) {
	deviceUUID := c.Param("uuid")
	
	var sensors []models.Sensors
	result := database.GetDB().Where("uuid = ?", deviceUUID).Find(&sensors)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	
	if len(sensors) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Nenhum sensor encontrado para esse dispositivo"})
		return
	}
	
	c.JSON(http.StatusOK, sensors)
}

// @Summary Cria um novo sensor
// @Description Cria um novo sensor no banco de dados
// @Tags sensores
// @Accept json
// @Produce json
// @Param sensor body models.Sensors true "Novo Sensor"
// @Success 201 {object} models.Sensors
// @Failure 400 {object} map[string]interface{}
// @Router /sensores [post]
func CreateSensor(c *gin.Context) {
	var sensor models.Sensors

	// Bind JSON body to sensor struct
	if err := c.ShouldBindJSON(&sensor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Insere no banco de dados
	result := database.GetDB().Create(&sensor)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, sensor)
}

// @Summary Deleta um sensor de um dispositivo específico
// @Description Deleta um sensor de um dispositivo com base no UUID do dispositivo e ID do sensor
// @Tags sensores
// @Accept json
// @Produce json
// @Param uuid path string true "UUID do dispositivo"
// @Param id path int true "ID do sensor"
// @Success 204 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /devices/{uuid}/sensores/{id} [delete]
func DeleteSensor(c *gin.Context) {
	// Recuperando os parâmetros da URL
	deviceUUID := c.Param("uuid")
	sensorID := c.Param("id")

	// Convertendo o ID do sensor para inteiro
	sensorIDInt, err := strconv.Atoi(sensorID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID do sensor inválido"})
		return
	}

	// Buscando o sensor no banco de dados pelo UUID do dispositivo e ID do sensor
	var sensor models.Sensors
	result := database.GetDB().Where("uuid = ? AND idSensor = ?", deviceUUID, sensorIDInt).First(&sensor)

	// Verificando se o sensor foi encontrado
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Sensor não encontrado"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		}
		return
	}

	// Deletando o sensor
	deleteResult := database.GetDB().Delete(&sensor)
	if deleteResult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": deleteResult.Error.Error()})
		return
	}

	// Verificando se o sensor foi realmente deletado
	if deleteResult.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sensor não encontrado"})
		return
	}

	// Retornando sucesso
	c.JSON(http.StatusNoContent, gin.H{"message": "Sensor deletado com sucesso"})
}
