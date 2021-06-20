package routes

import (
	"log"
	"net/http"
	"strings"
	"text/template"

	"github.com/jesperkha/survey-app/data"
)

type UserHandlerFunc func(res http.ResponseWriter, req *http.Request, user data.User)

var userHandlers = map[string]UserHandlerFunc{
	"dashboard": HandleUserDashboard,
}


func UserHandler(res http.ResponseWriter, req *http.Request) {
	path := strings.Split(req.URL.Path, "/")
	if len(path) != 4 {
		http.Redirect(res, req, "/error/404", http.StatusNotFound)
		return
	}

	userId := path[2]
	route := path[3]

	if user, auth := Authorize(req); user.UserId == userId && auth {
		if f, ok := userHandlers[route]; ok {
			f(res, req, user)
		} else {
			http.Redirect(res, req, "/error/404", http.StatusNotFound)
		}

		return
	}

	http.Redirect(res, req, "/login", http.StatusUnauthorized)
}


func HandleUserDashboard(res http.ResponseWriter, req *http.Request, user data.User) {
	template, err := template.ParseFiles("./Client/templates/dashboard.html")
	if err != nil {
		log.Fatal(err)
	}

	err = template.Execute(res, user)
	if err != nil {
		log.Fatal(err)
	}
}