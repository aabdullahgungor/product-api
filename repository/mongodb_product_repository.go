package repository

import (
	"context"
	"fmt"
	"strconv"

	"github.com/aabdullahgungor/product-api/database"
	"github.com/aabdullahgungor/product-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoDbProductRepository struct {

}

func (m *MongoDbProductRepository) GetAll() ([]models.Product, error) {

	db, err := database.GetMongoDB()
	if err != nil {
        panic(err)
    } 

	productCollection := db.Collection("product")
	
	var products []models.Product
	productCursor, err := productCollection.Find(context.TODO(), bson.M{})
	if err != nil {
        panic(err)
	}
	if err = productCursor.All(context.TODO(), &products); err != nil {
        panic(err)
	}
	
	return products, err

}

func (m *MongoDbProductRepository) GetById(id string) (models.Product, error) { 

	db, err := database.GetMongoDB()
	if err != nil {
        panic(err)
    } 

	productCollection := db.Collection("product")

	objId, _ := primitive.ObjectIDFromHex(id)
	
	var product models.Product
	err = productCollection.FindOne(context.TODO(), bson.M{"_id": objId}).Decode(&product)
	if err != nil {
		panic(err)
	}
	
	return product, err

}

func (m *MongoDbProductRepository) Create(product *models.Product) (string, error) {

	db, err := database.GetMongoDB()
	if err != nil {
        panic(err)
    } 

	productCollection := db.Collection("product")

	result, err := productCollection.InsertOne(context.TODO(), product)

	if err != nil {
        panic(err)
	}

	message := fmt.Sprintf("\ndisplay the ids of the newly inserted objects: %v", result.InsertedID)
	
	return message, err
}

func (m *MongoDbProductRepository) Edit(product *models.Product) (string, error) { 

	db, err := database.GetMongoDB()
	if err != nil {
        panic(err)
    } 

	productCollection := db.Collection("product")

	bsonBytes, err:= bson.Marshal(&product)
	
	if err != nil {
            panic(fmt.Errorf("can't marshal:%s", err))
    }

	var upt bson.M
	bson.Unmarshal(bsonBytes, &upt)

	update := bson.M{"$set": upt,}

	result, err := productCollection.UpdateOne(context.TODO(), bson.M{"_id": product.Id}, update)

	if err != nil {
        panic(err)
	}

	message := "Number of documents updated:"+ strconv.Itoa(int(result.ModifiedCount))

	return message, err

}

func (m *MongoDbProductRepository) Delete(id string) (string, error) { 
	
	db, err := database.GetMongoDB()
	if err != nil {
        panic(err)
    } 

	productCollection := db.Collection("product")

	objId, _ := primitive.ObjectIDFromHex(id)
	
	result, err := productCollection.DeleteOne(context.TODO(), bson.M{"_id": objId})

	// check for errors in the deleting
	if err != nil {
        panic(err)
	}

	// display the number of documents deleted

	message := "deleting the first result from the search filter\n"+ "Number of documents deleted:"+strconv.Itoa(int(result.DeletedCount))

	return message, err

}