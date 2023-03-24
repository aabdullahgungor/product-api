package models

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
)

type Brand struct {
	Id   bson.ObjectId `bson:"_id"`
	Name string        `bson:"name"`
}

func (brand Brand) ToString() string {
	result := fmt.Sprintf("\nbrand id: %s", brand.Id)
	result = result + fmt.Sprintf("\nbrand name: %s", brand.Name)
	return result
}
