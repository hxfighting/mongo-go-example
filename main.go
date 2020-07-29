package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	dbName         = "test"
	collectionName = dbName
)

func main() {
	ctx := context.Background()
	client := getClient()
	defer client.Disconnect(ctx)
	res, err := client.ListDatabases(ctx, bson.M{})
	if err != nil {
		log.Fatalln()
	}
	for _, database := range res.Databases {
		log.Println(database.Name)
	}
}

func getClient() *mongo.Client {
	clientOption := options.Client()
	clientOption.ApplyURI("mongodb://test:test@127.0.0.1:27017")
	//clientOption.SetReplicaSet("rs1")

	client, err := mongo.NewClient(clientOption)
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	return client
}
