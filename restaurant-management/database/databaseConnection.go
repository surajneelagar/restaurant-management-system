package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBinstance() *mongo.Client {
	Mongodb := "mongodb://localhost:27017"
	fmt.Print(Mongodb)

	client, err := mongo.NewClient(options.Client().ApplyURI(Mongodb))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connected to mongodb...")
	return client
}

var Client *mongo.Client = DBinstance()

func OpenCollections(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("restaurant").Collection(collectionName)
	return collection
}

// aJx9Jwk90Ah9pFVg
// mongodb+srv://surajneelagar:aJx9Jwk90Ah9pFVg@cluster0.lzjt4.mongodb.net/
