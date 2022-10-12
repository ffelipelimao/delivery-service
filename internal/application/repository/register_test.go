package repository_test

import (
	"testing"

	"github.com/ffelipelimao/delivery-service/internal/application/repository"
	"github.com/ffelipelimao/delivery-service/internal/domain"
	"github.com/ffelipelimao/delivery-service/internal/framework/database"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewRegisterRepo(t *testing.T) {
	r := repository.NewRegisterRepository(nil)
	assert.NotNil(t, r)
}

func TestRegisterRepositoryInsert(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	r := domain.Register{
		ID:   uuid.NewString(),
		From: "test",
	}

	repo := repository.RegisterRepositoryDB{DB: db}
	_, err := repo.Insert(r)

	assert.Nil(t, err)
}

func TestRegisterRepositoryInsertErr(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	r := domain.Register{}

	repo := repository.RegisterRepositoryDB{DB: db}
	_, err := repo.Insert(r)

	assert.NotNil(t, err)
}

func TestRegisterRepositoryGet(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	ID := "123456"
	r := domain.Register{
		ID:   ID,
		From: "test",
	}

	repo := repository.RegisterRepositoryDB{DB: db}
	_, err := repo.Insert(r)
	assert.Nil(t, err)

	res, err := repo.Get(ID)
	assert.Nil(t, err)

	assert.Equal(t, "test", res.From)
}

func TestRegisterRepositoryGetErr(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	ID := "$-$-$-$"
	r := domain.Register{
		ID:   ID,
		From: "test",
	}

	repo := repository.RegisterRepositoryDB{DB: db}
	_, err := repo.Insert(r)
	assert.Nil(t, err)

	_, err = repo.Get(ID)
	assert.NotNil(t, err)
}
