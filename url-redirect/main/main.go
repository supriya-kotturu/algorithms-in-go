package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/supriya-kotturu/algorithms-in-go/url-redirect/redirect"
)

func main() {
	port := 8080
	shortPathMap := redirect.NewPathMap()
	jsonFilePath, yamlFilePath, err := parseCommand()

	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", defaultHandler)
	mux.HandleFunc("/404", errorHandler)

	middleware := redirect.MapHandler(shortPathMap, mux)
	middleware, err = redirect.YAMLHandler(shortPathMap, yamlFilePath, middleware)

	if err != nil {
		log.Println(err)
	}

	middleware, err = redirect.JSONHandler(shortPathMap, jsonFilePath, middleware)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Running server on port : %d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), middleware)
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "test error message. try again", 400)
}
