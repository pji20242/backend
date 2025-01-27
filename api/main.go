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
		v1.GET("/users", handlers.ListUsers)
		v1.GET("/cooperativas", handlers.ListCooperativas)
		v1.GET("/devices", handlers.ListDevices)
		v1.GET("/devices/:uuid", handlers.GetDeviceData)
		v1.GET("/map", handlers.GetDeviceMap)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}
