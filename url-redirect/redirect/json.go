package redirect

import (
	"encoding/json"
	"net/http"
)

// JSONHandler sets up an http.HandlerFunc for redirecting paths based on a JSON configuration.
//
// Definition:
//
//	JSONHandler parses a JSON byte slice describing path-to-URL mappings, updates the provided PathMap,
//	and returns an http.HandlerFunc that redirects requests based on these mappings. If a path is not found,
//	the fallback http.Handler is used.
//
// Arguments:
//   - pathMap (*PathMap): Pointer to the PathMap to be updated with new mappings.
//   - jsn ([]byte): JSON byte slice containing an array of path-to-URL mappings.
//   - fallback (http.Handler): Handler to call when a path is not found in the map.
//
// Returns:
//   - (http.HandlerFunc): Handler function that performs the redirection.
//   - (error): Error encountered during JSON parsing or map update, if any.
func JSONHandler(pathMap *PathMap, filePath string, fallback http.Handler) (http.HandlerFunc, error) {
	jsn, err := getContent(filePath)

	if err != nil {
		return nil, err
	}

	pathList, err := parseJson(jsn)

	if err != nil {
		return nil, err
	}

	pathMap, err = convertListToMap(pathMap, pathList)

	if err != nil {
		return nil, err
	}

	return MapHandler(pathMap, fallback), nil
}

func parseJson(jsn []byte) ([]mapItem, error) {
	var list []mapItem

	err := json.Unmarshal(jsn, &list)

	if err != nil {
		return nil, err
	}

	return list, nil
}
