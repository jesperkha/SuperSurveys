package routes

import (
	"net/http"

	"github.com/jesperkha/survey-app/data"

	"github.com/gorilla/securecookie"
)


var KEY64 string = "9ckXVq5lewP2ICRBzNaIeDwrXcWSWzMPCI1GlDxaMBVnaXq4T9Sgu6sf5CD1tdXo"
var KEY32 string = "Kk26D48IQrjxo1SfKmNXMNFECCwCAStu"
var cookieHandler = securecookie.New([]byte(KEY64), []byte(KEY32))

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
			setEncodedCookie(res, "token", "userId", user.UserId)
			http.Redirect(res, req, "/users/dashboard", http.StatusFound)
		}

		return
	}
	
	http.Error(res, "Bad Request", http.StatusBadRequest)
}


func LogoutHandler(res http.ResponseWriter, req *http.Request) {
	// remove cookie
}


func Authorize(req *http.Request) (user data.User, authorized bool) {
	if user, err := data.GetUserById(getUserIdFromCookie(req)); err == nil {
		return user, true
	}

	return user, false
}