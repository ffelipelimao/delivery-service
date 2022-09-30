package server

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

func setup(g *gin.Engine, db *gorm.DB) {

	router := g.Group("/v1/delivery")
	objectRepository := repository.NewObjectRepository(db)
	registerRepository := repository.NewRegisterRepository(db)
	registerService := services.NewRegisterService(registerRepository, objectRepository)

	registerGetController := presentation.NewGetRegisterController(registerService)
	router.GET(registryByID, adaptRoute(registerGetController))

	registerCreateController := presentation.NewCreateRegisterController(registerService)
	router.POST(registry, adaptRoute(registerCreateController))
}
