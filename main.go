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
	userCollection = "users"
)

func main() {
	ctx := context.Background()
	client := getClient()
	defer client.Disconnect(ctx)
	db := client.Database(dbName)
	transaction(ctx, db)
}

func getClient() *mongo.Client {
	clientOption := options.Client()
	clientOption.ApplyURI("mongodb://test:test@127.0.0.1:27017")
	clientOption.SetReplicaSet("rs1")

	client, err := mongo.NewClient(clientOption)
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
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

func listDatabases(ctx context.Context, client *mongo.Client) {
	res, err := client.ListDatabases(ctx, bson.M{})
	if err != nil {
		log.Fatalln()
	}
	for _, database := range res.Databases {
		log.Println(database.Name)
	}
}
