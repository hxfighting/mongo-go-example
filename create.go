package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID   primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name string             `json:"name" bson:"name,omitempty"`
	Age  byte               `json:"age" bson:"age,omitempty"`
}

func createCollection(ctx context.Context, db *mongo.Database) {
	jsonSchema := bson.M{
		"bsonType": "object",
		"required": []string{"name", "age"},
		"properties": bson.M{
			"name": bson.M{
				"bsonType":    "string",
				"description": "the name of the user, which is required and must be a string",
			},
			"age": bson.M{
				"bsonType":    "int",
				"description": "the age of the user, which is required and must be a int",
				"minimum":     0,
				"maximum":     200,
			},
		},
	}
	validator := bson.M{
		"$jsonSchema": jsonSchema,
	}
	opts := options.CreateCollection().SetValidator(validator)
	if err := db.CreateCollection(ctx, userCollection, opts); err != nil {
		log.Fatalln(err.Error())
	}
}

func insertOne1(ctx context.Context, db *mongo.Database) {
	if _, err := db.Collection(userCollection).InsertOne(ctx, bson.M{"name": "test", "age": 3}); err != nil {
		log.Fatalln(err.Error())
	}
	opts := options.InsertOne().SetBypassDocumentValidation(true)
	if _, err := db.Collection(userCollection).InsertOne(ctx, bson.M{"name": 3, "age": "test"}, opts); err != nil {
		log.Fatalln(err.Error())
	}
}

func insertOne2(ctx context.Context, db *mongo.Database) {
	user := User{
		Name: "test",
		Age:  25,
	}
	if _, err := db.Collection(userCollection).InsertOne(ctx, user); err != nil {
		log.Fatalln(err.Error())
	}
}

func insertMany(ctx context.Context, db *mongo.Database) {
	users := []interface{}{
		User{Name: "test1", Age: 10},
		User{Name: "test2", Age: 13},
	}
	if _, err := db.Collection(userCollection).InsertMany(ctx, users); err != nil {
		log.Fatalln(err.Error())
	}
	users2 := []interface{}{
		bson.M{"name": "test3", "age": 15},
		bson.M{"name": "test4", "age": 16},
	}
	if _, err := db.Collection(userCollection).InsertMany(ctx, users2); err != nil {
		log.Fatalln(err.Error())
	}
}

func updateOne(ctx context.Context, db *mongo.Database) {
	id, _ := primitive.ObjectIDFromHex("5f226a0ef5d91bc30f112dda")
	filter := bson.M{"_id": id}
	update := bson.D{{
		"$set",
		bson.M{"name": "ffff"},
	}}
	if _, err := db.Collection(userCollection).UpdateOne(ctx, filter, update); err != nil {
		log.Fatalln(err.Error())
	}
}

func updateInsert(ctx context.Context, db *mongo.Database) {
	id, _ := primitive.ObjectIDFromHex("5f226a0ef5d91bc30f112ddd")
	filter := bson.M{"_id": id}
	update := bson.D{{
		"$set",
		bson.M{"name": "dddd"},
	}}
	opts := options.Update()
	opts.SetUpsert(true)
	if _, err := db.Collection(userCollection).UpdateOne(ctx, filter, update, opts); err != nil {
		log.Fatalln(err.Error())
	}
}

func updateMany(ctx context.Context, db *mongo.Database) {
	filter := bson.M{"age": 3}
	update := bson.D{{
		"$set",
		User{
			Name: "aaaaa",
		},
	}}
	if _, err := db.Collection(userCollection).UpdateMany(ctx, filter, update); err != nil {
		log.Fatalln(err.Error())
	}
}

func queryOne(ctx context.Context, db *mongo.Database) {
	id, _ := primitive.ObjectIDFromHex("5f226a0ef5d91bc30f112dda")
	filter := bson.M{"_id": id}
	res := db.Collection(userCollection).FindOne(ctx, filter)
	if res.Err() != nil {
		log.Fatalln(res.Err().Error())
	}
	user := User{}
	if err := res.Decode(&user); err != nil {
		log.Fatalln(err.Error())
	}
	log.Println(user)
}

func queryMany(ctx context.Context, db *mongo.Database) {
	filter := bson.M{"age": 3}
	opts := options.Find()
	opts.SetMaxTime(time.Second * 10)
	opts.SetLimit(1)
	cursor, err := db.Collection(userCollection).Find(ctx, filter, opts)
	if err != nil {
		log.Fatalln(err.Error())
	}
	users := []User{}
	//users := []bson.M{}
	if err = cursor.All(ctx, &users); err != nil {
		log.Fatalln(err.Error())
	}
	//// 第二种方式来解析数据，在解析大数量时，推荐第二种方式，第一种方式更耗性能
	//defer cursor.Close(ctx)
	//for cursor.Next(ctx) {
	//	user := User{}
	//// user := bson.M{}
	//	if err = cursor.Decode(&user); err != nil {
	//		log.Fatalln(err)
	//	}
	//	users = append(users, user)
	//}
	log.Println(users)
}

func deleteOne(ctx context.Context, db *mongo.Database) {
	id, _ := primitive.ObjectIDFromHex("5f226a0ef5d91bc30f112dda")
	filter := bson.M{"_id": id}
	if _, err := db.Collection(userCollection).DeleteOne(ctx, filter); err != nil {
		log.Fatalln(err.Error())
	}
}

func deleteMany(ctx context.Context, db *mongo.Database)  {
	filter := bson.M{"age": 3}
	if _, err := db.Collection(userCollection).DeleteMany(ctx, filter); err != nil {
		log.Fatalln(err.Error())
	}
}
