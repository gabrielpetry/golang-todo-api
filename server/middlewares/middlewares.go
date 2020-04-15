package middlewares

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"../models"

	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


const connectionString = "mongodb://localhost:27017"
const databaseName = "todo"
const collectionName = "todolist"

var collection *mongo.Collection

func init() {
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


func GetAllTask(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Access-Control-Allow-Origin", "*")
  payload := getAllTask()
  json.NewEncoder(w).Encode(payload)
}

func getAllTask() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var results []primitive.M
	for cur.Next(context.Background()) {
		var result bson.M
		e := cur.Decode(&result)
		if e != nil {
			log.Fatal(e)
		}
		// fmt.Println("cur..>", cur, "result", reflect.TypeOf(result), reflect.TypeOf(result["_id"]))
		results = append(results, result)

	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.Background())
	return results
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
  w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Header().Set("Access-Control-Allow-Methods", "POST")
  w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
  var task models.Todo
  _ = json.NewDecoder(r.Body).Decode(&task)
  fmt.Println(task, r.Body)
  insertOneTask(task)
  json.NewEncoder(w).Encode(task)
}




