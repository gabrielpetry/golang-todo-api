package models

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

const connectionString = "mongodb://localhost:27017"
const databaseName = "todo"
const collectionName = "todolist"

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

// Todo is the basic todo struct
type Todo struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Task   string             `json:"task"`
	Status bool               `json:"status,omitempty"`
}

func NewTodo() *Todo {
	return &Todo{}
}

// InsertOne handle the insertion of a new entry
func (todo *Todo) InsertOne() error {
	fmt.Println(todo)
	insertResult, err := collection.InsertOne(context.Background(), todo)

	if err != nil {
		return err
	}

	fmt.Println("Inserted a Single Record ", insertResult.InsertedID)
	return nil
}

// GetAll searchs all todos in the database
func (todo *Todo) GetAll() ([]Todo, error) {
	fmt.Println("Getting all todos")
	todos := []Todo{}
	cursor, _ := collection.Find(context.TODO(), bson.D{})
	// err = collection.FindOne(context.TODO(), filter).Decode(&result)
	log.Println(cursor)

	for cursor.Next(context.TODO()) {
		var elem Todo
		err := cursor.Decode(&elem)

		if err != nil {
			log.Fatal(err)
		}

		todos = append(todos, elem)
	}

	cursor.Close(context.TODO())
	return todos, nil
}
