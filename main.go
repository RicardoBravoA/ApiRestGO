package main

import (
	"fmt"
	"net/http"
	"log"
)

func main() {
	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request,){
		fmt.Fprintf(response, "Hello world from my server with GO")
	})

	server := http.ListenAndServe(":8080", nil)

	log.Fatal(server)
}