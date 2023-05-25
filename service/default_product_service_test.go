package service

import (
	"testing"
	"time"

	"github.com/aabdullahgungor/product-api/models"
	"github.com/aabdullahgungor/product-api/repository"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestDefaultProductService_GetAll(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepository := repository.NewMockIProductRepository(mockCtrl)
	mockRepository.EXPECT().GetAllProducts().Return([]models.Product{{
		Id:         primitive.ObjectID{'1'},
		Name:       "Mobile 1",
		Price:      40,
		Quantity:   4,
		Status:     true,
		Date:       time.Time{},
		CategoryId: primitive.ObjectID{'1'},
		Brand: models.Brand{
			Id:   primitive.ObjectID{'1'},
			Name: "Brand 1",
		},
		Colors: []string{"red", "green", "blue"},
	},
		{
			Id:         primitive.ObjectID{'2'},
			Name:       "Mobile 2",
			Price:      50,
			Quantity:   5,
			Status:     true,
			Date:       time.Time{},
			CategoryId: primitive.ObjectID{'1'},
			Brand: models.Brand{
				Id:   primitive.ObjectID{'2'},
				Name: "Brand 2",
			},
			Colors: []string{"red", "green", "blue"},
		}}, nil)

	productService := NewDefaultProductService(mockRepository)
	products, err := productService.GetAll()

	if assert.Nil(t, err) {
		if len(products) == 2 {
			t.Log("Product counts is matching, func run succesfuly")
		} else {
			t.Log("Product counts not matching, there is a problem in func")
		}
	} else {
		t.Log(err)
	}
}

func TestDefaultProductService_GetById(t *testing.T) {
	id := "1"
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepository := repository.NewMockIProductRepository(mockCtrl)
	mockRepository.EXPECT().GetProductById(gomock.Eq(id)).Return(models.Product{}, ErrProductNotFound)

	productService := NewDefaultProductService(mockRepository)
	_, err := productService.GetById(id)

	assert.ErrorIs(t, err, ErrProductNotFound)
}

func TestDefaultProductService_Create(t *testing.T) {

	product := models.Product{Name: "Mobile 1"}
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepository := repository.NewMockIProductRepository(mockCtrl)
	mockRepository.EXPECT().CreateProduct(&product).Return(nil).Times(1)

	productService := NewDefaultProductService(mockRepository)
	err := productService.Create(&product)

	if assert.Nil(t, err) {
		t.Log("Success Create Product")
	} else {
		t.Log("Product cannot create")
	}

}

func TestDefaultProductService_Edit(t *testing.T) {

	product := models.Product{Name: "Mobile 1"}
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepository := repository.NewMockIProductRepository(mockCtrl)
	mockRepository.EXPECT().EditProduct(&product).Return(nil).Times(1)

	productService := NewDefaultProductService(mockRepository)
	err := productService.Edit(&product)

	if assert.Nil(t, err) {
		t.Log("Success Update Product")
	} else {
		t.Log("Product cannot update")
	}
}

func TestDefaultProductService_Delete(t *testing.T) {

	id := "1"
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepository := repository.NewMockIProductRepository(mockCtrl)
	mockRepository.EXPECT().DeleteProduct(gomock.Eq(id)).Return(nil).Times(1)

	productService := NewDefaultProductService(mockRepository)
	err := productService.Delete(id)

	if assert.Nil(t, err) {
		t.Log("Success delete Product")
	} else {
		t.Log("Product cannot delete")
	}
}
