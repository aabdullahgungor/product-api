package models

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Brand struct {
	Id   primitive.ObjectID `bson:"_id,omitempty"`
	Name string        `bson:"name,omitempty"`
}

func (brand Brand) ToString() string {
	result := fmt.Sprintf("\nbrand id: %s", brand.Id.Hex())
	result = result + fmt.Sprintf("\nbrand name: %s", brand.Name)
	return result
}
