package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

type repository[T any] struct {
	db *mongo.Client
}

// This repository uses sql server
func InitializeDB() (*mongo.Client, error) {

	// Load the .env file.
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file is specified.")
	}

	// Fetch the uri from the environment variable.
	uri := os.Getenv("MONGODB_URI")

	if uri == "" {
		return nil, errors.New("you must set your 'MONGODB_URI' environment variable")
	}

	// Connect to the mongodb database.
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	// Error handling.
	if err != nil {
		return nil, fmt.Errorf("error connecting to MongoDB: %s", err)
	}

	fmt.Println("Connected to MongoDB!")
	// Return either the client or nil in case of error.
	return client, nil
}

// Function to store the userdata
func (r *repository[T]) Create(model T, collection string) (bool, error) {
	
	coll := r.db.Database("hirego").Collection(collection)
	_, err := coll.InsertOne(context.TODO(), model)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *repository[T]) GetByField(field string, field_value string, collection string) (interface{}, error) {
	coll := r.db.Database("hirego").Collection(collection)
    
	// Define the filter.
	filter := bson.D{{Key: field, Value: field_value}}
    
	// Define the result.
	var result map[string]interface{}

	// Find the data from the collection.
	err := coll.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *repository[T]) InsertData(workinforamtion interface{}, collection string) (bool, error) {
	coll := r.db.Database("hirego").Collection(collection)
	
	// Insert the data into the collection.
	_, err := coll.InsertOne(context.TODO(), workinforamtion)


	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *repository[T]) GetData(username string, collection string) (interface{}, error) {
	coll := r.db.Database("hirego").Collection(collection)
    
	// Define the filter.
	filter := bson.D{{Key: "username", Value: username}}
    
	// Define the data structure of the  result.
	var result interface{} 

	// Find the data from the collection.
	err := coll.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *repository[T]) GetAll(collection string) ([]map[string]interface{}, error) {
	coll := r.db.Database("hirego").Collection(collection)
    
	// Find all the data from the collection.
	cursor, err := coll.Find(context.TODO(), bson.D{})

	if err != nil {
		return nil, err
	}

	// Close the cursor after the function ends.
	defer cursor.Close(context.TODO())

	// Define the data structure of result.
	var results []map[string]interface{}
    
	// Iterate over the cursor to get the data.
	for cursor.Next(context.TODO()) {

		// Define the data structure to decode the individual elements later.
		var result map[string]interface{}

		// Decode the data from the cursor.
		err := cursor.Decode(&result)

		if err != nil {
			return nil, err
		}

		// Append the result to the results lists.
		results = append(results, result)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	fmt.Println("The result is here", results)

	return results, nil
}

func (r *repository[T]) GetAllByField(field string, field_value string, collection string) ([]map[string]interface{}, error) {
	coll := r.db.Database("hirego").Collection(collection)
    

	// Define the filter.
	filter := bson.D{{Key: field, Value: field_value}}

	// Find all the data from the collection.
	cursor, err := coll.Find(context.TODO(), filter)

	if err != nil {
		return nil, err
	}

	// Close the cursor after the function ends.
	defer cursor.Close(context.TODO())
    
	// Define the data structure of result.
	var results []map[string]interface{}
    
	// Iterate over the cursor to get the data.
	for cursor.Next(context.TODO()) {
		// Define the data structure to decode the individual elements later.
		var result map[string]interface{}

		// Decode the data from the cursor.
		err := cursor.Decode(&result)
		if err != nil {
			return nil, err
		}

		// Append the result to the results lists.
		results = append(results, result)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return results, nil
}