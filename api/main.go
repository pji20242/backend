package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pji20242/backend/api/database"
	_ "github.com/pji20242/backend/api/docs"
	"github.com/pji20242/backend/api/handlers"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"google.golang.org/api/idtoken"
)

const clientID = "803624329648-o2hggrtbmtdqeld9v8io4inuprus79am.apps.googleusercontent.com"

// AuthMiddleware extrai o token do header Authorization e o valida.
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
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
		// Substitua pelo seu Client ID registrado no Google.
		ctx := context.Background()
		payload, err := idtoken.Validate(ctx, token, clientID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("Token inválido: %v", err)})
			return
		}

		// TODO: PEGA DADOS DO USUÁRIOS DO BANCO DE DADOS
		// TODO: CRIA USUÁRIO SE NÃO EXISTIR

		// Opcional: armazena o payload no contexto para uso posterior nas rotas
		c.Set("tokenPayload", payload)
		c.Next()
	}
}

// @title API para Sistema AgroTech
// @version 1.0
// @description Esta API permite interações com dispositivos e usuários do sistema AgroTech
// @contact.name arthurcadore
// @contact.url arthurcadore.github.io
// @contact.email arthurbarcella.ifsc@gmail.com
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @host localhost:8080
// @BasePath /api/v1
func main() {
	database.InitDatabase()

	r := gin.Default()

	v1 := r.Group("/api/v1")
	v1.Use(AuthMiddleware())
	{
		// Endpoints para GET
		v1.GET("/users", handlers.ListUsers)
		v1.GET("/cooperativas", handlers.ListCooperativas)
		v1.GET("/devices", handlers.ListDevices)
		v1.GET("/devices/:uuid", handlers.GetDeviceData)
		v1.GET("/devices/:uuid/sensor/:idSensor", handlers.GetSensorData) // Novo endpoint para sensor específico
		v1.GET("/map", handlers.GetDeviceMap)
		v1.GET("/sensores", handlers.ListSensors)

		// Endpoints para POST
		v1.POST("/cooperativas", handlers.CreateCooperativa)
		v1.POST("/devices", handlers.CreateDevice)
		v1.POST("/users", handlers.CreateUser)
		v1.POST("/sensores", handlers.CreateSensor)

		// Endpoints para DELETE
		v1.DELETE("/cooperativas/:cnpj", handlers.DeleteCooperativa)
		v1.DELETE("/devices/:uuid", handlers.DeleteDevice)
		v1.DELETE("/users/:matricula", handlers.DeleteUser)
		v1.DELETE("/devices/:uuid/sensores/:id", handlers.DeleteSensor)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}
