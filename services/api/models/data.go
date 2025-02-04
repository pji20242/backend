package models

import "time"

type Data struct {
	Timestamp time.Time `gorm:"column:ts;primaryKey" json:"timestamp"`
	IDSensor  int       `gorm:"column:idSensor" json:"id_sensor"`
	Valor     float64   `gorm:"column:valor" json:"valor"`
}

func (Data) TableName() string {
	return "dados"
}
