package handlers

import (
	"net/http"
	"strconv"

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

// @Summary Cria um novo usuário
// @Description Cria um novo usuário no banco de dados
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "Novo Usuário"
// @Success 201 {object} models.User
// @Failure 400 {object} map[string]interface{}
// @Router /users [post]
func CreateUser(c *gin.Context) {
	var user models.User

	// Bind JSON body to user struct
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Insere no banco de dados
	result := database.GetDB().Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// @Summary Deleta um usuário
// @Description Deleta um usuário do banco de dados
// @Tags users
// @Accept json
// @Produce json
// @Param matricula path string true "Matricula do usuário"
// @Success 204 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /users/{matricula} [delete]
func DeleteUser(c *gin.Context) {
	matricula := c.Param("matricula")

	// Verifica se a matricula está vazia
	if matricula == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Matricula vazia"})
		return
	}

	// Converte a matricula para int
	matriculaInt, err := strconv.Atoi(matricula)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Deleta o usuário no banco de dados
	result := database.GetDB().Where("matricula = ?", matriculaInt).Delete(&models.User{})
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
