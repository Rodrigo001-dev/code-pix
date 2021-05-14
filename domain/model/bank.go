package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type Bank struct {
	Base `valid:"required"`
	// código do banco
	Code string `json:"code" gorm:"type:varchar(20)" valid:"notnull"`
	// nome do banco
	Name     string     `json:"name" gorm:"type:varchar(255)" valid:"notnull"`
	Acconuts []*Account `gorm:"ForeignKey:BankID" valid:"-"`
}

// criando um método para a validação do banco e para isso o método está
// entrelassado com o Bank
func (bank *Bank) isValid() error {
	_, err := govalidator.ValidateStruct(bank)
	if err != nil {
		return err
	}

	return nil
}

// essa fiunção vai me retornar um banco ou um erro
func NewBank(code string, name string) (*Bank, error) {
	bank := Bank{
		Code: code,
		Name: name,
	}

	bank.ID = uuid.NewV4().String()
	bank.CreatedAt = time.Now()

	err := bank.isValid()
	if err != nil {
		// se tiver algun erro vai retornar nulo no lugar do banco(nil) e vai
		// retornar o erro
		return nil, err
	}

	return &bank, nil
}
