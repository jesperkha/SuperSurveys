package routes

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
	"text/template"

	"github.com/jesperkha/SuperSurveys/app/data"
)


func UsersHandler(res http.ResponseWriter, req *http.Request) (errorCode int) {
	path := strings.Split(req.URL.Path, "/")
	if len(path) != 3 || path[1] != "users" {
		return 404
	}

	route := path[2]

	if user, auth := Authorize(req); auth {
		if f, ok := userHandlers[route]; ok {
			return f(res, req, user)
		} else {
			return 404
		}
	}

	log.Print("Not authorized") // Debug
	http.Redirect(res, req, "/login", http.StatusUnauthorized)
	return 0
}


func PageUserDashboard(res http.ResponseWriter, req *http.Request, user data.User) (errorCode int) {
	template, err := template.ParseFiles("./client/templates/dashboard.html")
	if err != nil {
		return 500
	}

	err = template.Execute(res, user)
	if err != nil {
		return 500
	}

	return 0
}


func ServeUserProfile(res http.ResponseWriter, req *http.Request, user data.User) (errorCode int) {
	if req.Method != "GET" {
		return 400
	}

	userJson, err := json.Marshal(user)
	if err != nil {
		return 500
	}

	res.Write([]byte(userJson))
	return 0
}


func PageCreateSurvey(res http.ResponseWriter, req *http.Request, user data.User) (errorCode int) {
	if req.Method == "GET" {
		http.ServeFile(res, req, "./client/create_survey.html")
		return 0
	}

	if req.Method == "POST" {
		raw, err := io.ReadAll(req.Body)
		if err != nil {
			return 500
		}

		var parsed []data.Question
		err = json.Unmarshal(raw, &parsed)
		if err != nil {
			return 400
		}

		// Todo
		// Make submission to database with questions
		// also get the survey title and description with request
		log.Print(parsed)
		return 0
	}

	return 0
}