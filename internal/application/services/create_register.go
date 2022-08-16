package services

import (
	"context"

	"github.com/ffelipelimao/delivery-service/internal/application/repository"
	"github.com/ffelipelimao/delivery-service/internal/domain"
	"github.com/google/uuid"
)

type RegisterService struct {
	registerRepository repository.RegisterRepository
	objectRepository   repository.ObjectRepository
}

func NewRegisterService(registerRepo repository.RegisterRepository, objectRepo repository.ObjectRepository) Service {
	return RegisterService{
		registerRepository: registerRepo,
		objectRepository:   objectRepo,
	}
}

func (s RegisterService) CreateRegister(ctx context.Context, register domain.Register) (domain.Register, error) {

	register.ID = uuid.NewString()

	for _, object := range register.Objects {
		object.ID = uuid.NewString()
		object.RegisterID = register.ID

		_, err := s.objectRepository.Insert(*object)
		if err != nil {
			return domain.Register{}, err
		}

	}

	_, err := s.registerRepository.Insert(register)
	if err != nil {
		return domain.Register{}, err
	}

	return register, nil
}
