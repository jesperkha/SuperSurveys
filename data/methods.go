package data

import (
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

// Gets a document from the mongodb database and writes the data to the result variable.
// The result parameter MUST be passed in as a pointer to a slice.
func GetDocumentById(id string, result interface{}) {

	// Gets documents by given id
	
	ctx, cancel := getContext()
	defer cancel()

	collection := client.Database("Survey-App").Collection("surveys")
	cursor, err := collection.Find(ctx, bson.M{ "surveyId": id }) // UUID
	if err != nil { log.Fatal(err) }


	// Writes document data to given result slice

	if err = cursor.All(ctx, result); err != nil {
		log.Fatal(err)
	}
}