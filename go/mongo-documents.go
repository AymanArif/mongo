package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(ctx)

	languagesDatabase := client.Database("languages")
	golangCollection := languagesDatabase.Collection("golang")
	nodejsCollection := languagesDatabase.Collection("nodejs")

	nodejsResult, err := nodejsCollection.InsertOne(ctx, bson.D{
		{Key: "title", Value: "NodeJS Developer with Array document"}, // Key and Value are optional. Better to explicitly put it for clean code perspective
		{"name", "Ayman"},
		{"tags", bson.A{"systems programming"}},
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(nodejsResult.InsertedID)

	golangResult, err := golangCollection.InsertMany(ctx, []interface{}{

		bson.D{
			{"nodejscollectionId", nodejsResult.InsertedID},
			{Key: "title", Value: "Golang Developer with Array document"}, // Key and Value are optional. Better to explicitly put it for clean code perspective
			{"name", "Ayman"},
			{"tags", bson.A{"web", "cloud"}},
		},
		bson.D{
			{"nodejscollectionId", nodejsResult.InsertedID},
			{Key: "title", Value: "NodeJS Developer with Array document"}, // Key and Value are optional. Better to explicitly put it for clean code perspective
			{"name", "Ayman"},
			{"tags", bson.A{"systems programming"}},
		},
	})

	fmt.Println(golangResult.InsertedIDs)

}
