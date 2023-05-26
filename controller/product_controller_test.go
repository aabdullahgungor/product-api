package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/aabdullahgungor/product-api/models"
	"github.com/aabdullahgungor/product-api/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCarController_GetAllProducts(t *testing.T) {

	t.Run("Error", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := service.NewMockIProductService(mockCtrl)
		mockService.EXPECT().GetAll().Return([]models.Product{}, errors.New("hata!")).AnyTimes()

		w := httptest.NewRecorder()
		gin.SetMode(gin.ReleaseMode)
		ctx, r := gin.CreateTestContext(w)
		productTestController := NewProductController(mockService)
		productTestController.GetAllProducts(ctx)

		req, _ := http.NewRequest("GET", "api/v1/products", nil)
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusNotFound, w.Code)
		fmt.Println(w.Code)
	})

	t.Run("Success", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := service.NewMockIProductService(mockCtrl)
		mockService.EXPECT().GetAll().Return([]models.Product{}, nil).AnyTimes()

		w := httptest.NewRecorder()
		gin.SetMode(gin.ReleaseMode)
		ctx, r := gin.CreateTestContext(w)
		productTestController := NewProductController(mockService)
		productTestController.GetAllProducts(ctx)

		req, _ := http.NewRequest("GET", "api/v1/products", nil)
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		fmt.Println(w.Code)

	})
}

func TestCarController_GetProductbyId(t *testing.T) {
	t.Run("Error", func(t *testing.T) {
		var id string
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := service.NewMockIProductService(mockCtrl)
		mockService.EXPECT().GetById(id).Return(models.Product{}, service.ErrProductNotFound).AnyTimes()

		w := httptest.NewRecorder()
		gin.SetMode(gin.ReleaseMode)
		ctx, r := gin.CreateTestContext(w)
		productTestController := NewProductController(mockService)
		productTestController.GetProductById(ctx)

		req, _ := http.NewRequest("GET", "api/v1/products/:id", nil)
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusNotFound, w.Code)
		fmt.Println(w.Code)
	})

	t.Run("Success", func(t *testing.T) {
		var id string
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := service.NewMockIProductService(mockCtrl)
		mockService.EXPECT().GetById(id).Return(models.Product{
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
		}, nil).AnyTimes()

		w := httptest.NewRecorder()
		gin.SetMode(gin.ReleaseMode)
		ctx, r := gin.CreateTestContext(w)
		productTestController := NewProductController(mockService)
		productTestController.GetProductById(ctx)

		req, _ := http.NewRequest("GET", "api/v1/products/:id", nil)
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		fmt.Println(w.Code)

		var responseData models.Product
		json.NewDecoder(w.Body).Decode(&responseData)
		assert.Equal(t, "Mobile 1", responseData.Name)
		t.Log("\nProduct name is: ", responseData.Name)

	})
}

func TestCarController_CreateProduct(t *testing.T) {
	t.Run("ErrorCreate", func(t *testing.T) {
		product := models.Product{
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
		}
		jsonValue, _ := json.Marshal(product)
		byteProduct := bytes.NewBuffer(jsonValue)

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := service.NewMockIProductService(mockCtrl)
		mockService.EXPECT().Create(&product).Return(errors.New("hata")).AnyTimes()

		w := httptest.NewRecorder()
		gin.SetMode(gin.ReleaseMode)
		ctx, r := gin.CreateTestContext(w)
		productTestController := NewProductController(mockService)
		productTestController.CreateProduct(ctx)
		req, err := http.NewRequest("POST", "api/v1/products", byteProduct)
		if err != nil {
			fmt.Println(err)
		}
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusNotAcceptable, w.Code)
		t.Log(w.Body.String())
	})

	t.Run("Success", func(t *testing.T) {
		product := models.Product{
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
		}
		jsonValue, _ := json.Marshal(product)
		byteProduct := bytes.NewBuffer(jsonValue)

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := service.NewMockIProductService(mockCtrl)
		mockService.EXPECT().Create(&product).Return(nil).AnyTimes()

		w := httptest.NewRecorder()
		gin.SetMode(gin.ReleaseMode)
		ctx, r := gin.CreateTestContext(w)
		productTestController := NewProductController(mockService)
		productTestController.CreateProduct(ctx)
		req, err := http.NewRequest("POST", "api/v1/products", byteProduct)
		if err != nil {
			fmt.Println(err)
		}
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusCreated, w.Code)
		t.Log(w.Body.String())

	})
}

func TestCarController_EditProduct(t *testing.T) {
	t.Run("ErrorEdit", func(t *testing.T) {
		product := models.Product{
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
		}
		jsonValue, _ := json.Marshal(product)
		byteProduct := bytes.NewBuffer(jsonValue)

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := service.NewMockIProductService(mockCtrl)
		mockService.EXPECT().Edit(&product).Return(errors.New("hata")).AnyTimes()

		w := httptest.NewRecorder()
		gin.SetMode(gin.ReleaseMode)
		ctx, r := gin.CreateTestContext(w)
		productTestController := NewProductController(mockService)
		productTestController.EditProduct(ctx)
		req, err := http.NewRequest("PUT", "api/v1/products", byteProduct)
		if err != nil {
			fmt.Println(err)
		}
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusNotAcceptable, w.Code)
		t.Log(w.Body.String())
	})

	t.Run("Success", func(t *testing.T) {
		product := models.Product{
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
		}
		jsonValue, _ := json.Marshal(product)
		byteProduct := bytes.NewBuffer(jsonValue)

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := service.NewMockIProductService(mockCtrl)
		mockService.EXPECT().Edit(&product).Return(nil).AnyTimes()

		w := httptest.NewRecorder()
		gin.SetMode(gin.ReleaseMode)
		ctx, r := gin.CreateTestContext(w)
		productTestController := NewProductController(mockService)
		productTestController.EditProduct(ctx)
		req, err := http.NewRequest("PUT", "api/v1/products", byteProduct)
		if err != nil {
			fmt.Println(err)
		}
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusCreated, w.Code)
		t.Log(w.Body.String())
	})
}

func TestCarController_DeleteProduct(t *testing.T) {
	t.Run("Error", func(t *testing.T) {
		var id string
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := service.NewMockIProductService(mockCtrl)
		mockService.EXPECT().Delete(id).Return(service.ErrProductNotFound).AnyTimes()

		w := httptest.NewRecorder()
		gin.SetMode(gin.ReleaseMode)
		ctx, _ := gin.CreateTestContext(w)
		productTestController := NewProductController(mockService)
		productTestController.DeleteProduct(ctx)

		assert.Equal(t, http.StatusNotFound, w.Code)
		t.Log(w.Body.String())
	})

	t.Run("Success", func(t *testing.T) {
		var id string
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := service.NewMockIProductService(mockCtrl)
		mockService.EXPECT().Delete(id).Return(nil).AnyTimes()

		w := httptest.NewRecorder()
		gin.SetMode(gin.ReleaseMode)
		ctx, _ := gin.CreateTestContext(w)
		productTestController := NewProductController(mockService)
		productTestController.DeleteProduct(ctx)

		assert.Equal(t, http.StatusAccepted, w.Code)
		t.Log(w.Body.String())
	})
}
