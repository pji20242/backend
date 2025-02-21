package database

import (
	"log"

	"github.com/pji20242/backend/api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	dsn := "root:rootpass@tcp(database:3306)/pjiot?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Ativa o modo debug para exibir as queries executadas
	db = db.Debug()

	// Migração dos modelos
	db.AutoMigrate(&models.User{}, &models.Cooperativa{}, &models.Device{}, &models.Data{}, &models.Sensors{})
	DB = db
}

func GetDB() *gorm.DB {
	return DB
}
