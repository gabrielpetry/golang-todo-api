package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

const connectionString = "mongodb://localhost:27017/?connectTimeoutMS=5000"
const databaseName = "todo"
const collectionName = "todolist"

// Init makes the connection creates de collection
func Init() {
	log.Println("Trying connection to MongoDb")
	// set client options
	clientOption := options.Client()
	clientOption.SetConnectTimeout(2 * time.Second)
	clientOption.SetSocketTimeout(2 * time.Second)
	clientOption.SetMaxConnIdleTime(1 * time.Second)
	clientOption.ApplyURI(connectionString)
	// connect to mongo
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to mongodb")

	collection = client.Database(databaseName).Collection(collectionName)

	log.Println("Collection created and ready!")
}

// GetCollectionPointer returns the pointer to the collection
func GetCollectionPointer() *mongo.Collection {
	return collection
}
