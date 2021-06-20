package routes

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"text/template"

	"github.com/jesperkha/survey-app/data"
)

// Proto
func debug(str string) {
	log.Print(str)
}


func SurveyHandler(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/survey" {
		http.Redirect(res, req, "/error/404", http.StatusNotFound)
		return
	}

	if req.Method == "GET" {
		SurveyGET(res, req)
		return
	}

	if req.Method == "POST" {
		SurveyPOST(res, req)
		return
	}

	http.Error(res, "Bad Request", http.StatusBadRequest)
}


func SurveyGET(res http.ResponseWriter, req *http.Request) {
	surveyID := req.FormValue("id")
	result, err := data.GetSurveysById(surveyID)
	if err != nil {
		http.Redirect(res, req, "/error/500", http.StatusInternalServerError)
		return
	}

	numSurveys := len(result)

	if numSurveys == 1 {
		template, _ := template.ParseFiles("./Client/templates/survey.html")
		template.Execute(res, result[0])
	}
	
	if numSurveys == 0 {
		http.Redirect(res, req, "/error/404", http.StatusNotFound)
	}
}


func SurveyPOST(res http.ResponseWriter, req *http.Request) {
	response, err := ioutil.ReadAll(req.Body)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	answers := string(response)
	surveyId := req.URL.Query().Get("id")

	if num, err := data.InsertSubmission(surveyId, answers); err != nil || num == 0 {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusOK)
	debug(fmt.Sprintf("Survey data submitted: %s",  string(response)))
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