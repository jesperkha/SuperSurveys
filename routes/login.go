package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jesperkha/survey-app/data"

	"github.com/gorilla/securecookie"
)


var KEY64 string = "9ckXVq5lewP2ICRBzNaIeDwrXcWSWzMPCI1GlDxaMBVnaXq4T9Sgu6sf5CD1tdXo"
var KEY32 string = "Kk26D48IQrjxo1SfKmNXMNFECCwCAStu"
var cookieHandler = securecookie.New([]byte(KEY64), []byte(KEY32))
var cookieName = "token"
var cookieValue = "userId"


func setCookie(res http.ResponseWriter, id string) {
	if encoded, err := cookieHandler.Encode(cookieValue, id); err == nil {
		cookie := &http.Cookie{
			Name: cookieName,
			Value: encoded,
		}
	
		http.SetCookie(res, cookie)
	} else {
		// Proto
		log.Print("Error encoding login.go")
	}
}


func LoginHandler(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/login" {
		http.Error(res, "404 Not Found", http.StatusNotFound)
		return
	}

	if req.Method == "GET" {
		http.ServeFile(res, req, "./Client/login.html")
		return
	}

	if req.Method == "POST" {
		username := req.FormValue("username")
		password := req.FormValue("password")

		if user, err := data.GetUser(username, password); err == nil {
			setCookie(res, user.UserId)
			url := fmt.Sprintf("/users/%s/dashboard", user.UserId)
			http.Redirect(res, req, url, http.StatusFound)
		}

		return
	}
	
	http.Error(res, "Bad Request", http.StatusBadRequest)
}


func LogoutHandler(res http.ResponseWriter, req *http.Request) {
	// remove cookie
}


func getUserId(req *http.Request) (id string) {
	var decoded string

	if cookie, err := req.Cookie(cookieName); err == nil {
		if err = cookieHandler.Decode(cookieValue, cookie.Value, &decoded); err == nil {
			id = decoded
		}
	} else {
		log.Print("Error decoding login.go")
	}

	return id
}


func Authorize(req *http.Request) (user data.User, authorized bool) {
	if user, err := data.GetUserById(getUserId(req)); err == nil {
		return user, true
	}

	return user, false
}