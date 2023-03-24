package models

import (
	"fmt"
	"strings"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Product struct {
	Id         bson.ObjectId `bson:"_id"`
	Name       string        `bson:"name"`
	Price      float64       `bson:"price"`
	Quantity   int64         `bson:"quantity"`
	Status     bool          `bson:"status"`
	Date       time.Time     `bson:"date"`
	CategoryId bson.ObjectId `bson:"categoryId"`
	Brand      Brand         `bson:"brand"`
	Colors     []string      `bson:"colors"`
}

func (product Product) ToString() string {
	result := fmt.Sprintf("id: %s", product.Id)
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
