package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBMQTT *gorm.DB

func InitMQTTDatabase() {
	dsn := "host=dbmqtt user=postgres password=postgres dbname=mqtt port=5432 sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to MQTT database:", err)
	}

	// Ativa o modo debug para exibir as queries executadas
	db = db.Debug()

	// Se necessário, adicionar migração de modelos específicos do banco MQTT
	// db.AutoMigrate(&models.SomeMQTTModel{})

	DBMQTT = db
}

func GetMQTTDB() *gorm.DB {
	return DBMQTT
}
