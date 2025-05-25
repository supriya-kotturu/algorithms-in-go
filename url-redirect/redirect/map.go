package redirect

import "net/http"

// MapHandler returns an http.HandlerFunc that attempts to map request paths to their corresponding URLs.
//
// Definition:
//   MapHandler checks if the incoming request path exists in the provided PathMap. If a match is found,
//   it redirects the request to the mapped URL. Otherwise, it delegates the request to the fallback handler.
//
// Arguments:
//   - pathMap (*PathMap): A pointer to a PathMap that holds path-to-URL mappings.
//   - fallback (ttp.Handler): A fallback handler to invoke if the path is not found in pathMap.
//
// Returns:
//   - http.HandlerFunc: An HTTP handler function that performs the described mapping and fallback logic.
func MapHandler(pathMap *PathMap, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Path
		dest, ok := pathMap.get(url)

		if ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}
