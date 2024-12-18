package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/pji20242/backend/api/database"
	_ "github.com/pji20242/backend/api/docs"
	"github.com/pji20242/backend/api/handlers"
)

// @title AgroTech API
// @version 1.0
// @description API para o sistema AgroTech
// @host localhost:8080
// @BasePath /api/v1
func main() {
	database.InitDatabase()

	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/users", handlers.ListUsers)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}
