package router

import (
	"github.com/aabdullahgungor/product-api/controller"
	"github.com/aabdullahgungor/product-api/docs"
	"github.com/aabdullahgungor/product-api/repository"
	"github.com/aabdullahgungor/product-api/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {

	productRepo := repository.NewMongoDbProductRepository()
	productService := service.NewDefaultProductService(productRepo)
	productController := controller.NewProductController(productService)

	main := router.Group("api/v1")
	{
		products := main.Group("products")
		{
			products.GET("/", productController.GetAllProducts)
			products.GET("/:id", productController.GetProductById)
			products.POST("/", productController.CreateProduct)
			products.PUT("/", productController.EditProduct)
			products.DELETE("/:id", productController.DeleteProduct)
		}

		// programmatically set swagger info
		docs.SwaggerInfo.Title = "Product API"
		docs.SwaggerInfo.Description = "This is a sample server ProductApi server."
		docs.SwaggerInfo.Version = "1.0"
		docs.SwaggerInfo.Host = ""
		docs.SwaggerInfo.BasePath = "/api/v1"
		docs.SwaggerInfo.Schemes = []string{"http", "https"}
		// add swagger
		main.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	return router
}
