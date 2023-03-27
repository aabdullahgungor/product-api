package controller

import (
	"errors"
	"net/http"

	"github.com/aabdullahgungor/product-api/models"
	"github.com/aabdullahgungor/product-api/service"
	"github.com/gin-gonic/gin"
)

type productController struct {
	service service.IProductService
}

func NewProductController(ps service.IProductService) *productController {
	return &productController{service: ps}
}

func (ps *productController) GetAllProducts(c *gin.Context) {
	products, err := ps.service.GetAll()
	if err != nil {
		c.IndentedJSON(http.StatusNotFound,gin.H{"error": "Products cannot show: " + err.Error(), })
		return
	}
	c.Header("Content-Type", "application/json")
	c.IndentedJSON(http.StatusOK, products)
}

func (ps *productController) GetProductById(c *gin.Context) { 
	str_id := c.Param("id")
	product, err := ps.service.GetById(str_id)
	if err != nil {
		if errors.Is(err, service.ErrIDIsNotValid) {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "id is not valid"+err.Error()})
			return
		} else if  errors.Is(err, service.ErrProductNotFound) {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Product cannot be found"+err.Error()})
			return
		}
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	} 
	c.Header("Content-Type", "application/json")
	c.IndentedJSON(http.StatusOK, product)
}


func (ps *productController) CreateProduct(c *gin.Context) { 
	var product models.Product
	err := c.ShouldBindJSON(&product)

	if err != nil {
		c.IndentedJSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = ps.service.Create(&product)

	if err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, gin.H{
			"error": "cannot create product: " + err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message":"Product has been created","product_id": product.Id})

}		

func (ps *productController) EditProduct(c *gin.Context) { 
	var product models.Product
	err := c.ShouldBindJSON(&product)

	if err != nil {
		c.IndentedJSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = ps.service.Edit(&product)

	if err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, gin.H{
			"error": "cannot edit product: " + err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message":"Product has been edited","product_id": product.Id})

}		

func (ps *productController) DeleteProduct(c *gin.Context) { 
	str_id := c.Param("id")
	err := ps.service.Delete(str_id)
	if err != nil {
		if errors.Is(err, service.ErrIDIsNotValid) {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "id is not valid"+err.Error()})
			return
		} else if  errors.Is(err, service.ErrProductNotFound) {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Product cannot be found"+err.Error()})
			return
		}
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}

	c.IndentedJSON(http.StatusAccepted, gin.H{"message":"Product has been deleted","product_id": str_id})	

}		