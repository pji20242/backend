package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/pji20242/backend/api/database"
	_ "github.com/pji20242/backend/api/docs"
	"github.com/pji20242/backend/api/handlers"
)

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
	{
		// Endpoints para GET
		v1.GET("/users", handlers.ListUsers)
		v1.GET("/cooperativas", handlers.ListCooperativas)
		v1.GET("/devices", handlers.ListDevices)
		v1.GET("/devices/:uuid", handlers.GetDeviceData)
		v1.GET("/devices/:uuid/sensor/:idSensor", handlers.GetSensorData) // Novo endpoint para sensor específico
		v1.GET("/map", handlers.GetDeviceMap)
		v1.GET("/sensores", handlers.ListSensors)

		// Novo endpoint para retornar os sensores de um dispositivo específico pelo UUID
		v1.GET("/sensores/:uuid", handlers.GetSensorsByUUID)

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
