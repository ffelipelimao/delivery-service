package adapter

import (
	"github.com/ffelipelimao/delivery-service/internal/application/presentation"
	"github.com/gin-gonic/gin"
)

func AdaptRoute(controller presentation.Controller) func(c *gin.Context) {

	return func(c *gin.Context) {
		//TODO: Refactor adapter to handle with errors, query params, paths...

		request := presentation.HttpRequest{Body: c.Request.Body}
		response := controller.Handle(request)
		c.JSON(response.StatusCode, response.Body)
	}
}
