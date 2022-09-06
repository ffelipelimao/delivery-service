package presentation

type GetRegisterController struct {
}

func NewGetRegisterController() Controller {
	return &GetRegisterController{}
}

func (s *GetRegisterController) Handle(req HttpRequest) HttpResponse {
	return ok(req.Params)
}
