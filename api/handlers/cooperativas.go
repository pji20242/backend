package handlers

import (
	"log"
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
	log.Println("Iniciando a busca por cooperativas")
	result := database.GetDB().Find(&cooperativas)

	if result.Error != nil {
		log.Printf("Erro ao consultar as cooperativas: %v", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	log.Printf("Cooperativas encontradas: %v", cooperativas)
	c.JSON(http.StatusOK, cooperativas)
}

// @Summary Cria uma nova cooperativa
// @Description Cria uma nova cooperativa no banco de dados
// @Tags cooperativas
// @Accept json
// @Produce json
// @Param cooperativa body models.Cooperativa true "Nova Cooperativa"
// @Success 201 {object} models.Cooperativa
// @Failure 400 {object} map[string]interface{}
// @Router /cooperativas [post]
func CreateCooperativa(c *gin.Context) {
	var cooperativa models.Cooperativa

	// Bind JSON body to cooperativa struct
	if err := c.ShouldBindJSON(&cooperativa); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Insere no banco de dados
	result := database.GetDB().Create(&cooperativa)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, cooperativa)
}

// @Summary Deleta uma cooperativa
// @Description Deleta uma cooperativa com base no CNPJ
// @Tags cooperativas
// @Accept json
// @Produce json
// @Param cnpj path string true "CNPJ da Cooperativa"
// @Success 204 {string} string "Deletado com sucesso"
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /cooperativas/{cnpj} [delete]
func DeleteCooperativa(c *gin.Context) {
	cnpj := c.Param("cnpj")

	// Deleta a cooperativa no banco de dados
	result := database.GetDB().Where("cnpj = ?", cnpj).Delete(&models.Cooperativa{})
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cooperativa não encontrada"})
		return
	}

	c.JSON(http.StatusNoContent, "Deletado com sucesso")
}
