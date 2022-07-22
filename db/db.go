package db

import (
	"context"
	"fmt"
	"log"

	"github.com/afroluxe/afroluxe-be/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var env config.Env = config.LoadEnv()

func connectDB() *mongo.Database {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(env.MONGODB_URI))
	if err != nil {
		log.Fatalf("Error connecting to DB:: \n %v", err)
	}
	fmt.Printf("DB Connected Successful DB:: \n  %v", env.DB_NAME)
	return client.Database(env.DB_NAME)
}

var db *mongo.Database = connectDB()

func CollectionInstance(collectionName string) *mongo.Collection {
	return db.Collection(collectionName)
}
