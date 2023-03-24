package models

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Brand struct {
	Id   primitive.ObjectId `bson:"_id"`
	Name string        `bson:"name"`
}

func (brand Brand) ToString() string {
	result := fmt.Sprintf("\nbrand id: %s", brand.Id)
	result = result + fmt.Sprintf("\nbrand name: %s", brand.Name)
	return result
}
