package data

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client *mongo.Client
	clientContext context.Context
	Timeout int = 5
)


func getContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Duration(Timeout) * time.Second)
}


func getCollection(name string) *mongo.Collection {
	return Client.Database("Survey-App").Collection(name)
}


func ConnectClient() (err error) {
	token := os.Getenv("MONGODB_PASSWORD")
	URI := fmt.Sprintf("mongodb+srv://jesperkha:%s@cluster-1.d5rss.mongodb.net/?retryWrites=true&w=majority", token)
	Client, err = mongo.NewClient(options.Client().ApplyURI(URI))
	if err != nil {
		return err
	}

	clientContext, cancel := getContext()
	defer cancel()

	err = Client.Connect(clientContext)
	if err != nil {
		return err
	}

	return nil
}


func CloseClient() {
	Client.Disconnect(clientContext)
}
