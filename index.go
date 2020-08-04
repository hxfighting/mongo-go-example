package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func indexExample(ctx context.Context, db *mongo.Database) {
	// 创建一个有过期时间的索引
	indexOption := options.Index()
	indexOption.SetName("uid_index")
	indexOption.SetExpireAfterSeconds(600)
	mod := mongo.IndexModel{
		Keys:    bson.M{"uid": 1},
		Options: indexOption,
	}
	_, err := db.Collection(userCollection).Indexes().CreateOne(ctx, mod)
	if err != nil {
		log.Fatal(err)
	}

	// 创建一个文本索引
	indexOption = options.Index()
	indexOption.SetName("name_index")
	mod = mongo.IndexModel{
		Keys:    bson.M{"name": "text"},
		Options: indexOption,
	}
	_, err = db.Collection(userCollection).Indexes().CreateOne(ctx,mod)
	if err != nil {
		log.Fatal(err)
	}

	// 删除索引
	_, err = db.Collection(userCollection).Indexes().DropOne(ctx,"uid_index")
	if err != nil {
		log.Fatal(err)
	}
}
