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
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	// Necessário para comparar erros de busca
	"gorm.io/gorm"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// AuthMiddleware extrai e valida o token, e cria o usuário se não existir.
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientID := os.Getenv("CLIENT_ID")
		clientSecret := os.Getenv("CLIENT_SECRET")

		if clientID == "" || clientSecret == "" {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "CLIENT_ID ou CLIENT_SECRET não definidos"})
			return
		}

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header não informado"})
			return
		}

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
			RedirectURL:  "postmessage", // Usado para troca de código
			Endpoint:     google.Endpoint,
		}

		// Troca o código pelo token
		token, err := config.Exchange(context.Background(), code)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Falha ao trocar código por token: %v", err)})
			return
		}

		idToken, ok := token.Extra("id_token").(string)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token JWT não encontrado"})
			return
		}

		parser := new(jwt.Parser)
		parsed, _, err := parser.ParseUnverified(idToken, jwt.MapClaims{})
		if err != nil {
			log.Printf("Error parsing token: %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error parsing token: %v", err)})
			return
		}

		claims, ok := parsed.Claims.(jwt.MapClaims)
		if !ok {
			log.Printf("Failed to parse token claims")
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse token claims"})
			return
		}

		email, emailOk := claims["email"].(string)
		name, nameOk := claims["name"].(string)
		matricula, matriculaOk := claims["sub"].(string)
		if !emailOk || !nameOk || !matriculaOk {
			log.Printf("Token não contém os campos necessários")
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Token não contém email e/ou name e/ou matricula"})
			return
		}

		log.Printf("Dados extraídos do token - Email: %s, Nome: %s, Matricula: %s", email, name, matricula)

		// Busca o usuário pelo email
		var user models.User
		result := database.GetDB().Where("email = ?", email).First(&user)
		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				// Se não encontrado, cria o usuário
				user = models.User{
					Nome:      name,
					Email:     email,
					User:      strings.Split(email, "@")[0],
					Senha:     "",
					Ativo:     true,
					Matricula: matricula,
				}
				if err := database.GetDB().Create(&user).Error; err != nil {
					log.Printf("Erro ao criar usuário: %v", err)
					c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Falha ao criar usuário"})
					return
				}
				log.Printf("Usuário criado com sucesso: %+v", user)
			} else {
				log.Printf("Erro ao buscar usuário: %v", result.Error)
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar usuário"})
				return
			}
		} else {
			log.Printf("Usuário já existe: %+v", user)
		}

		// Armazena o usuário no contexto para uso posterior
		c.Set("user", user)
		c.Next()
	}
}

func main() {
	// Inicializa o banco de dados com debug
	database.InitDatabase()

	r := gin.Default()

	// Grupo de rotas com middleware de autenticação
	v1 := r.Group("/api/v1")
	v1.Use(AuthMiddleware())
	{
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
