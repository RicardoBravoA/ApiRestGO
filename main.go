package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(request http.Request, response http.ResponseWriter){
		fmt.Fprintf(response, "Hello world from my server with GO")
	})
}