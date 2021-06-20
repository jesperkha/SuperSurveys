package data

import (
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/securecookie"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*

	Surveys

*/


type Survey struct {
	SurveyId     string
	CreatorId    string
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
	cursor, err := collection.Find(ctx, bson.M{ "creatorId": id }) // UUID
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
	Settings  Settings
}

type Settings struct {
	Theme string
}

var encrypter = securecookie.New(securecookie.GenerateRandomKey(64), securecookie.GenerateRandomKey(32))


func (user *User) EncodePassword() (err error) {
	encoded, err := encrypter.Encode("password", user.Password)
	if err != nil {
		return err
	}

	user.Password = encoded
	return nil
}


func (user *User) DecodePassword() (err error) {
	var decoded string
	err = encrypter.Decode("password", user.Password, &decoded)
	if err != nil {
		return err
	}

	user.Password = decoded
	return nil
}


func (user *User) GetSurveys() (surveys []Survey, err error) {
	surveys, err = GetSurveysById(user.UserId)
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

	err = user.EncodePassword()
	if err != nil {
		return user, err
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


func getUserByFilter(filter bson.M) (user User, err error) {
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

	user.DecodePassword()
	return user, nil
}


func GetUser(username string, password string) (user User, err error) {
	return getUserByFilter(bson.M{ "username": username, "password": password })
}


func GetUserById(id string) (user User, err error) {
	return getUserByFilter(bson.M{ "userId": id })
}

