package main

import (
	"fmt"
	"net/http"

	"github.com/supriya-kotturu/algorithms-in-go/url-redirect/redirect"
)

func main() {
	port := 8080
	mux := http.NewServeMux()
	mux.HandleFunc("/", defaultHandler)
	mux.HandleFunc("/404", errorHandler)

	routeMap := map[string]string{
		"/google": "https://google.com",
		"/ddg":    "https://duckduckgo.com",
	}

	middleware := redirect.MapHandler(routeMap, mux)

	fmt.Printf("Running server on port : %d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), middleware)
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "test error message. try again", 400)
}
