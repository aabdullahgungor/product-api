package models

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"
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
