package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pji20242/backend/api/database"
	"github.com/pji20242/backend/api/models"
)

// @Summary Lista ACLs MQTT
// @Description Obtém a lista de ACLs MQTT
// @Tags mqttacls
// @Produce json
// @Success 200 {array} models.MqttAcl
// @Router /mqttacls [get]
func ListMqttAcls(c *gin.Context) {
	var acls []models.MqttAcl
	result := database.GetMQTTDB().Find(&acls)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, acls)
}

// @Summary Cria uma nova ACL MQTT
// @Description Cria uma nova ACL MQTT no banco de dados
// @Tags mqttacls
// @Accept json
// @Produce json
// @Param acl body models.MqttAcl true "Nova ACL MQTT"
// @Success 201 {object} models.MqttAcl
// @Failure 400 {object} map[string]interface{}
// @Router /mqttacls [post]
func CreateMqttAcl(c *gin.Context) {
	var acl models.MqttAcl

	// Bind JSON body to acl struct
	if err := c.ShouldBindJSON(&acl); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Insere no banco de dados
	result := database.GetMQTTDB().Create(&acl)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, acl)
}

// @Summary Deleta uma ACL MQTT
// @Description Deleta uma ACL MQTT do banco de dados
// @Tags mqttacls
// @Accept json
// @Produce json
// @Param username path string true "Nome de usuário MQTT"
// @Param topic path string true "Nome de tópico MQTT"
// @Success 204 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /mqttacls/{username}/{topic} [delete]
func DeleteMqttAcl(c *gin.Context) {
	username := c.Param("username")
	topic := c.Param("topic")

	// Verifica se os parâmetros estão vazios
	if username == "" || topic == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username ou tópico vazio"})
		return
	}

	// Deleta a ACL no banco de dados
	result := database.GetMQTTDB().Where("username = ? AND topic = ?", username, topic).Delete(&models.MqttAcl{})
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "ACL não encontrada"})
		return
	}

	c.JSON(http.StatusNoContent, "Deletado com sucesso")
}
