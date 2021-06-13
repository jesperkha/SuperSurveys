package main

import (
	"log"
	"net/http"

	"github.com/jesperkha/survey-app/data"
	"github.com/jesperkha/survey-app/routes"
	"github.com/joho/godotenv"
)


func main() {
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
	
	routes.Handlers()
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./Client/js/"))))
	log.Print("Listening on port :3000")
	http.ListenAndServe(":3000", nil)
}