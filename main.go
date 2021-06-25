package main

import (
	"log"
	"net/http"

	"github.com/jesperkha/survey-app/data"
	"github.com/jesperkha/survey-app/routes"
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

	// Route handlers
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		http.ServeFile(res, req, "./Client/index.html")
	})

	http.HandleFunc("/survey", routes.SurveyHandler)
	http.HandleFunc("/error/", routes.HandleError)
	http.HandleFunc("/login", routes.LoginHandler)
	http.HandleFunc("/users/", routes.UsersRouteHandler)
	
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./Client/js/"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./Client/css/"))))
	
	log.Print("Listening on port :3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}