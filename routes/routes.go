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

func Handlers() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		http.ServeFile(res, req, "./Client/index.html")
	})

	http.HandleFunc("/survey", FetchSurvey)
	http.HandleFunc("/submitSurveyData", SubmitSurveyData)
	http.HandleFunc("/error/", HandleError)
}


func FetchSurvey(res http.ResponseWriter, req *http.Request) {
	surveyID := req.FormValue("id")
	result, err := data.GetSurveysById(surveyID)
	if err != nil {
		HandleErrorCode(res, req, 500)
	}

	numSurveys := len(result)

	if numSurveys == 1 {
		template, _ := template.ParseFiles("./Client/templates/survey.html")
		template.Execute(res, result[0])
	}
	
	if numSurveys == 0 {
		// Html file missing or not in database
		HandleErrorCode(res, req, 404)
	}
} 


func SubmitSurveyData(res http.ResponseWriter, req *http.Request) {
	response, err := ioutil.ReadAll(req.Body)
	if err != nil {
		res.WriteHeader(400)
	}

	answers := string(response)
	surveyId := req.URL.Query().Get("id")

	num, err := data.InsertSubmission(surveyId, answers)
	if err != nil || num == 0 {
		res.WriteHeader(500)
	}

	debug(fmt.Sprintf("Survey data submitted: %s",  string(response)))
}


func HandleError(res http.ResponseWriter, req *http.Request) {
	errorType := strings.ReplaceAll(req.URL.Path, "/error/", "")

	reqErr := RequestError{
		Type: errorType,
		Msg: ErrorMessages[errorType],
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


func HandleErrorCode(res http.ResponseWriter, req *http.Request, code int) {
	http.Redirect(res, req, fmt.Sprintf("/error/%d", code), code)
	debug(fmt.Sprintf("Redirected with error code %d", code))
}