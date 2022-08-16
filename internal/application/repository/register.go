package repository

import (
	"github.com/ffelipelimao/delivery-service/internal/domain"
	"github.com/jinzhu/gorm"
)

type RegisterRepository interface {
	Insert(register domain.Register) (domain.Register, error)
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
