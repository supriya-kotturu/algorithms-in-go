package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := 8080

	deps := &DefaultURLRedirectDeps{}
	u := NewURLRedirect(deps)

	jsonFilePath, yamlFilePath, err := u.deps.parseCommand()

	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", u.deps.defaultHandler)
	mux.HandleFunc("/404", u.deps.errorHandler)

	middleware := u.deps.mapMiddleware(u.pathMap, mux)
	middleware, err = u.deps.yamlMiddleware(u.pathMap, yamlFilePath, middleware)

	if err != nil {
		log.Println(err)
	}

	middleware, err = u.deps.jsonMiddleware(u.pathMap, jsonFilePath, middleware)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Running server on port : %d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), middleware)
}
