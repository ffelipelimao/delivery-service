package presentation_test

import (
	"testing"

	"github.com/ffelipelimao/delivery-service/internal/application/helpers"
	"github.com/ffelipelimao/delivery-service/internal/application/presentation"
	mock_services "github.com/ffelipelimao/delivery-service/internal/application/services/mocks"
	"github.com/ffelipelimao/delivery-service/internal/domain"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestNewGetRegister(t *testing.T) {
	ctrl := presentation.NewGetRegisterController(nil)
	assert.NotNil(t, ctrl)
}

func TestGetRegister_ok(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	registerFake := domain.Register{
		ID:   "123",
		From: "test",
	}
	serviceMock := mock_services.NewMockService(ctrlMock)
	serviceMock.EXPECT().Get(gomock.Any(), "123").Return(registerFake, nil)

	controller := presentation.GetRegisterController{
		RegisterService: serviceMock,
	}

	request := helpers.HttpRequest{
		Params: map[string]string{
			"id": "123",
		},
	}

	response := controller.Handle(request)
	assert.Equal(t, 200, response.StatusCode)
}

func Test_GetRegister_internalError(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	serviceMock := mock_services.NewMockService(ctrlMock)
	serviceMock.EXPECT().Get(gomock.Any(), "123").Return(domain.Register{}, assert.AnError)

	controller := presentation.GetRegisterController{
		RegisterService: serviceMock,
	}

	request := helpers.HttpRequest{
		Params: map[string]string{
			"id": "123",
		},
	}

	response := controller.Handle(request)
	assert.Equal(t, 500, response.StatusCode)
}
