package redirect

import (
	"net/http"

	"gopkg.in/yaml.v2"
)

// YAMLHandler parses the provided YAML and returns an http.HandlerFunc that maps paths to URLs.
// If the path is not found in the YAML, the fallback http.Handler is called instead.
//
// Arguments:
//   - pathMap   *PathMap      - a pointer to the PathMap to populate with parsed paths
//   - yml       []byte        - YAML data in the format:
//     path: /some-path
//     url: https://www.some-url.com/demo
//   - fallback  http.Handler  - handler to call if the path is not found
//
// Returns:
//   - http.HandlerFunc - the handler function that performs the redirect logic
//   - error            - error if YAML parsing or map conversion fails
func YAMLHandler(pathMap *PathMap, filePath string, fallback http.Handler) (http.HandlerFunc, error) {
	yml, err := getContent(filePath)

	if err != nil {
		return nil, err
	}

	pathList, err := parseYaml(yml)

	if err != nil {
		return nil, err
	}

	pathMap, err = convertListToMap(pathMap, pathList)

	if err != nil {
		return nil, err
	}

	return MapHandler(pathMap, fallback), nil
}

func parseYaml(yml []byte) ([]mapItem, error) {
	var list []mapItem
	err := yaml.Unmarshal(yml, &list)
	if err != nil {
		return nil, err
	}
	return list, nil
}
