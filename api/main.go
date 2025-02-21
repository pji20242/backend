package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/pji20242/backend/api/database"
	"github.com/pji20242/backend/api/handlers"
	"github.com/pji20242/backend/api/models"
	"golang.org/x/oauth2"        // Importação correta do pacote OAuth2
	"golang.org/x/oauth2/google" // Importação para configuração do Google OAuth2

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// AuthMiddleware extrai o token do header Authorization e o valida.
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientID := os.Getenv("CLIENT_ID")
		clientSecret := os.Getenv("CLIENT_SECRET")

		// Verifica se as variáveis de ambiente estão definidas
		if clientID == "" || clientSecret == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "CLIENT_ID ou CLIENT_SECRET não definidos"})
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

		code := parts[1]

		// Configura o cliente OAuth2
		config := &oauth2.Config{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			RedirectURL:  "postmessage",   // Usado para troca de código
			Endpoint:     google.Endpoint, // Endpoint do Google OAuth2
		}

		// Troca o código pelo token
		token, err := config.Exchange(context.Background(), code)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Falha ao trocar código por token: %v", err)})
			return
		}
		fmt.Println(token)

		idToken, ok := token.Extra("id_token").(string)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token JWT não encontrado"})
			return
		}
		fmt.Println("ID Token:", idToken)

		parser := new(jwt.Parser)
		// Parse the token without verifying the signature
		parsed, _, err := parser.ParseUnverified(idToken, jwt.MapClaims{})
		if err != nil {
			log.Fatalf("Error parsing token: %v", err)
		}

		claims, ok := parsed.Claims.(jwt.MapClaims)
		if !ok {
			log.Fatal("Failed to parse token claims")
		}

		email, emailOk := claims["email"].(string)
		name, nameOk := claims["name"].(string)
		if !emailOk || !nameOk {
			log.Fatal("Token does not contain email and/or name")
		}

		// Busca o usuário no banco de dados pelo email
		var user models.User
		if err := database.GetDB().Where("email = ?", email).First(&user).Error; err != nil {
			// Se o usuário não existir, cria um novo usuário
			user = models.User{
				Nome:  name, // Assumindo que o nome está no payload
				Email: email,
				User:  strings.Split(email, "@")[0], // Gera um nome de usuário a partir do email
				Senha: "",                           // Senha pode ser deixada em branco ou gerada automaticamente
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

	// Grupo para rotas da API v1
	v1 := r.Group("/api/v1")
	v1.Use(AuthMiddleware())
	{
		// Rotas públicas
		v1.GET("/map", handlers.GetDeviceMap)
		v1.GET("/users", handlers.ListUsers)
		v1.GET("/cooperativas", handlers.ListCooperativas)
		v1.GET("/devices", handlers.ListDevices)
		v1.GET("/devices/:uuid", handlers.GetDeviceData)
		v1.GET("/devices/:uuid/sensor/:idSensor", handlers.GetSensorData)
		v1.GET("/sensores", handlers.ListSensors)
		v1.POST("/cooperativas", handlers.CreateCooperativa)
		v1.POST("/devices", handlers.CreateDevice)
		v1.POST("/users", handlers.CreateUser)
		v1.POST("/sensores", handlers.CreateSensor)
		v1.DELETE("/cooperativas/:cnpj", handlers.DeleteCooperativa)
		v1.DELETE("/devices/:uuid", handlers.DeleteDevice)
		v1.DELETE("/users/:matricula", handlers.DeleteUser)
		v1.DELETE("/devices/:uuid/sensores/:id", handlers.DeleteSensor)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}
