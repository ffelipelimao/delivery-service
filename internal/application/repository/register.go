package repository

import (
	"github.com/ffelipelimao/delivery-service/internal/domain"
	"github.com/jinzhu/gorm"
)

//go:generate mockgen -destination=./mocks/register.go -source=./register.go
type RegisterRepository interface {
	Insert(register domain.Register) (domain.Register, error)
	Get(ID string) (domain.Register, error)
}

type RegisterRepositoryDB struct {
	DB *gorm.DB
}

func NewRegisterRepository(DB *gorm.DB) RegisterRepository {
	return RegisterRepositoryDB{
		DB: DB,
	}
}

func (repo RegisterRepositoryDB) Insert(register domain.Register) (domain.Register, error) {

	err := repo.DB.Create(register).Error
	if err != nil {
		return domain.Register{}, err
	}

	return register, nil
}

func (repo RegisterRepositoryDB) Get(ID string) (domain.Register, error) {
	var register domain.Register

	err := repo.DB.First(&register, ID).Error
	if err != nil {
		return domain.Register{}, err
	}

	return register, nil
}
