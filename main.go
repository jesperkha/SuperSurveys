package main

import (
	"fmt"
	"net/http"
)



func main() {
	fs := http.FileServer(http.Dir("./Client"))

	err := http.ListenAndServe(":3000", fs)
	if err != nil {
		fmt.Printf("Listen err: %s", err)
	}
}