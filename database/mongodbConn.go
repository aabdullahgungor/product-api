package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetMongoDB() (*mongo.Database, error) {
                           
    databaseURL := "mongodb+srv://abdullahgungor:Ag7410@cluster0.xbwcqpz.mongodb.net/?retryWrites=true&w=majority"
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()   
    // mongo.Connect return mongo.Client method
    client, err := mongo.Connect(ctx, options.Client().ApplyURI(databaseURL))
    if err != nil {
        panic(err)
    }
    
    db := client.Database("productdb")
     
    return db,  nil
}










