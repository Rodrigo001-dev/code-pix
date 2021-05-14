package model

import (
	"time"

	"github.com/asaskevich/govalidator"
)

func init() {
	// vai validar todos os campos que são obrigados serem enviados
	govalidator.SetFieldsRequiredByDefault(true)
}

type Base struct {
	// id do banco, chave identificadora unica
	ID        string    `json:"id" valid:"uuid"`
	CreatedAt time.Time `json:"created_at" valid:"-"`
	UpdatedAt time.Time `json:"updated_at" valid:"-"`
}
