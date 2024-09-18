package config

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func ConnectMongoDB() {
	var err error
	MongoClient, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(EnvConfig.MONGODB_URI))
	if err != nil {
		log.Fatal(err)
	}
	// Send a ping to confirm a successful connection
  if err := MongoClient.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
    panic(err)
  }
  fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
	fmt.Println("Live MongoDB Database connected.")
}

func DisconnectMongoDB() {
	err := MongoClient.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
}

func GetCollection(collectionName string) *mongo.Collection {
	return MongoClient.Database(EnvConfig.DB_NAME).Collection(collectionName)
}
