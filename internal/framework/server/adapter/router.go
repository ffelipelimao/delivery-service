package adapter

import (
	"encoding/json"

	"github.com/ffelipelimao/delivery-service/internal/application/presentation"
	"github.com/gin-gonic/gin"
)

func AdaptRoute(controller presentation.Controller) func(c *gin.Context) {

	return func(c *gin.Context) {
		request := presentation.HttpRequest{Body: c.Request.Body}
		response := controller.Handle(request)

		c.Header("Content-Type", "application/json")
		c.Writer.WriteHeader(response.StatusCode)
		responseData, _ := json.Marshal(response.Body)
		c.Writer.Write(responseData)
	}
}
