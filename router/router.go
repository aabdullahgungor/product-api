package router

import (
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine{
	main := router.Group("api/v1") 
	{
		products := main.Group("products")
		{
			products.GET("/", )
			products.GET("/:id",)
			products.POST("/", )
			products.PUT("/", )
			products.DELETE("/:id", )
		}
	}
	return router
}