package utils

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDb() (*mongo.Client, context.CancelFunc) {
	mongoDbPath := os.Getenv("MONGODBPATH")
	if mongoDbPath == "" {
		log.Fatalln(
			"[ERROR]\tPlease set the environment variable : MONGODBPATH\n",
			"\tex.: MONGODBPATH = 'mongodb://localhost:27017'",
		)
	}

	clientOptions := options.Client().ApplyURI(mongoDbPath)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatalln("[ERROR]\t", err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatalln("[ERROR]\t", err.Error())
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalln("[ERROR]\t", err.Error())
	}
	log.Print("[INFO]\tConnected to MongoDB for CommentsProject.")
	return client, cancel
}

func GetCommentCollection(client *mongo.Client) *mongo.Collection {
	collection := client.Database("CommentsDB").Collection("Comments")
	return collection
}
