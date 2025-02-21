package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/pji20242/backend/api/database"
	"github.com/pji20242/backend/api/handlers"
	"github.com/pji20242/backend/api/models"
	"google.golang.org/api/idtoken"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)



// AuthMiddleware extrai o token do header Authorization e o valida.
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientID := os.Getenv("CLIENT_ID")
		// Verifica se o CLIENT_ID foi definido
		if clientID == "" {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "CLIENT_ID não definido"})
			return
		}
		// Extrai o header Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header não informado"})
			return
		}

		// Espera o formato "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Formato do header Authorization inválido"})
			return
		}

		token := parts[1]
		// Valida o ID Token usando a biblioteca do Google.
		ctx := context.Background()
		payload, err := idtoken.Validate(ctx, token, clientID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("Token inválido: %v", err)})
			return
		}

		// Extrai o email do payload (assumindo que o email está no payload)
		email, ok := payload.Claims["email"].(string)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Email não encontrado no token"})
			return
		}

		// Busca o usuário no banco de dados pelo email
		var user models.User
		if err := database.GetDB().Where("email = ?", email).First(&user).Error; err != nil {
			// Se o usuário não existir, cria um novo usuário
			user = models.User{
				Nome:  payload.Claims["name"].(string), // Assumindo que o nome está no payload
				Email: email,
				User:  strings.Split(email, "@")[0], // Gera um nome de usuário a partir do email
				Senha: "", // Senha pode ser deixada em branco ou gerada automaticamente
				Ativo: true,
			}
			if err := database.GetDB().Create(&user).Error; err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Falha ao criar usuário"})
				return
			}
		}

		// Armazena o usuário no contexto para uso posterior nas rotas
		c.Set("user", user)
		c.Next()
	}
}

func main() {
	database.InitDatabase()

	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		// Rotas públicas
		v1.GET("/map", handlers.GetDeviceMap)

		// Rotas protegidas por autenticação
		authGroup := v1.Group("/")
		authGroup.Use(AuthMiddleware())
		{
			authGroup.GET("/users", handlers.ListUsers)
			authGroup.GET("/cooperativas", handlers.ListCooperativas)
			authGroup.GET("/devices", handlers.ListDevices)
			authGroup.GET("/devices/:uuid", handlers.GetDeviceData)
			authGroup.GET("/devices/:uuid/sensor/:idSensor", handlers.GetSensorData)
			authGroup.GET("/sensores", handlers.ListSensors)
			authGroup.POST("/cooperativas", handlers.CreateCooperativa)
			authGroup.POST("/devices", handlers.CreateDevice)
			authGroup.POST("/users", handlers.CreateUser)
			authGroup.POST("/sensores", handlers.CreateSensor)
			authGroup.DELETE("/cooperativas/:cnpj", handlers.DeleteCooperativa)
			authGroup.DELETE("/devices/:uuid", handlers.DeleteDevice)
			authGroup.DELETE("/users/:matricula", handlers.DeleteUser)
			authGroup.DELETE("/devices/:uuid/sensores/:id", handlers.DeleteSensor)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}