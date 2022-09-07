package main

import (
	"github.com/ffelipelimao/delivery-service/internal/application/presentation"
	"github.com/gin-gonic/gin"
)

func AdaptRoute(controller presentation.Controller) func(c *gin.Context) {

	return func(c *gin.Context) {
		queries := c.Request.URL.Query()
		presentationParam := make(map[string]string)

		for _, param := range c.Params {
			presentationParam[param.Key] = param.Value
		}

		request := presentation.HttpRequest{Params: presentationParam, Query: queries, Body: c.Request.Body}
		response := controller.Handle(request)
		c.JSON(response.StatusCode, response.Body)
	}
}
