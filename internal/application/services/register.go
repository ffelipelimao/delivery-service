package services

import (
	"context"

	"github.com/ffelipelimao/delivery-service/internal/application/repository"
	"github.com/ffelipelimao/delivery-service/internal/domain"
	"github.com/google/uuid"
)

type RegisterService struct {
	RegisterRepository repository.RegisterRepository
	ObjectRepository   repository.ObjectRepository
}

func NewRegisterService(registerRepo repository.RegisterRepository, objectRepo repository.ObjectRepository) Service {
	return RegisterService{
		RegisterRepository: registerRepo,
		ObjectRepository:   objectRepo,
	}
}

func (s RegisterService) Create(ctx context.Context, register domain.Register) (domain.Register, error) {
	register.ID = uuid.NewString()

	for _, object := range register.Objects {
		object.ID = uuid.NewString()
		object.RegisterID = register.ID

		_, err := s.ObjectRepository.Insert(*object)
		if err != nil {
			return domain.Register{}, err
		}

	}

	_, err := s.RegisterRepository.Insert(register)
	if err != nil {
		return domain.Register{}, err
	}

	return register, nil
}

func (s RegisterService) Get(ctx context.Context, ID string) (domain.Register, error) {
	register, err := s.RegisterRepository.Get(ID)
	if err != nil {
		return domain.Register{}, err
	}

	return register, nil

}
