package service

import (
	"errors"
	"strconv"

	"github.com/aabdullahgungor/product-api/models"
	"github.com/aabdullahgungor/product-api/repository"
)

var (
	ErrIDIsNotValid    = errors.New("id is not valid")
	ErrNameIsNotEmpty = errors.New("Product title cannot be empty")
	ErrProductNotFound = errors.New("Product cannot be found")
)

type DefaultProductService struct {
	productRepo repository.IProductRepository
}

func NewDefaultProductService(pRepo repository.IProductRepository) *DefaultProductService {
	return &DefaultProductService{
		productRepo: pRepo,
	}
}

func (d *DefaultProductService) GetAll() ([]models.Product, error) {
	return d.productRepo.GetAllProducts()	
}

func (d *DefaultProductService) GetById(id string) (models.Product, error) {
	
	_ , err := strconv.Atoi(id)
	if err != nil {
		return models.Product{}, ErrIDIsNotValid
	}

	product, err := d.productRepo.GetProductById(id)

	if err != nil {
		return models.Product{}, repository.ErrProductNotFound
	}

	return product, nil
}

func (d *DefaultProductService) Create(product *models.Product) error {

	if product.Name == "" {
		return ErrNameIsNotEmpty
	}

	return d.productRepo.CreateProduct(product)
}

func (d *DefaultProductService) Edit(product *models.Product) error {
	
	if product.Id.String() == "" {
		return  ErrIDIsNotValid
	}
	if product.Name == "" {
		return ErrNameIsNotEmpty
	}

	err := d.productRepo.EditProduct(product)

	if err != nil {
		return  repository.ErrProductNotFound
	}

	return nil
}

func (d *DefaultProductService) Delete(id string) error {

	_ , err := strconv.Atoi(id)
	if err != nil {
		return  ErrIDIsNotValid
	}
	
	err = d.productRepo.DeleteProduct(id)

	if err != nil {
		return repository.ErrProductNotFound
	}

	return nil
}