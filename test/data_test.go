package test

import (
	"testing"

	"github.com/jesperkha/SuperSurveys/app/data"
	"github.com/joho/godotenv"
)

var (
	testSurveyId = "e0b7bf26db9d4bb2adfd52582e22844c"
)

func TestClientConnect(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		t.Error(err)
	}
	err = data.ConnectClient()
	if err != nil {
		t.Error(err)
	}
}

func TestGetSurvey(t *testing.T) {
	result, err := data.GetSurveysById(testSurveyId)
	if err != nil {
		t.Error(err)
	}
	if len(result) != 1 {
		t.Error("Too many documents fetched")
	}
}

func TestInsertSubmission(t *testing.T) {
	num, err := data.InsertSubmission(testSurveyId, [][]string{{"TEST_SUBMISSION"}})
	if err != nil {
		t.Error(err)
	}
	if num != 1 {
		t.Error("Multiple documents modified")
	}
}

func TestGetUser(t *testing.T) {
	user, err := data.GetUser("TestUser", "1234")
	if err != nil {
		t.Error(err)
	}
	if user.Username != "TestUser" {
		t.Error("Wrong user fetched")
	}
}

func TestClientDisconnect(t *testing.T) {
	data.CloseClient()
}

