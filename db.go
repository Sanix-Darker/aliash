package main

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection
var ctx = context.TODO()

var MONGODB_URL = os.Getenv("MONGODB_URL")

var DATABASE = os.Getenv("DATABASE")

func init() {
	clientOptions := options.Client().ApplyURI(MONGODB_URL)

	client, err := mongo.Connect(ctx, clientOptions)
	Must(err)

	Must(client.Ping(ctx, nil))

	collection = client.Database(DATABASE).Collection("aliases")
}

// createAliases to create a new alias
func createAliases(as *Aliases) error {
	_, err := collection.InsertOne(ctx, as)
	return err
}

// getAllAliases To get all aliases
// passing bson.D{{}} matches all documents in the collection
func getAllAliases() ([]*Aliases, error) {
	filter := bson.D{{}}
	return filterAliasessBy(filter)
}

// getAliasess to get aliases from mongodb
// with a given filter sjon filter
// A slice of aliases for storing the decoded documents
func filterAliasessBy(filter interface{}) ([]*Aliases, error) {
	var aliases []*Aliases

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return aliases, err
	}

	for cur.Next(ctx) {
		var t Aliases
		err := cur.Decode(&t)
		if err != nil {
			return aliases, err
		}

		aliases = append(aliases, &t)
	}

	if err := cur.Err(); err != nil {
		return aliases, err
	}

	// once exhausted, close the cursor
	cur.Close(ctx)

	if len(aliases) == 0 {
		return aliases, mongo.ErrNoDocuments
	}

	return aliases, nil
}
