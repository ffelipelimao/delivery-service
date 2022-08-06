package presentation

type DeliveryController struct {
}

func NewDeliveryController() Controller {
	return &DeliveryController{}
}

func (s *DeliveryController) Handle(req HttpRequest) HttpResponse {
	var e interface{}
	return ok(e)
}
