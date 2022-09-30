package presentation

import "github.com/ffelipelimao/delivery-service/internal/application/helpers"

type Controller interface {
	Handle(req helpers.HttpRequest) helpers.HttpResponse
}
