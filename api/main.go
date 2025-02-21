package main

import (
	"context"
	"errors"
	"log"
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

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientID := os.Getenv("CLIENT_ID")
		clientSecret := os.Getenv("CLIENT_SECRET")

		if clientID == "" || clientSecret == "" {
			log.Printf("CLIENT_ID ou CLIENT_SECRET não definidos")
		}

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			log.Printf("Authorization header não informado")
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			log.Printf("Formato do header Authorization inválido")
		}

		code := parts[1]
		log.Printf("Código de autorização extraído: %s", code)

		config := &oauth2.Config{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			RedirectURL:  "postmessage",
			Endpoint:     google.Endpoint,
		}

		token, err := config.Exchange(context.Background(), code)
		if err != nil {
			log.Printf("Falha ao trocar código por token: %v", err)
			// Aqui você pode continuar com o fluxo, mas sem bloquear a execução
		}

		// Verifica se o token contém o campo id_token
		idToken, ok := token.Extra("id_token").(string)
		if !ok {
			log.Printf("Token JWT não encontrado ou malformado")
		}

		parser := new(jwt.Parser)
		parsed, _, err := parser.ParseUnverified(idToken, jwt.MapClaims{})
		if err != nil {
			log.Printf("Erro ao analisar token: %v", err)
		}

		claims, ok := parsed.Claims.(jwt.MapClaims)
		if !ok {
			log.Printf("Falha ao converter claims do token")
		}

		email, emailOk := claims["email"].(string)
		name, nameOk := claims["name"].(string)
		matricula, matriculaOk := claims["sub"].(string)
		if !emailOk || !nameOk || !matriculaOk {
			log.Printf("Token não contém os campos necessários")
		}

		log.Printf("Dados extraídos do token - Email: %s, Nome: %s, Matricula: %s", email, name, matricula)

		// Verifica se o usuário já existe no banco de dados
		var user models.User
		db := database.GetDB()
		result := db.Where("email = ?", email).First(&user)

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// Usuário não encontrado, então cria um novo
			user = models.User{
				Nome:      name,
				Email:     email,
				User:      strings.Split(email, "@")[0],
				Senha:     "",
				Ativo:     true,
				Matricula: matricula,
			}
			if err := db.Create(&user).Error; err != nil {
				log.Printf("Erro ao criar usuário: %v", err)
			} else {
				log.Printf("Usuário criado com sucesso: %+v", user)
			}
		} else if result.Error != nil {
			log.Printf("Erro ao buscar usuário: %v", result.Error)
		} else {
			log.Printf("Usuário já existe: %+v", user)
		}

		// Armazena o usuário no contexto, mesmo que com erros
		c.Set("user", user)
		log.Printf("Usuário armazenado no contexto: %+v", user)

		// Continua a execução da próxima função na pilha de middleware
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
