package models

import (
	"context"
	"fmt"
	"log"

	"../database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection = database.GetCollectionPointer()

// Todo is the basic todo struct
type Todo struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Task      string             `json:"task" bson="task"`
	Completed bool               `json:"completed" bson="copleted"`
}

// InsertOne handle the insertion of a new entry
func (todo *Todo) InsertOne() error {
	insertResult, err := database.GetCollectionPointer().InsertOne(context.Background(), todo)

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
	log.Println(database.GetCollectionPointer())
	cursor, err := database.GetCollectionPointer().Find(context.TODO(), bson.D{})
	if err != nil {
		return todos, err
	}

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
