package main

import (
	"github.com/ffelipelimao/delivery-service/internal/application/presentation"
	"github.com/ffelipelimao/delivery-service/internal/framework/server/adapter"
	"github.com/gin-gonic/gin"
)

func Setup(g *gin.Engine) {

	router := g.Group("/api/v1/delivery")

	RegisterGetController := presentation.NewGetRegisterController()
	router.GET("/Register", adapter.AdaptRoute(RegisterGetController))
}
