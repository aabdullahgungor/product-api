package repository

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/aabdullahgungor/product-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ErrProductNotFound = errors.New("FromRepository - product not found")
)

type MongoDbProductRepository struct {
	connectionPool *mongo.Database
}

func NewMongoDbProductRepository() *MongoDbProductRepository {
	databaseURL := "mongodb+srv://<username>:<password>@cluster0.xbwcqpz.mongodb.net/?retryWrites=true&w=majority"
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// mongo.Connect return mongo.Client method
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(databaseURL))
	if err != nil {
		panic(err)
	}

	db := client.Database("productdb")

	return &MongoDbProductRepository{
		connectionPool: db,
	}
}

func (m *MongoDbProductRepository) GetAllProducts() ([]models.Product, error) {

	productCollection := m.connectionPool.Collection("product")

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

func (m *MongoDbProductRepository) GetProductById(id string) (models.Product, error) {

	productCollection := m.connectionPool.Collection("product")

	objId, _ := primitive.ObjectIDFromHex(id)

	var product models.Product
	err := productCollection.FindOne(context.TODO(), bson.M{"_id": objId}).Decode(&product)
	if err != nil {
		return models.Product{}, ErrProductNotFound
	}

	return product, nil

}

func (m *MongoDbProductRepository) CreateProduct(product *models.Product) error {

	productCollection := m.connectionPool.Collection("product")

	result, err := productCollection.InsertOne(context.TODO(), product)

	if err != nil {
		panic(err)
	}

	log.Printf("\ndisplay the ids of the newly inserted objects: %v", result.InsertedID)

	return err
}

func (m *MongoDbProductRepository) EditProduct(product *models.Product) error {

	productCollection := m.connectionPool.Collection("product")

	bsonBytes, err := bson.Marshal(&product)

	if err != nil {
		panic(fmt.Errorf("can't marshal:%s", err))
	}

	var upt bson.M
	bson.Unmarshal(bsonBytes, &upt)

	update := bson.M{"$set": upt}

	result, err := productCollection.UpdateOne(context.TODO(), bson.M{"_id": product.Id}, update)

	if err != nil {
		panic(err)
	}

	log.Println("Number of documents updated:" + strconv.Itoa(int(result.ModifiedCount)))

	return err
}

func (m *MongoDbProductRepository) DeleteProduct(id string) error {

	productCollection := m.connectionPool.Collection("product")

	objId, _ := primitive.ObjectIDFromHex(id)

	result, err := productCollection.DeleteOne(context.TODO(), bson.M{"_id": objId})

	// check for errors in the deleting
	if err != nil {
		panic(err)
	}

	// display the number of documents deleted
	log.Println("deleting the first result from the search filter\n" + "Number of documents deleted:" + strconv.Itoa(int(result.DeletedCount)))

	return err
}
