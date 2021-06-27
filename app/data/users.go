package data

import (
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

// Todo
// Make actual encryption and decryption

func EncodePassword(password string) (encoded string, err error) {

	return encoded, nil
}


func DecodePassword(password string) (decoded string, err error) {

	return decoded, nil
}


func (user *User) GetSurveys() (surveys []Survey, err error) {
	surveys, err = GetSurveysById(user.UserId)
	if err != nil {
		return surveys, err
	}

	return surveys, nil
}


func InsertUser(username string, password string, email string) (user User, err error) {
	encoded, err := EncodePassword(user.Password)
	if err != nil {
		return user, err
	}

	user = User{
		UserId:    uuid.NewString(),
		Username:  username,
		Password:  encoded,
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

	return user, nil
}


func GetUser(username string, password string) (user User, err error) {
	return getUserByFilter(bson.M{ "username": username, "password": password })
}


func GetUserById(id string) (user User, err error) {
	return getUserByFilter(bson.M{ "userId": id })
}

