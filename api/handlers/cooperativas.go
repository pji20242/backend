package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pji20242/backend/api/database"
	"github.com/pji20242/backend/api/models"
)

// @Summary Lista cooperativas
// @Description Obtém a lista de todas as cooperativas
// @Tags cooperativas
// @Produce json
// @Success 200 {array} models.Cooperativa
// @Router /cooperativas [get]
func ListCooperativas(c *gin.Context) {
	var cooperativas []models.Cooperativa
	result := database.GetDB().Find(&cooperativas)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, cooperativas)
}
