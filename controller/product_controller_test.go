package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
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

		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = &http.Request{
			Header: make(http.Header),
			URL:    &url.URL{},
		}
		ctx.Request.Method = "GET"
		ctx.Request.Header.Set("Content-Type", "application/json")
		productTestController := NewProductController(mockService)
		productTestController.GetAllProducts(ctx)

		assert.Equal(t, http.StatusNotFound, w.Code)
		fmt.Println(w.Code)
	})

	t.Run("Success", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := service.NewMockIProductService(mockCtrl)
		mockService.EXPECT().GetAll().Return([]models.Product{}, nil).AnyTimes()

		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = &http.Request{
			Header: make(http.Header),
			URL:    &url.URL{},
		}
		ctx.Request.Method = "GET"
		ctx.Request.Header.Set("Content-Type", "application/json")
		productTestController := NewProductController(mockService)
		productTestController.GetAllProducts(ctx)

		assert.Equal(t, http.StatusOK, w.Code)
		fmt.Println(w.Code)

	})
}

func TestCarController_GetProductbyId(t *testing.T) {
	t.Run("Error", func(t *testing.T) {
		id := "1"
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := service.NewMockIProductService(mockCtrl)
		mockService.EXPECT().GetById(id).Return(models.Product{}, service.ErrProductNotFound).AnyTimes()

		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = &http.Request{
			Header: make(http.Header),
			URL:    &url.URL{},
		}
		ctx.Request.Method = "GET"
		ctx.Request.Header.Set("Content-Type", "application/json")
		params := []gin.Param{
			{
				Key:   "id",
				Value: "1",
			},
		}
		ctx.Params = params
		productTestController := NewProductController(mockService)
		productTestController.GetProductById(ctx)

		assert.Equal(t, http.StatusNotFound, w.Code)
		fmt.Println(w.Code)
	})

	t.Run("Success", func(t *testing.T) {
		id := "1"
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

		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = &http.Request{
			Header: make(http.Header),
			URL:    &url.URL{},
		}
		ctx.Request.Method = "GET"
		ctx.Request.Header.Set("Content-Type", "application/json")
		params := []gin.Param{
			{
				Key:   "id",
				Value: "1",
			},
		}
		ctx.Params = params
		productTestController := NewProductController(mockService)
		productTestController.GetProductById(ctx)

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

		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = &http.Request{
			Header: make(http.Header),
			URL:    &url.URL{},
		}
		ctx.Request.Method = "POST"
		ctx.Request.Header.Set("Content-Type", "application/json")
		ctx.Request.Body = io.NopCloser(byteProduct)
		productTestController := NewProductController(mockService)
		productTestController.CreateProduct(ctx)

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

		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = &http.Request{
			Header: make(http.Header),
			URL:    &url.URL{},
		}
		ctx.Request.Method = "POST"
		ctx.Request.Header.Set("Content-Type", "application/json")
		ctx.Request.Body = io.NopCloser(byteProduct)
		productTestController := NewProductController(mockService)
		productTestController.CreateProduct(ctx)

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

		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = &http.Request{
			Header: make(http.Header),
			URL:    &url.URL{},
		}
		ctx.Request.Method = "PUT"
		ctx.Request.Header.Set("Content-Type", "application/json")
		ctx.Request.Body = io.NopCloser(byteProduct)
		productTestController := NewProductController(mockService)
		productTestController.EditProduct(ctx)

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

		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = &http.Request{
			Header: make(http.Header),
			URL:    &url.URL{},
		}
		ctx.Request.Method = "PUT"
		ctx.Request.Header.Set("Content-Type", "application/json")
		ctx.Request.Body = io.NopCloser(byteProduct)
		productTestController := NewProductController(mockService)
		productTestController.EditProduct(ctx)

		assert.Equal(t, http.StatusCreated, w.Code)
		t.Log(w.Body.String())
	})
}

func TestCarController_DeleteProduct(t *testing.T) {
	t.Run("Error", func(t *testing.T) {
		id := "1"
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := service.NewMockIProductService(mockCtrl)
		mockService.EXPECT().Delete(id).Return(service.ErrProductNotFound).AnyTimes()

		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = &http.Request{
			Header: make(http.Header),
			URL:    &url.URL{},
		}
		ctx.Request.Method = "DELETE"
		ctx.Request.Header.Set("Content-Type", "application/json")
		params := []gin.Param{
			{
				Key:   "id",
				Value: "1",
			},
		}
		ctx.Params = params
		productTestController := NewProductController(mockService)
		productTestController.DeleteProduct(ctx)

		assert.Equal(t, http.StatusNotFound, w.Code)
		t.Log(w.Body.String())
	})

	t.Run("Success", func(t *testing.T) {
		id := "1"
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := service.NewMockIProductService(mockCtrl)
		mockService.EXPECT().Delete(id).Return(nil).AnyTimes()

		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = &http.Request{
			Header: make(http.Header),
			URL:    &url.URL{},
		}
		ctx.Request.Method = "DELETE"
		ctx.Request.Header.Set("Content-Type", "application/json")
		params := []gin.Param{
			{
				Key:   "id",
				Value: "1",
			},
		}
		ctx.Params = params
		productTestController := NewProductController(mockService)
		productTestController.DeleteProduct(ctx)

		assert.Equal(t, http.StatusAccepted, w.Code)
		t.Log(w.Body.String())
	})
}
