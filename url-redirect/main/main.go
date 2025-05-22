package main

import (
	"net/http"

	"github.com/supriya-kotturu/algorithms-in-go/url-redirect/redirect"
)

func main() {
	routeMap := make(map[string]string)

	routeMap["/google"] = "https://google.com"
	routeMap["/ddg"] = "https://duckduckgo.com"

	fallbackHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Not found!"))
	})

	// mux := http.NewServeMux()
	handler := redirect.MapHandler(routeMap, fallbackHandler)

	http.ListenAndServe(":8080", handler)
}
