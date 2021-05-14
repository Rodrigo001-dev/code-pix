package usecase

import (
	"errors"

	"github.com/Rodrigo001-de/code-pix/domain/model"
)

type PixUseCase struct {
	PixKeyRepository model.PixKeyRepositoryInterface
}

func (p *PixUseCase) RegisterKey(key string, kind string, accountId string) (*model.PixKey, error) {
	account, err := p.PixKeyRepository.FindAccount(accountId)
	if err != nil {
		return nil, err
	}

	// se a conta existir eu vou criar uma chave
	pixKey, err := model.NewPixKey(kind, account, key)
	// se não validar, vai retornar um erro
	if err != nil {
		return nil, err
	}

	p.PixKeyRepository.RegisterKey(pixKey)
	if pixKey.ID == "" {
		return nil, errors.New("unable to create new key at the moment")
	}

	return pixKey, nil
}

// o FindKey vai receber uma key e uma kind e vai retornar uma KeyPix, ele vai
// siplesmente chamar o repositório FindKeyByKind e vai retornar o pixKey
func (p *PixUseCase) FindKey(key string, kind string) (*model.PixKey, error) {
	pixKey, err := p.PixKeyRepository.FindKeyByKind(key, kind)
	if err != nil {
		return nil, err
	}

	return pixKey, nil
}
