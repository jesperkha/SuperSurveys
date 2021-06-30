package routes

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"

	"github.com/jesperkha/SuperSurveys/app/data"
)

func SurveyHandler(res http.ResponseWriter, req *http.Request) (errorCode int) {
	if req.URL.Path != "/survey" {
		return 404
	}

	if req.Method == "GET" {
		errorCode, err := SurveyGET(res, req)
		// Debug
		if err != nil {
			log.Print(err)
		}
		return errorCode
	}

	if req.Method == "POST" {
		errorCode, err := SurveyPOST(res, req)
		// Debug
		if err != nil {
			log.Print(err)
		}
		return errorCode
	}

	return 400
}


func SurveyGET(res http.ResponseWriter, req *http.Request) (errorCode int, err error) {
	surveyID := req.FormValue("id")
	result, err := data.GetSurveysById(surveyID)
	if err != nil {
		return 500, err
	}

	if len(result) >= 1 {
		template, _ := template.ParseFiles("./client/templates/survey.html")
		template.Execute(res, result[0])
		return 0, nil
	}

	return 404, err
}


func SurveyPOST(res http.ResponseWriter, req *http.Request) (errorCode int, err error) {
	response, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return 500, err
	}

	var submission SurveyResponse
	if err = json.Unmarshal(response, &submission); err != nil {
		return 500, err
	}

	if num, err := data.InsertSubmission(submission.Id, submission.Data); err != nil || num == 0 {
		return 500, err
	}

	res.WriteHeader(http.StatusOK)
	log.Print(submission) // Debug
	return 0, nil
}