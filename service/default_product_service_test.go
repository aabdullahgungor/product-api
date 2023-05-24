package service

import (
	"testing"

	"github.com/aabdullahgungor/product-api/repository"
	"github.com/golang/mock/gomock"
)

func TestDefaultProductService_GetAll(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepository := repository.NewMockIProductRepository(mockCtrl)
	mockRepository.EXPECT().GetAllProducts().Return([]model.Product{}, nil)
}

func TestDefaultProductService_GetById(t *testing.T) {
}

func TestDefaultProductService_Create(t *testing.T) {
}

func TestDefaultProductService_Edit(t *testing.T) {
}

func TestDefaultProductService_Delete(t *testing.T) {
}
