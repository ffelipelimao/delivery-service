package services_test

import (
	"context"
	"testing"

	mock_repository "github.com/ffelipelimao/delivery-service/internal/application/repository/mocks"
	"github.com/ffelipelimao/delivery-service/internal/application/services"
	"github.com/ffelipelimao/delivery-service/internal/domain"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestNewService(t *testing.T) {
	svc := services.NewRegisterService(nil, nil)
	assert.NotNil(t, svc)
}

func TestServiceCreateOk(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	objectRepoMock := mock_repository.NewMockObjectRepository(ctrlMock)
	objectRepoMock.EXPECT().Insert(gomock.Any()).Return(domain.Object{}, nil)

	registerRepoMock := mock_repository.NewMockRegisterRepository(ctrlMock)
	registerRepoMock.EXPECT().Insert(gomock.Any()).Return(domain.Register{}, nil)

	registerFake := domain.Register{
		From: "test",
		Objects: []*domain.Object{
			{
				Name: "test_obj",
			},
		},
	}

	svc := services.RegisterService{
		RegisterRepository: registerRepoMock,
		ObjectRepository:   objectRepoMock,
	}

	_, err := svc.Create(context.TODO(), registerFake)
	assert.Nil(t, err)
}

func TestServiceCreateObjRepoErr(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	objectRepoMock := mock_repository.NewMockObjectRepository(ctrlMock)
	objectRepoMock.EXPECT().Insert(gomock.Any()).Return(domain.Object{}, assert.AnError)

	registerFake := domain.Register{
		From: "test",
		Objects: []*domain.Object{
			{
				Name: "test_obj",
			},
		},
	}

	svc := services.RegisterService{
		ObjectRepository: objectRepoMock,
	}

	_, err := svc.Create(context.TODO(), registerFake)
	assert.NotNil(t, err)
}

func TestServiceCreateRegisterRepoErr(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	objectRepoMock := mock_repository.NewMockObjectRepository(ctrlMock)
	objectRepoMock.EXPECT().Insert(gomock.Any()).Return(domain.Object{}, nil)

	registerRepoMock := mock_repository.NewMockRegisterRepository(ctrlMock)
	registerRepoMock.EXPECT().Insert(gomock.Any()).Return(domain.Register{}, assert.AnError)

	registerFake := domain.Register{
		From: "test",
		Objects: []*domain.Object{
			{
				Name: "test_obj",
			},
		},
	}

	svc := services.RegisterService{
		RegisterRepository: registerRepoMock,
		ObjectRepository:   objectRepoMock,
	}

	_, err := svc.Create(context.TODO(), registerFake)
	assert.NotNil(t, err)
}

func TestServiceGetOk(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	registerFake := domain.Register{
		ID:   "1234",
		From: "test",
		Objects: []*domain.Object{
			{
				Name: "test_obj",
			},
		},
	}

	registerRepoMock := mock_repository.NewMockRegisterRepository(ctrlMock)
	registerRepoMock.EXPECT().Get("1234").Return(registerFake, nil)

	svc := services.RegisterService{
		RegisterRepository: registerRepoMock,
	}

	res, err := svc.Get(context.TODO(), "1234")
	assert.Nil(t, err)

	assert.Equal(t, "test", res.From)
	assert.Equal(t, "1234", res.ID)
}

func TestServiceGetRepoErr(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	registerRepoMock := mock_repository.NewMockRegisterRepository(ctrlMock)
	registerRepoMock.EXPECT().Get("1234").Return(domain.Register{}, assert.AnError)

	svc := services.RegisterService{
		RegisterRepository: registerRepoMock,
	}

	_, err := svc.Get(context.TODO(), "1234")
	assert.NotNil(t, err)
}
