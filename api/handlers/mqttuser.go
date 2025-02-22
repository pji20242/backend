package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pji20242/backend/api/database"
	"github.com/pji20242/backend/api/models"
)

// @Summary Lista usuários MQTT
// @Description Obtém a lista de usuários MQTT
// @Tags mqttusers
// @Produce json
// @Success 200 {array} models.MqttUser
// @Router /mqttusers [get]
func ListMqttUsers(c *gin.Context) {
	var users []models.MqttUser
	result := database.GetMQTTDB().Find(&users)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, users)
}

// @Summary Cria um novo usuário MQTT
// @Description Cria um novo usuário MQTT no banco de dados
// @Tags mqttusers
// @Accept json
// @Produce json
// @Param user body models.MqttUser true "Novo Usuário MQTT"
// @Success 201 {object} models.MqttUser
// @Failure 400 {object} map[string]interface{}
// @Router /mqttusers [post]
func CreateMqttUser(c *gin.Context) {
	var user models.MqttUser

	// Bind JSON body to user struct
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Insere no banco de dados
	result := database.GetMQTTDB().Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// @Summary Deleta um usuário MQTT
// @Description Deleta um usuário MQTT do banco de dados
// @Tags mqttusers
// @Accept json
// @Produce json
// @Param username path string true "Nome de usuário MQTT"
// @Success 204 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /mqttusers/{username} [delete]
func DeleteMqttUser(c *gin.Context) {
	username := c.Param("username")

	// Verifica se o username está vazio
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username vazio"})
		return
	}

	// Deleta o usuário no banco de dados
	result := database.GetMQTTDB().Where("username = ?", username).Delete(&models.MqttUser{})
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
		return
	}

	c.JSON(http.StatusNoContent, "Deletado com sucesso")
}
