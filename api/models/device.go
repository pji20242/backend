package models

type Device struct {
	UUID      string  `gorm:"column:uuid;primaryKey" json:"uuid"`
	HWVersion string  `gorm:"column:hw_version" json:"hardware_version"`
	FWVersion string  `gorm:"column:fw_version" json:"firmware_version"`
	Latitude  float64 `gorm:"column:latitude" json:"latitude"`
	Longitude float64 `gorm:"column:longitude" json:"longitude"`
	Peso      float64 `gorm:"column:peso" json:"peso"`
	Alt       float64 `gorm:"column:alt" json:"altitude"`
}

func (Device) TableName() string {
	return "dispositivo"
}
