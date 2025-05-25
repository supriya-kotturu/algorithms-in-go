// Package redirect provides utilities for managing URL redirection mappings with thread-safe access.
//
// The package provides the following key functionalities:
//   - NewPathMap: Initializes and returns a new, empty PathMap.
//   - get: Retrieves the URL for a given path, returning the URL and a boolean indicating existence.
//   - set: Adds a new path-to-URL mapping if the path does not already exist, returning an error if it does.
//   - convertListToMap: Populates a PathMap from a list of mapItem structs, each representing a path-URL pair.
//
// Typical usage involves creating a PathMap, populating it with path-URL pairs, and performing thread-safe
// lookups or insertions as needed for redirect logic.
package redirect

import (
	"fmt"
	"os"
	"sync"
)

// mapItem represents a single path-to-URL mapping, typically used while unmarshalling YAML or JSON data.
type mapItem struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

// PathMap is a concurrency-safe structure for storing and retrieving mappings from paths to URLs.
// It uses a sync.RWMutex to allow multiple concurrent readers or one writer at a time, ensuring safe access
// in concurrent environments such as web servers or redirect services.
//
// PathMap contains:
//   - paths (map[string]string): that stores the mapping from a path to its corresponding URL.
//   - m (sync.RWMutex): to protect the map from concurrent access.
type PathMap struct {
	paths map[string]string
	m     sync.RWMutex
}

// NewPathMap returns an instance of PathMap
func NewPathMap() *PathMap {
	return &PathMap{
		paths: make(map[string]string),
	}
}

func (p *PathMap) get(path string) (string, bool) {
	p.m.RLock()
	defer p.m.RUnlock()

	url, ok := p.paths[path]
	return url, ok
}

func (p *PathMap) set(path string, url string) (string, error) {
	p.m.Lock()
	defer p.m.Unlock()

	if v, ok := p.paths[path]; ok {
		return v, fmt.Errorf("path %s already exists", path)
	}

	p.paths[path] = url
	return url, nil
}

func convertListToMap(pathMap *PathMap, list []mapItem) (*PathMap, error) {
	for _, k := range list {
		pathMap.set(k.Path, k.URL)
	}
	return pathMap, nil
}

func getContent(filePath string) ([]byte, error) {
	content, err := os.ReadFile(filePath)

	if err != nil {
		return []byte{}, err
	}

	return content, nil
}
