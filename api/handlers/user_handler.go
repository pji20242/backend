package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pji20242/backend/api/database"
	"github.com/pji20242/backend/api/models"
)

// @Summary Lista usuários
// @Description Obtém a lista de usuários
// @Tags users
// @Produce json
// @Success 200 {array} models.User
// @Router /users [get]
func ListUsers(c *gin.Context) {
	var users []models.User
	result := database.GetDB().Find(&users)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, users)
}
