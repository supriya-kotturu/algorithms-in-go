package main

import (
	"net/http"

	"github.com/supriya-kotturu/algorithms-in-go/url-redirect/redirect"
)

type URLRedirectDeps interface {
	parseCommand() (string, string, error)

	defaultHandler(http.ResponseWriter, *http.Request)
	errorHandler(http.ResponseWriter, *http.Request)

	mapMiddleware(*redirect.PathMap, http.Handler) http.HandlerFunc
	yamlMiddleware(*redirect.PathMap, string, http.HandlerFunc) (http.HandlerFunc, error)
	jsonMiddleware(*redirect.PathMap, string, http.HandlerFunc) (http.HandlerFunc, error)
}

type URLRedirect struct {
	pathMap *redirect.PathMap
	deps    URLRedirectDeps
}

type DefaultURLRedirectDeps struct{}

func (d *DefaultURLRedirectDeps) parseCommand() (string, string, error) {
	return parseCommand()
}

func (d *DefaultURLRedirectDeps) defaultHandler(w http.ResponseWriter, r *http.Request) {
	defaultHandler(w, r)
}

func (d *DefaultURLRedirectDeps) errorHandler(w http.ResponseWriter, r *http.Request) {
	errorHandler(w, r)
}

func (d *DefaultURLRedirectDeps) mapMiddleware(m *redirect.PathMap, h http.Handler) http.HandlerFunc {
	return redirect.MapHandler(m, h)
}

func (d *DefaultURLRedirectDeps) yamlMiddleware(m *redirect.PathMap, p string, h http.HandlerFunc) (http.HandlerFunc, error) {
	return redirect.YAMLHandler(m, p, h)
}

func (d *DefaultURLRedirectDeps) jsonMiddleware(m *redirect.PathMap, p string, h http.HandlerFunc) (http.HandlerFunc, error) {
	return redirect.JSONHandler(m, p, h)
}

func NewURLRedirect(deps URLRedirectDeps) *URLRedirect {
	pathMap := redirect.NewPathMap()

	if deps == nil {
		return &URLRedirect{
			pathMap: pathMap,
			deps:    &DefaultURLRedirectDeps{},
		}
	}

	return &URLRedirect{
		pathMap: pathMap,
		deps:    deps,
	}
}
