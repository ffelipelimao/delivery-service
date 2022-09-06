package presentation

import "io"

type HttpRequest struct {
	Params map[string]string
	Query  map[string][]string
	Body   io.Reader
}

type HttpResponse struct {
	StatusCode int
	Body       any
}

type Error struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Cause   string `json:"cause"`
}

func ok(body interface{}) HttpResponse {
	return HttpResponse{StatusCode: 200, Body: body}
}

func badRequest(message string, err string) HttpResponse {
	return HttpResponse{StatusCode: 400, Body: Error{
		Message: message,
		Status:  400,
		Cause:   err,
	}}
}

func internalServerError(message string, err string) HttpResponse {
	return HttpResponse{StatusCode: 500, Body: Error{
		Message: message,
		Status:  500,
		Cause:   err,
	}}
}
