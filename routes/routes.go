package routes

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"

	"github.com/gorilla/securecookie"
	"github.com/jesperkha/survey-app/data"
)

// Secure cookies for auth

var KEY64 string = "9ckXVq5lewP2ICRBzNaIeDwrXcWSWzMPCI1GlDxaMBVnaXq4T9Sgu6sf5CD1tdXo"
var KEY32 string = "Kk26D48IQrjxo1SfKmNXMNFECCwCAStu"
var cookieHandler = securecookie.New([]byte(KEY64), []byte(KEY32))


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

var RouteHandlers = map[string]HandlerFunc {
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
		http.ServeFile(res, req, "./Client/index.html")
		return
	}
	
	split := strings.Split(req.URL.Path, "/")
	route := split[1]
	if handler, ok := RouteHandlers[route]; ok {
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

	tmp, err := template.ParseFiles("./Client/templates/error.html")
	if err != nil {
		log.Print(err)
	}

	err = tmp.Execute(res, reqErr)
	if err != nil {
		log.Print(err)
	}
}