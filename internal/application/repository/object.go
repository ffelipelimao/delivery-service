package repository

import (
	"github.com/ffelipelimao/delivery-service/internal/domain"
	"github.com/jinzhu/gorm"
)

type ObjectRepository interface {
	Insert(object domain.Object) (domain.Object, error)
}

type ObjectRepositoryDB struct {
	DB *gorm.DB
}

func NewObjectRepository(DB *gorm.DB) ObjectRepository {
	return ObjectRepositoryDB{
		DB: DB,
	}
}

func (repo ObjectRepositoryDB) Insert(object domain.Object) (domain.Object, error) {

	err := repo.DB.Create(object).Error
	if err != nil {
		return domain.Object{}, err
	}

	return object, nil
}
