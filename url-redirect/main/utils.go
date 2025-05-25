package main

import (
	"errors"
	"flag"
	"os"
)

func parseCommand() (string, string, error) {
	json := flag.String("json", "../data/map.json", "The file path to the url mappings in JSON format.")
	yaml := flag.String("yaml", "../data/map.yml", "The file path to the url mappings in YAML format.")

	flag.Parse()

	_, jsonErr := os.Stat(*json)
	_, yamlErr := os.Stat(*yaml)

	if jsonErr == nil && yamlErr == nil {
		return *json, *yaml, nil
	}

	jsonNotExists := os.IsNotExist(jsonErr)
	yamlNotExists := os.IsNotExist(yamlErr)

	if jsonNotExists && yamlNotExists {
		return "", "", errors.New("files not found: \n" + *json + "\n" + *yaml)
	} else if jsonNotExists {
		return "", *yaml, errors.New("json file not found: " + *json)
	} else if yamlNotExists {
		return *json, "", errors.New("yaml file not found: " + *yaml)
	} else {
		return "", "", errors.New("cannot find the files")
	}
}
