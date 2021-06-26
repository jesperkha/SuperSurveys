package routes

import (
	"net/http"

	"github.com/jesperkha/survey-app/data"
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


func Authorize(req *http.Request) (user data.User, authorized bool) {
	if user, err := data.GetUserById(getUserIdFromCookie(req)); err == nil {
		return user, true
	}

	return user, false
}


func LoginHandler(res http.ResponseWriter, req *http.Request) (errorCode int) {
	if req.URL.Path != "/login" {
		return 404
	}

	if req.Method == "GET" {
		http.ServeFile(res, req, "./Client/auth/login.html")
		return 0
	}

	if req.Method == "POST" {
		username := req.FormValue("username")
		password := req.FormValue("password")

		if user, err := data.GetUser(username, password); err == nil {
			setEncodedCookie(res, "token", "userId", user.UserId)
			http.Redirect(res, req, "/users/dashboard", http.StatusFound)
		}

		return 0
	}
	
	return 400
}


func LogoutHandler(res http.ResponseWriter, req *http.Request) (errorCode int) {
	// remove cookie
	return 0
}


func SignupHandler(res http.ResponseWriter, req *http.Request) (errorCode int) {
	// create user account
	return 0
}

