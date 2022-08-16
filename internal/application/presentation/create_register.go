package presentation

import (
	"context"
	"encoding/json"
	"io/ioutil"

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

func (s *CreteRegisterController) Handle(req HttpRequest) HttpResponse {
	ctx := context.TODO()
	var register domain.Register

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return badRequest("invalid input", err.Error())
	}
	err = json.Unmarshal(body, &register)
	if err != nil {
		return badRequest("invalid input", err.Error())
	}

	register, err = s.registerService.CreateRegister(ctx, register)
	if err != nil {
		return internalServerError("internal server error", err.Error())
	}

	return ok(register)
}
