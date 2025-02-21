package models

type User struct {
	Matricula string `gorm:"column:matricula;primaryKey" json:"matricula"`
	Nome      string `gorm:"column:nome;not null" json:"nome"`
	Email     string `gorm:"column:email;unique;not null" json:"email"`
	Senha     string `gorm:"column:senha;not null" json:"-"`
	User      string `gorm:"column:user;unique;not null" json:"user"`
	Ativo     bool   `gorm:"column:ativo;default:true" json:"ativo"`
	Licencas  int    `gorm:"column:licencas;default:0" json:"licencas"`
}

func (User) TableName() string {
	return "usuario"
}
