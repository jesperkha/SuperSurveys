package data

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func InsertSurvey(survey Survey) (err error) {
	ctx, cancel := getContext()
	defer cancel()

	surveyCollection := getCollection("surveys")
	_, err = surveyCollection.InsertOne(ctx, survey)
	if err != nil {
		return err
	}

	return nil
}


func GetSurveysById(id string) (result []Survey, err error) {
	ctx, cancel := getContext()
	defer cancel()

	collection := getCollection("surveys")
	cursor, err := collection.Find(ctx, bson.M{"surveyId": id}) // UUID
	if err != nil {
		return result, err
	}

	// Writes document data to given result slice
	err = cursor.All(ctx, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}


func InsertSubmission(surveyId string, answers [][]string) (numModified int64, err error) {
	ctx, cancel := getContext()
	defer cancel()

	update := bson.D{primitive.E{Key: "$push", Value: bson.A{primitive.E{Key: "answers", Value: answers}}}}
	id := bson.M{"surveyId": surveyId}

	collection := getCollection("surveys")
	result, err := collection.UpdateOne(ctx, id, update)
	if err != nil {
		return -1, err
	}

	return result.ModifiedCount, nil
}
