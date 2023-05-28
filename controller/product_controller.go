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

// GetProducts            godoc
// @Summary		Get products array
// @Description	Responds with the list of all products as JSON.
// @Tags			products
// @Produce		json
// @Success		200	{object}	models.Product
// @Router			/products [get]
func (ps *productController) GetAllProducts(c *gin.Context) {
	products, err := ps.service.GetAll()
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Products cannot show: " + err.Error()})
		return
	}
	c.Header("Content-Type", "application/json")
	c.IndentedJSON(http.StatusOK, products)
}

// GetProduct          godoc
// @Summary		Get single product by id
// @Description	Returns the product whose id value matches the id.
// @Tags			products
// @Produce		json
// @Param			id path	string true "search product by id"
// @Success		200		{object}	models.Product
// @Failure 	400     error message
// @Failure 	406     error message
// @Router			/products/{id} [get]
func (ps *productController) GetProductById(c *gin.Context) {
	str_id := c.Param("id")
	product, err := ps.service.GetById(str_id)
	if err != nil {
		if errors.Is(err, service.ErrIDIsNotValid) {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		} else if errors.Is(err, service.ErrProductNotFound) {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Header("Content-Type", "application/json")
	c.IndentedJSON(http.StatusOK, product)
}

// CreateProduct          godoc
// @Summary		Add a new product
// @Description	Takes a product JSON and store in DB. Return saved JSON.
// @Tags			products
// @Produce		json
// @Param			product body	models.Product	true "Product JSON"
// @Success		200		{object}  models.Product
// @Failure 	400     error message
// @Failure 	406     error message
// @Router			/products [post]
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

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Product has been created"})
}

// EditProduct          godoc
// @Summary		Edit an product
// @Description	Takes a product JSON and edit an in DB. Return saved JSON.
// @Tags			products
// @Produce		json
// @Param			product body	models.Product	true "Product JSON"
// @Success		200		{object}  models.Product
// @Failure 	400     error message
// @Failure 	406     error message
// @Router			/products [put]
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

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Product has been edited", "product_id": product.Id})
}

// DeleteProduct          godoc
// @Summary		Delete an product
// @Description	Remove an product from DB by id.
// @Tags			products
// @Produce		json
// @Param			id path	string true "delete product by id"
// @Success		200		{object} models.Product
// @Failure 	400     error message
// @Failure 	406     error message
// @Router			/products/{id} [delete]
func (ps *productController) DeleteProduct(c *gin.Context) {
	str_id := c.Param("id")
	err := ps.service.Delete(str_id)
	if err != nil {
		if errors.Is(err, service.ErrIDIsNotValid) {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": "id is not valid" + err.Error()})
			return
		} else if errors.Is(err, service.ErrProductNotFound) {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Product cannot be found" + err.Error()})
			return
		}
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusAccepted, gin.H{"message": "Product has been deleted", "product_id": str_id})
}
