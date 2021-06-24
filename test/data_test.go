package test

import (
	"testing"

	"github.com/jesperkha/survey-app/data"
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
	_, err := data.GetSurveysById(testSurveyId)
	if err != nil {
		t.Error(err)
	}
}

func TestInsertSubmission(t *testing.T) {
	_, err := data.InsertSubmission(testSurveyId, [][]string{{"TEST_SUBMISSION"}})
	if err != nil {
		t.Error(err)
	}
}

func TestGetUser(t *testing.T) {
	_, err := data.GetUser("TestUser", "1234")
	if err != nil {
		t.Error(err)
	}
}

func TestClientDisconnect(t *testing.T) {
	data.CloseClient()
}

