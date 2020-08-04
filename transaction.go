package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type Author struct {
	ID   primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name string             `json:"name"`
}

type Article struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title    string             `json:"title"`
	AuthorID primitive.ObjectID `json:"author_id" bson:"author_id"`
}

func transaction(ctx context.Context, db *mongo.Database) {
	authorCollection := "author"
	articleCollection := "article"
	err := db.Client().UseSession(ctx, func(sessionContext mongo.SessionContext) error {
		if err := sessionContext.StartTransaction(); err != nil {
			return err
		}
		authorID := primitive.NewObjectID()
		author := Author{
			ID:   authorID,
			Name: "aa",
		}
		if _, err := db.Collection(authorCollection).InsertOne(sessionContext, author); err != nil {
			_ = sessionContext.AbortTransaction(sessionContext)
			return err
		}
		article := Article{
			ID:       primitive.NewObjectID(),
			Title:    "test",
			AuthorID: authorID,
		}
		if _, err := db.Collection(articleCollection).InsertOne(sessionContext, article); err != nil {
			_ = sessionContext.AbortTransaction(sessionContext)
			return err
		}
		return sessionContext.CommitTransaction(sessionContext)
	})
	if err != nil {
		log.Fatalln(err)
	}
}
