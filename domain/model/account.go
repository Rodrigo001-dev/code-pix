package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type Account struct {
	Base      `valid:"required"`
	OwnerName string `gorm:"column:owner_name;type:varchar(255);not null" valid:"notnull"`
	Bank      *Bank  `valid:"-"`
	// adicionando o ID do banco(chave estrangeira)
	BankID  string    `gorm:"column:bank_id;type:uuid;not null" valid:"-"`
	Number  string    `json:"number" gorm:"type:varchar(20)" valid:"notnull"`
	PixKeys []*PixKey `gorm:"ForeignKey:AccountID" valid:"-"`
}

func (account *Account) isValid() error {
	_, err := govalidator.ValidateStruct(account)
	if err != nil {
		return err
	}

	return nil
}

func NewAccount(bank *Bank, number string, ownername string) (*Account, error) {
	account := Account{
		OwnerName: ownername,
		Bank:      bank,
		Number:    number,
	}

	account.ID = uuid.NewV4().String()
	account.CreatedAt = time.Now()

	err := account.isValid()
	if err != nil {
		return nil, err
	}

	return &account, nil
}
