package routes

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"

	"github.com/gorilla/securecookie"
	"github.com/jesperkha/SuperSurveys/app/data"
)

// Secure cookies for auth

var CookieHandler *securecookie.SecureCookie


// Error type for error page handler

type RequestError struct {
	Type string
	Msg  string
}


// Response type coming from client upon submitting a form

type SurveyResponse struct {
	Data [][]string
	Id   string
}


// Handlers for incomming requests

type HandlerFunc func(res http.ResponseWriter, req *http.Request) (errorCode int)

var routeHandlers = map[string]HandlerFunc {
	"login": LoginHandler,
	"survey": SurveyHandler,
	"users": UsersHandler,
}


// Handlers for endpoints starting with /users

type UserHandlerFunc func(res http.ResponseWriter, req *http.Request, user data.User) (errorCode int)

var userHandlers = map[string]UserHandlerFunc {
	"dashboard": PageUserDashboard,
	"create": PageCreateSurvey,
	"get": ServeUserProfile,
}


func RouteHandler(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/" {
		http.ServeFile(res, req, "./client/index.html")
		return
	}
	
	split := strings.Split(req.URL.Path, "/")
	route := split[1]
	if handler, ok := routeHandlers[route]; ok {
		if errorCode := handler(res, req); errorCode != 0 {
			http.Redirect(res, req, fmt.Sprintf("/error/%d", errorCode), http.StatusMovedPermanently)
		}
	} else {
		http.Redirect(res, req, "/error/404", http.StatusNotFound)
	}
}


func HandleError(res http.ResponseWriter, req *http.Request) {
	errorType := strings.ReplaceAll(req.URL.Path, "/error/", "")
	reqErr := RequestError{ Type: errorType }

	tmp, err := template.ParseFiles("./client/templates/error.html")
	if err != nil {
		log.Print(err)
	}

	err = tmp.Execute(res, reqErr)
	if err != nil {
		log.Print(err)
	}
}