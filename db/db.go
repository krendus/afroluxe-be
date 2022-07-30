package db

import (
	"context"
	"fmt"
	"log"

	"github.com/afroluxe/afroluxe-be/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var env = config.LoadEnv()

func connectDB() *mongo.Database {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(env.MongodbUri))
	if err != nil {
		log.Fatalf("Error connecting to DB:: \n %v", err)
	}
	fmt.Printf("DB Connected Successful DB:: \n  %v", env.DbName)
	return client.Database(env.DbName)
}

var db = connectDB()

func CollectionInstance(collectionName string) *mongo.Collection {
	return db.Collection(collectionName)
}
