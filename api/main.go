package main

import (
	"context"
	"errors"
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

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			log.Printf("Authorization header não informado")
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			log.Printf("Formato do header Authorization inválido")
		}

		token := parts[1]
		log.Printf("Código de autorização extraído: %s", token)

		parser := new(jwt.Parser)
		parsed, _, err := parser.ParseUnverified(token, jwt.MapClaims{})
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

type AuthCode struct {
	Code string `json:"code"`
}

func Authenticate(c *gin.Context) {
	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")

	if clientID == "" || clientSecret == "" {
		log.Printf("CLIENT_ID ou CLIENT_SECRET não definidos")
	}

	config := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  "postmessage",
		Endpoint:     google.Endpoint,
	}

	// Recupera o código de autorização
	var authCode AuthCode
	if err := c.ShouldBindJSON(&authCode); err != nil {
		log.Printf("Falha ao recuperar código de autorização: %v", err)
		// Aqui você pode continuar com o fluxo, mas sem bloquear a execução
	}
	code := authCode.Code

	token, err := config.Exchange(context.Background(), code)
	if err != nil {
		log.Printf("Falha ao trocar código por token: %v", err)
	}

	// Verifica se o token contém o campo id_token
	idToken, ok := token.Extra("id_token").(string)
	if !ok {
		log.Printf("Token JWT não encontrado ou malformado")
	}

	c.JSON(http.StatusOK, gin.H{"token": idToken})
}

func main() {
	// Inicializa o banco de dados com debug
	database.InitDatabase()
	database.InitMQTTDatabase()

	r := gin.Default()

	// Grupo de rotas com middleware de autenticação
	v1 := r.Group("/api/v1")
	{
		v1.POST("/auth", Authenticate)

		// Grupo protegido com autenticação
		protected := v1.Group("/")
		protected.Use(AuthMiddleware())
		{
			protected.GET("/map", handlers.GetDeviceMap)
			protected.GET("/users", handlers.ListUsers)
			protected.GET("/cooperativas", handlers.ListCooperativas)
			protected.GET("/devices", handlers.ListDevices)
			protected.GET("/devices/:uuid", handlers.GetDeviceData)
			protected.GET("/devices/:uuid/sensor/:idSensor", handlers.GetSensorData)
			protected.GET("/sensores", handlers.ListSensors)
			protected.GET("/sensores/:uuid", handlers.GetSensorsByUUID)
			protected.POST("/cooperativas", handlers.CreateCooperativa)
			protected.POST("/devices", handlers.CreateDevice)
			protected.POST("/users", handlers.CreateUser)
			protected.POST("/sensores", handlers.CreateSensor)
			protected.DELETE("/cooperativas/:cnpj", handlers.DeleteCooperativa)
			protected.DELETE("/devices/:uuid", handlers.DeleteDevice)
			protected.DELETE("/users/:matricula", handlers.DeleteUser)
			protected.DELETE("/devices/:uuid/sensores/:id", handlers.DeleteSensor)

			// Rotas MQTT
			protected.GET("/mqttusers", handlers.ListMqttUsers)
			protected.POST("/mqttusers", handlers.CreateMqttUser)
			protected.DELETE("/mqttusers/:username", handlers.DeleteMqttUser)
			protected.GET("/mqttacls", handlers.ListMqttAcls)
			protected.POST("/mqttacls", handlers.CreateMqttAcl)
			protected.DELETE("/mqttacls/:username/:topic", handlers.DeleteMqttAcl)
		}
	}

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Inicia o servidor
	r.Run(":8080")
}
