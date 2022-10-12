package repository_test

import (
	"testing"

	"github.com/ffelipelimao/delivery-service/internal/application/repository"
	"github.com/ffelipelimao/delivery-service/internal/domain"
	"github.com/ffelipelimao/delivery-service/internal/framework/database"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewObjectRepo(t *testing.T) {
	r := repository.NewObjectRepository(nil)
	assert.NotNil(t, r)
}

func TestObjectRepositoryInsert(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	obj := domain.Object{
		ID:         uuid.NewString(),
		Name:       "teste",
		RegisterID: uuid.NewString(),
	}

	repo := repository.ObjectRepositoryDB{DB: db}
	_, err := repo.Insert(obj)

	assert.Nil(t, err)
}

func TestObjectRepositoryInsertErr(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	obj := domain.Object{}

	repo := repository.ObjectRepositoryDB{DB: db}
	_, err := repo.Insert(obj)

	assert.NotNil(t, err)
}
