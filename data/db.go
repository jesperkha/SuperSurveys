package data

import (
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*

	Surveys

*/


type Survey struct {
	SurveyId     string
	Creator      string
	Name         string
	Desc         string
	NumQuestions int
	Questions    []Question
	Answers      []string
}

type Question struct {
	Class    string
	Optional bool
	Prompt   string
	Options  []string
}


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
	cursor, err := collection.Find(ctx, bson.M{ "surveyId": id }) // UUID
	if err != nil {
		return nil, err
	}

	// Writes document data to given result slice
	err = cursor.All(ctx, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}


func InsertSubmission(surveyId string, answers string) (numModified int64, err error) {
	ctx, cancel := getContext()
	defer cancel()

	update := bson.D{primitive.E{Key: "$push", Value: bson.D{primitive.E{Key: "answers", Value: answers}}}}
	id := bson.M{"surveyId": surveyId}

	collection := getCollection("surveys")
	result, err := collection.UpdateOne(ctx, id, update)
	if err != nil {
		return -1, err
	}

	return result.ModifiedCount, nil
}


/*

	Users

*/


type User struct {
	UserId    string
	Username  string
	Password  string
	Email     string
	Timestamp string
}


func (usr *User) Verify() bool {
	_, err := GetUser(usr.Username, usr.Password)
	return err == nil
}


func (usr *User) GetSurveys(surveyId string) (surveys []Survey, err error) {
	surveys, err = GetSurveysById(usr.UserId)
	if err != nil {
		return surveys, err
	}

	return surveys, nil
}


func InsertUser(username string, password string, email string) (user User, err error) {
	user = User{
		UserId:    uuid.NewString(),
		Username:  username,
		Password:  password,
		Email:     email,
		Timestamp: time.Now().String(),
	}

	ctx, cancel := getContext()
	defer cancel()

	userCollection := getCollection("users")
	_, err = userCollection.InsertOne(ctx, user)
	if err != nil {
		return user, err
	}

	return user, nil
}


func GetUser(username string, password string) (user User, err error) {
	filter := bson.M{ "username": username, "password": password }

	ctx, cancel := getContext()
	defer cancel()

	userCollection := getCollection("users")
	result := userCollection.FindOne(ctx, filter)

	err = result.Err()
	if err != nil {
		return user, err
	}

	err = result.Decode(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}

