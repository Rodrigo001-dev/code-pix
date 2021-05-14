package model

import (
	"time"

	"errors"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type PixKeyRepositoryInterface interface {
	// vai tentar registrara uma chave e se não conseguir vai retornar um erro
	RegisterKey(pixKey *PixKey) (*PixKey, error)
	// vai tentar encontrar uma chave basiado pelo tipo(kind) que for passado
	// email ou CPF
	FindKeyByKind(key string, kind string) (*PixKey, error)
	// vai tentar adicionar um banco e se não conseguir vai retornar um erro
	AddBank(bank *Bank) error
	// vai tentar adicionar uma conta e se não conseguir vai retornar um erro
	AddAccount(account *Account) error
	// vai tentar encontrar uma conta e se não conseguir vai retornar um erro
	FindAccount(id string) (*Account, error)
}

type PixKey struct {
	Base      `valid:"required"`
	Kind      string   `json:"kind" valid:"notnull"`
	Key       string   `json:"key" valid:"notnull"`
	AccountID string   `gorm:"column:account_id;type:uuid;not null" valid:"-"`
	Account   *Account `valid:"-"`
	Status    string   `json:"status"`
}

func (pixKey *PixKey) isValid() error {
	_, err := govalidator.ValidateStruct(pixKey)

	if pixKey.Kind != "email" && pixKey.Kind != "cpf" {
		return errors.New("invalid type of key")
	}

	if pixKey.Status != "active" && pixKey.Status != "inactive" {
		return errors.New("invalid status")
	}

	if err != nil {
		return err
	}

	return nil
}

func NewPixKey(kind string, account *Account, key string) (*PixKey, error) {
	pixKey := PixKey{
		Kind:    kind,
		Key:     key,
		Account: account,
		Status:  "active",
	}

	pixKey.ID = uuid.NewV4().String()
	pixKey.CreatedAt = time.Now()

	err := pixKey.isValid()
	if err != nil {
		return nil, err
	}

	return &pixKey, nil
}
