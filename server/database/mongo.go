package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

const connectionString = "mongodb://localhost:27017"
const databaseName = "todo"
const collectionName = "todolist"

// Init makes the connection creates de collection
func Init() {
	// set client options
	clientOption := options.Client().ApplyURI(connectionString)
	// connect to mongo
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to mongodb")

	collection = client.Database(databaseName).Collection(collectionName)

	fmt.Println("Collection created and ready!")
}

// GetCollectionPointer returns the pointer to the collection
func GetCollectionPointer() *mongo.Collection {
	log.Println("Collection", collection)
	return collection
}
