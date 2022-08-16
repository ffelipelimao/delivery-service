package presentation

type Controller interface {
	Handle(req HttpRequest) HttpResponse
}
