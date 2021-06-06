package data

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var clientContext context.Context
const Timeout int = 5


// Shorthand for getting context with global timeout value

func getContext() (ctx context.Context, cancel context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Duration(Timeout) * time.Second)
}


// Connects to database and creates new mongo.Client and loads .env file
func ConnectClient() {

	// Set up .env file

	err := godotenv.Load("./.env")
	if err != nil { log.Fatal(err) }


	// Set up MongoDB client

	URI := fmt.Sprintf("mongodb+srv://jesperkha:%s@cluster-1.d5rss.mongodb.net/?retryWrites=true&w=majority", os.Getenv("MONGODB_PASSWORD"))
	client, err = mongo.NewClient(options.Client().ApplyURI(URI))
	if err != nil { log.Fatal(err) }

	clientContext, cancel := getContext()
	defer cancel()

	err = client.Connect(clientContext)
	if err != nil { log.Fatal(err) }

}


// Fully disconnects from database.
// A new client can be made by calling Init() again, but it is not recommended.
func CloseClient() {
	client.Disconnect(clientContext)
}
