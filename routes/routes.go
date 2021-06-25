package routes

import (
	"log"
	"net/http"
	"strings"
	"text/template"
)


func setEncodedCookie(res http.ResponseWriter, name string, key string, value interface{}) {
	if encoded, err := cookieHandler.Encode(key, value); err == nil {
		cookie := &http.Cookie{
			Name: name,
			Value: encoded,
			Path: "/",
		}
	
		http.SetCookie(res, cookie)
	}
}


func getUserIdFromCookie(req *http.Request) (id string) {
	var decoded string
	id = ""

	if cookie, err := req.Cookie("token"); err == nil {
		if err = cookieHandler.Decode("userId", cookie.Value, &decoded); err == nil {
			id = decoded
		}
	}

	return id
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