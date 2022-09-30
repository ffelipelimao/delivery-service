package presentation

import (
	"context"
	"encoding/json"
	"io/ioutil"

	"github.com/ffelipelimao/delivery-service/internal/application/helpers"
	"github.com/ffelipelimao/delivery-service/internal/application/services"
	"github.com/ffelipelimao/delivery-service/internal/domain"
)

type CreateRegisterController struct {
	RegisterService services.Service
}

func NewCreateRegisterController(service services.Service) Controller {
	return &CreateRegisterController{
		RegisterService: service,
	}
}

func (s *CreateRegisterController) Handle(req helpers.HttpRequest) helpers.HttpResponse {
	ctx := context.TODO()
	var register domain.Register

	body, _ := ioutil.ReadAll(req.Body)
	err := json.Unmarshal(body, &register)
	if err != nil {
		return helpers.BadRequest("invalid input", err.Error())
	}

	register, err = s.RegisterService.Create(ctx, register)
	if err != nil {
		return helpers.InternalServerError("internal server error", err.Error())
	}

	return helpers.Ok(register)
}
