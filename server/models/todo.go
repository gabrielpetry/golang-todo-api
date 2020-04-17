package models

import (
	"context"
	"fmt"
	"log"

	"../database"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection = database.GetCollectionPointer()

// Todo is the basic todo struct
type Todo struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Task      string             `json:"task" bson:"task" validate:"required"`
	Completed bool               `json:"completed" bson:"copleted" validate:"required"`
}

// Validate runs govalidator for the struct
func (todo *Todo) Validate() error {
	validate := validator.New()
	return validate.Struct(todo)
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

// Update updatse an element
func (todo *Todo) Update(id string) error {
	log.Println("Updating", todo)
	oid, err := primitive.ObjectIDFromHex(id)
	collection := database.GetCollectionPointer()
	doc, err := collection.UpdateOne(
		context.Background(),
		bson.M{"_id": bson.M{"$eq": oid}},
		bson.M{"$set": todo},
	)
	log.Println(doc, err)
	return err
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
