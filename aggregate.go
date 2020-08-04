package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CountResult struct {
	UID   int64   `json:"uid" bson:"_id"`
	Total float32 `json:"Total" bson:"total"`
}

func aggregate(ctx context.Context, db *mongo.Database) {
	matchStage := bson.D{{"$match", bson.M{"status": 1}}}
	groupStage := bson.D{{
		"$group",
		bson.D{
			{"_id", "$uid"},
			{"total", bson.D{{
				"$sum", "$score"}}},
		}}}
	cursor, err := db.Collection(userCollection).Aggregate(ctx, mongo.Pipeline{matchStage, groupStage})
	if err != nil {
		log.Fatalln(err)
	}
	res := []CountResult{}
	if err = cursor.All(ctx, &res); err != nil {
		log.Fatalln(err)
	}
	log.Println(res)
}
