package main 

import (
	"log"
	"net/http"
)

func main() {
	router := NewRouter()

	NewDBConnection()

	log.Fatal(http.ListenAndServe(":8080", router))	
}