package repository

import (
	"github.com/aabdullahgungor/product-api/models"
)

type IProductRepository interface {
	GetAllProducts() ([]models.Product, error)
	GetProductById(id string) (models.Product, error)
	CreateProduct(product *models.Product) error
	EditProduct(product *models.Product) error
	DeleteProduct(id string) error
}