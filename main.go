package main

import (
	"log"
	"net/http"
	"project/data"
	"project/routes"
)

const Port string = ":3000"

func main() {

	data.ConnectClient()
	defer data.CloseClient()

	// HTTP Server

	http.HandleFunc("/", routes.Index)
	http.HandleFunc("/survey", routes.FetchSurvey)

	err := http.ListenAndServe(Port, nil)
	if err != nil { log.Fatal(err) }
}