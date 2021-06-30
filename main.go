package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/securecookie"
	"github.com/jesperkha/SuperSurveys/app/data"
	"github.com/jesperkha/SuperSurveys/app/routes"
	"github.com/joho/godotenv"
)


func main() {
	// Load env vars
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Mongo client connect
	err = data.ConnectClient()
	if err != nil {
		log.Fatal(err)
	}
	
	log.Print("MongoDB client connected")
	defer data.CloseClient()

	// Set cookie handler with env vars
	routes.CookieHandler = securecookie.New([]byte(os.Getenv("KEY64")), []byte(os.Getenv("KEY32")))
	
	// Route handlers
	http.HandleFunc("/", routes.RouteHandler)
	// http.HandleFunc("/error/", routes.HandleError)

	var filePrefixes = map[string]string {
		"/js/": "./client/js/",
		"/css/": "./client/css/",
	}

	for key, value := range filePrefixes {
		http.Handle(key, http.StripPrefix(key, http.FileServer(http.Dir(value))))
	}
	
	log.Print("Listening on port :3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}