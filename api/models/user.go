package models

type User struct {
	Matricula  int    `gorm:"column:matricula;primaryKey" json:"matricula"`
	Nome       string `gorm:"column:nome;not null" json:"nome"`
	Email      string `gorm:"column:email;unique;not null" json:"email"`
	Senha      string `gorm:"column:senha;not null" json:"-"` // Omit from JSON response
	Privilegio string `gorm:"column:privilegio" json:"privilegio"`
	Ativo      bool   `gorm:"column:ativo;default:true" json:"ativo"`
}

func (User) TableName() string {
	return "usuario"
}
