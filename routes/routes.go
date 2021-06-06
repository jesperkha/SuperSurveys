package routes

import (
	"log"
	"net/http"
	"project/data"
	"text/template"
)

// Handles main page (index.html)

func Index(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "./Client/index.html")
}


// Structs to represent survey data in a usable manner

type Survey struct {
	SurveyId	 string
	Name         string
	Desc         string
	NumQuestions int
	Questions    []Question
}

type Question struct {
	Class    string
	Optional bool
	Prompt   string
	Options  []string
}


// Fetches survey by ID from database and serves it to the client

func FetchSurvey(res http.ResponseWriter, req *http.Request) {

	// Fetch document data

	surveyID := req.FormValue("id")
	var result []Survey
	data.GetDocumentById(surveyID, &result)


	// If the number of surveys returned is 0 the client will be served the 404 page
	// If a survey exists, its sent to a html template

	if len(result) != 0 {

		// Send data to html template

		template, err := template.ParseFiles("./Client/surveyTemplate.html")
		if err != nil { log.Fatal(err) }

		err = template.Execute(res, result[0])
		if err != nil { log.Fatal(err) }

	} else {

		// Send to 404 page
		http.ServeFile(res, req, "./Client/404.html")

	}

}