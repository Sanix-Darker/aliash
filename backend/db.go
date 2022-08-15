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
var ctxB = context.Background()

var MONGODB_URL = os.Getenv("MONGODB_URL")
var DATABASE = os.Getenv("DATABASE")
var MONGO_INITDB_ROOT_USERNAME = os.Getenv("MONGO_INITDB_ROOT_USERNAME")
var MONGO_INITDB_ROOT_PASSWORD = os.Getenv("MONGO_INITDB_ROOT_PASSWORD")

func init() {
	credential := options.Credential{
		Username: MONGO_INITDB_ROOT_USERNAME,
		Password: MONGO_INITDB_ROOT_PASSWORD,
	}
	clientOptions := options.Client().ApplyURI(MONGODB_URL).SetAuth(credential)

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
func getAllAliases() []*Aliases {
	filter := bson.D{{}}
	aliases, _ := filterAliasesBy(filter)

	var sanitizedAliases []*Aliases
	for _, as := range aliases {
		sanitizedAliases = append(sanitizedAliases, as)
	}

	return sanitizedAliases
}

// getAliasess to get aliases from mongodb
// with a given filter sjon filter
// A slice of aliases for storing the decoded documents
func filterAliasesBy(filter interface{}) ([]*Aliases, error) {
	var aliases []*Aliases

	opts := options.Find().SetLimit(100)
	cur, err := collection.Find(ctx, filter, opts)
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

func searchAliases(field, value string) ([]*Aliases, error) {
	var aliases []*Aliases

	query := bson.M{field: bson.M{"$regex": value, "$options": "im"}}
	opts := options.Find().SetLimit(10)
	cursor, err := collection.Find(ctx, query, opts)
	if err != nil {
		return aliases, err
	}

	// var sites []bson.M
	if err = cursor.All(ctx, &aliases); err != nil {
		return aliases, err
	}

	return aliases, err
}
