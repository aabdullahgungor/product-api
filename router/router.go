package router

import (
	"github.com/aabdullahgungor/product-api/controller"
	"github.com/aabdullahgungor/product-api/repository"
	"github.com/aabdullahgungor/product-api/service"
	"github.com/gin-gonic/gin"
)



func ConfigRoutes(router *gin.Engine) *gin.Engine{

	productRepo := repository.NewMongoDbProductRepository()
	productService := service.NewDefaultProductService(productRepo)
	productController := controller.NewProductController(productService)

	main := router.Group("api/v1") 
	{
		products := main.Group("products")
		{
			products.GET("/",  productController.GetAllProducts)
			products.GET("/:id", productController.GetProductById)
			products.POST("/", productController.CreateProduct)
			products.PUT("/", productController.EditProduct)
			products.DELETE("/:id", productController.DeleteProduct)
		}
	}
	return router
}