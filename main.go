package main

import (
	"cleancode/controller"
	"cleancode/middleware"

	_ "cleancode/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title 		Office Reservation API
// @version		1.0
// @description API to calculate monthly revenue and display office reservations.
// @BaseUrl  	/
func main() {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.POST("/calculate", controller.CalculateHandler)
	router.GET("/manual", middleware.ValidateAuth(), controller.ManualHandler)
	router.Run(":8080")
}
