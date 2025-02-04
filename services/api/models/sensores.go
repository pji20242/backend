package models

type Sensors struct {
	ID		uint      `gorm:"column:idSensor;primaryKey" json:"id"`
	UUID   string    `gorm:"column:uuid" json:"uuid"`
	Tipo   string    `gorm:"column:tipo" json:"tipo"`
	Unidade string   `gorm:"column:unidade" json:"unidade"`
}

func (Sensors) TableName() string {
	return "sensor"
}
