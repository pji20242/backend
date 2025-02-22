package models

type MqttUser struct {
	ID             int    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Username       string `gorm:"column:username;unique;not null" json:"username"`
	Password       string `gorm:"column:pw;not null" json:"password"`
	MosquittoSuper bool   `gorm:"column:mosquitto_super;not null;default:false" json:"mosquitto_super"`
}

func (MqttUser) TableName() string {
	return "mqttusers"
}

type MqttAcl struct {
	ID       int    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Username string `gorm:"column:username;not null" json:"username"`
	Topic    string `gorm:"column:topic;not null" json:"topic"`
	Rw       int    `gorm:"column:rw;not null;default:1" json:"rw"`
}

func (MqttAcl) TableName() string {
	return "mqttacls"
}
