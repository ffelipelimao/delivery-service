package main

import (
	"github.com/ffelipelimao/delivery-service/internal/application/presentation"
	"github.com/ffelipelimao/delivery-service/internal/application/repository"
	"github.com/ffelipelimao/delivery-service/internal/application/services"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

const (
	registryByID = "/register/:id"
	registry     = "/register"
)

func Setup(g *gin.Engine, db *gorm.DB) {

	router := g.Group("/v1/delivery")

	registerGetController := presentation.NewGetRegisterController()
	router.GET(registryByID, AdaptRoute(registerGetController))

	objectRepository := repository.NewObjectRepository(db)
	registerRepository := repository.NewRegisterRepository(db)
	registerCreteService := services.NewRegisterService(registerRepository, objectRepository)
	registerCreateController := presentation.NewCreateRegisterController(registerCreteService)
	router.POST(registry, AdaptRoute(registerCreateController))
}
