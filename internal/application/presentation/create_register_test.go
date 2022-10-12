package presentation_test

import (
	"bytes"
	"testing"

	"github.com/ffelipelimao/delivery-service/internal/application/helpers"
	"github.com/ffelipelimao/delivery-service/internal/application/presentation"
	mock_services "github.com/ffelipelimao/delivery-service/internal/application/services/mocks"
	"github.com/ffelipelimao/delivery-service/internal/domain"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestNewCreateRegister(t *testing.T) {
	ctrl := presentation.NewCreateRegisterController(nil)
	assert.NotNil(t, ctrl)
}

func TestCreateRegister_ok(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	registerFake := domain.Register{
		From: "test",
	}
	serviceMock := mock_services.NewMockService(ctrlMock)
	serviceMock.EXPECT().Create(gomock.Any(), registerFake).Return(registerFake, nil)

	controller := presentation.CreateRegisterController{
		RegisterService: serviceMock,
	}

	r := bytes.NewReader([]byte("{\"from\":\"test\"}"))
	request := helpers.HttpRequest{
		Body: r,
	}

	response := controller.Handle(request)
	assert.Equal(t, 200, response.StatusCode)
}

func TestCreateRegister_badRequest(t *testing.T) {
	controller := presentation.CreateRegisterController{}

	r := bytes.NewReader([]byte("{"))
	request := helpers.HttpRequest{
		Body: r,
	}

	response := controller.Handle(request)
	assert.Equal(t, 400, response.StatusCode)
}

func TestCreateRegister_internalServerError(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	registerFake := domain.Register{
		From: "test",
	}
	serviceMock := mock_services.NewMockService(ctrlMock)
	serviceMock.EXPECT().Create(gomock.Any(), registerFake).Return(domain.Register{}, assert.AnError)

	controller := presentation.CreateRegisterController{
		RegisterService: serviceMock,
	}

	r := bytes.NewReader([]byte("{\"from\":\"test\"}"))
	request := helpers.HttpRequest{
		Body: r,
	}

	response := controller.Handle(request)
	assert.Equal(t, 500, response.StatusCode)
}
