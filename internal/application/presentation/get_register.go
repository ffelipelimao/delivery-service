package presentation

import (
	"context"

	"github.com/ffelipelimao/delivery-service/internal/application/helpers"
	"github.com/ffelipelimao/delivery-service/internal/application/services"
)

type GetRegisterController struct {
	RegisterService services.Service
}

func NewGetRegisterController(service services.Service) Controller {
	return &GetRegisterController{
		RegisterService: service,
	}
}

func (s *GetRegisterController) Handle(req helpers.HttpRequest) helpers.HttpResponse {
	ctx := context.TODO()

	ID := req.Params["id"]

	register, err := s.RegisterService.Get(ctx, ID)
	if err != nil {
		return helpers.InternalServerError("internal server error", err.Error())
	}

	return helpers.Ok(register)
}
