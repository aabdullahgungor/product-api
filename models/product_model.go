package models

import (
	"fmt"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	Id         primitive.ObjectID `bson:"_id,omitempty"`
	Name       string        `bson:"name,omitempty"`
	Price      float64       `bson:"price,omitempty"`
	Quantity   int64         `bson:"quantity,omitempty"`
	Status     bool          `bson:"status,omitempty"`
	Date       time.Time     `bson:"date,omitempty"`
	CategoryId primitive.ObjectID `bson:"categoryId,omitempty"`
	Brand      Brand         `bson:"brand,omitempty"`
	Colors     []string      `bson:"colors,omitempty"`
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
