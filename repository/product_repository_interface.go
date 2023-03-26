package repository

import (
	"github.com/aabdullahgungor/product-api/models"
)

type IProductRepository interface {
	GetAll() ([]models.Product, error)
	GetById(id string) (models.Product, error)
	Create(product *models.Product) error
	Edit(product *models.Product) error
	Delete(id string) error
}