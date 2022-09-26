package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Start(dbConnection *gorm.DB) {
	app := gin.Default()

	setup(app, dbConnection)
	app.Run(":8080")
}
