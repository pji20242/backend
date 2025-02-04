package models

type Cooperativa struct {
	CNPJ     string `gorm:"column:cnpj;primaryKey" json:"cnpj"`
	Endereco string `gorm:"column:endereco" json:"endereco"`
	Email    string `gorm:"column:email" json:"email"`
	Nome     string `gorm:"column:nome;not null" json:"nome"`
}

func (Cooperativa) TableName() string {
	return "cooperativa"
}
