package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"text/template"

	"github.com/jesperkha/survey-app/data"
)


func SurveyHandler(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/survey" {
		http.Redirect(res, req, "/error/404", http.StatusNotFound)
		return
	}

	if req.Method == "GET" {
		if errorCode, err := SurveyGET(res, req); errorCode != 0 {
			http.Redirect(res, req, fmt.Sprintf("/error/%d", errorCode), errorCode)
			log.Print(err) // Debug
		}
	}

	if req.Method == "POST" {
		if errorCode, err := SurveyPOST(res, req); errorCode != 0 {
			res.WriteHeader(errorCode)
			log.Print(err) // Debug
		}
	}
}


func SurveyGET(res http.ResponseWriter, req *http.Request) (errorCode int, err error) {
	surveyID := req.FormValue("id")
	result, err := data.GetSurveysById(surveyID)
	if err != nil {
		return 500, err
	}

	if len(result) >= 1 {
		template, _ := template.ParseFiles("./Client/templates/survey.html")
		template.Execute(res, result[0])
		return 0, nil
	}

	return 404, err
}


type SurveyResponse struct {
	Data [][]string
	Id 	 string
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


type RequestError struct {
	Type string
	Msg  string
}

func HandleError(res http.ResponseWriter, req *http.Request) {
	errorType := strings.ReplaceAll(req.URL.Path, "/error/", "")

	reqErr := RequestError{
		Type: errorType,
	}

	tmp, err := template.ParseFiles("./Client/templates/error.html")
	if err != nil {
		log.Fatal(err)
	}

	err = tmp.Execute(res, reqErr)
	if err != nil {
		log.Fatal(err)
	}
}