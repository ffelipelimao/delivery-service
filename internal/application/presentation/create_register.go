package presentation

import (
	"context"
	"encoding/json"
	"io/ioutil"

	"github.com/ffelipelimao/delivery-service/internal/application/helpers"
	"github.com/ffelipelimao/delivery-service/internal/application/services"
	"github.com/ffelipelimao/delivery-service/internal/domain"
)

type CreteRegisterController struct {
	registerService services.Service
}

func NewCreateRegisterController(service services.Service) Controller {
	return &CreteRegisterController{
		registerService: service,
	}
}

func (s *CreteRegisterController) Handle(req helpers.HttpRequest) helpers.HttpResponse {
	ctx := context.TODO()
	var register domain.Register

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return helpers.BadRequest("invalid input", err.Error())
	}
	err = json.Unmarshal(body, &register)
	if err != nil {
		return helpers.BadRequest("invalid input", err.Error())
	}

	register, err = s.registerService.Create(ctx, register)
	if err != nil {
		return helpers.InternalServerError("internal server error", err.Error())
	}

	return helpers.Ok(register)
}
