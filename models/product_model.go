package models

import (
	"fmt"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	Id         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name       string        `bson:"name,omitempty" json:"name"`
	Price      float64       `bson:"price,omitempty" json:"price"`
	Quantity   int64         `bson:"quantity,omitempty" json:"quantity"`
	Status     bool          `bson:"status,omitempty" json:"status"`
	Date       time.Time     `bson:"date,omitempty" json:"date"`
	CategoryId primitive.ObjectID `bson:"categoryId,omitempty" json:"category_id"`
	Brand      Brand         `bson:"brand,omitempty" json:"brand"`
	Colors     []string      `bson:"colors,omitempty" json:"colors"`
}

func (product Product) ToString() string {
	result := fmt.Sprintf("id: %s", product.Id.Hex())
	result = result + fmt.Sprintf("\nname: %s", product.Name)
	result = result + fmt.Sprintf("\nprice: %0.1f", product.Price)
	result = result + fmt.Sprintf("\nquantity: %d", product.Quantity)
	result = result + fmt.Sprintf("\nstatus: %t", product.Status)
	result = result + fmt.Sprintf("\ndate: %s", product.Date.Format("2006-01-02"))
	result = result + fmt.Sprintf("\ncategory id: %s", product.CategoryId)
	result = result + product.Brand.ToString()
	result = result + fmt.Sprintf("\ncolors: %s", strings.Join(product.Colors, ", "))
	return result
}
