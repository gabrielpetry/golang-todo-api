package models

import (
	"context"

	log "github.com/sirupsen/logrus"

	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"host.local/go/golang-todo-api/src/database"
)

var bookCollection *mongo.Collection = database.GetCollectionPointer()

// Todo is the basic todo struct
type Book struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name" validate:"required"`
	Description string             `json:"description" bson:"description"`
	CreatedAt   string             `json:"created_at" bson:"created_at"`
	UpdatedAt   string             `json:"updated_at" bson:"updated_at"`
	Order       string             `json:"order" bson:"order" validate:"required"`
}

// Validate runs govalidator for the struct
func (todo *Todo) ValidateBook() error {
	validate := validator.New()
	return validate.Struct(todo)
}

// InsertOne insert a new entry
func (todo *Todo) InsertOneBook() error {
	_, err := database.GetCollectionPointer().InsertOne(context.Background(), todo)

	if err != nil {
		return err
	}

	return nil
}

// Update updates an element
func (todo *Todo) UpdateBook(id string) error {
	log.Println("Updating", todo)
	oid, err := primitive.ObjectIDFromHex(id)
	collection := database.GetCollectionPointer()
	doc, err := collection.UpdateOne(
		context.Background(),
		bson.M{"_id": bson.M{"$eq": oid}},
		bson.M{"$set": todo},
	)

	if err != nil {
		log.Error("Error updating", doc, err)
	}

	return err
}

// Delete removes an entry
func (todo *Todo) DeleteBook(id string) error {
	log.Println("Deleting", todo)
	oid, err := primitive.ObjectIDFromHex(id)
	collection := database.GetCollectionPointer()
	doc, err := collection.DeleteOne(
		context.Background(),
		bson.M{"_id": bson.M{"$eq": oid}},
	)

	if err != nil {
		log.Error("Error deleting", doc, err)
	}

	return err
}

// GetAll searchs all todos in the database
func (todo *Todo) GetAllBooks() ([]Todo, error) {
	todos := []Todo{}
	cursor, err := database.GetCollectionPointer().Find(context.TODO(), bson.D{})
	if err != nil {
		log.Error("Error with collection pointer", err)
		return todos, err
	}

	for cursor.Next(context.TODO()) {
		var elem Todo
		err := cursor.Decode(&elem)

		if err != nil {
			log.Error("Error decoding cursor", err)
		}

		todos = append(todos, elem)
	}

	cursor.Close(context.TODO())
	return todos, nil
}
