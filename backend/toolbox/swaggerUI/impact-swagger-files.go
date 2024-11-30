package swaggerui

import (
	"encoding/json"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"

	"github.com/Tom-Mendy/SentryLink/schemas"
)

func impactSwaggerFiles(routes []schemas.Route) {
	var filePathOfFiles = []string{
		"docs/docs.go",
		"docs/swagger.json",
		"docs/swagger.yaml",
	}
	for _, route := range routes {
		for _, file := range filePathOfFiles {
			processFile(file, route)
		}
	}
}

func processFile(filePath string, route schemas.Route) {



	fileData, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file %s: %s\n", filePath, err)
		return
	}

	var paths map[string]interface{}

	if isGOFile(filePath) {
		jsonValueInGoFile, err := updateDocTemplate(filePath)
		err = json.Unmarshal([]byte(jsonValueInGoFile), &paths)
		if err != nil {
			fmt.Printf("Error reading file %s: %s\n", filePath, err)
			return
		}
	} else if isJSONFile(filePath) {
		err = json.Unmarshal(fileData, &paths)
		if err != nil {
			fmt.Printf("Error unmarshalling JSON file %s: %s\n", filePath, err)
			return
		}
	} else if isYAMLFile(filePath) {
		err = yaml.Unmarshal(fileData, &paths)
		if err != nil {
			fmt.Printf("Error unmarshalling YAML file %s: %s\n", filePath, err)
			return
		}
	} else {
		fmt.Printf("Unsupported file type %s\n", filePath)
		return
	}

	if paths == nil {
		paths = make(map[string]interface{})
	}
	if _, ok := paths["paths"]; !ok {
		paths["paths"] = make(map[string]interface{})
	}
	pathsMap := paths["paths"].(map[string]interface{})
	pathsMap[route.Path] = buildRouteEntry(route)

	if isGOFile(filePath) {
		_, err := json.MarshalIndent(paths, "", "  ")
		if err != nil {
			fmt.Printf("Error serializing JSON for file %s: %v\n", filePath, err)
			return
		}
		newActualFilePath := "tmp.json"
		err = updateDocTemplateWithJSON(filePath, newActualFilePath)
		if err != nil {
			fmt.Printf("Error updating docTemplate in file %s: %v\n", filePath, err)
			return
		}

	} else if isJSONFile(filePath) {
		updatedJSON, err := json.MarshalIndent(paths, "", "  ")
		if err != nil {
			fmt.Printf("Error serializing JSON for file %s: %v\n", filePath, err)
			return
		}

		err = os.WriteFile(filePath, updatedJSON, 0644)
		if err != nil {
			fmt.Printf("Error writing JSON to file %s: %v\n", filePath, err)
			return
		}
	} else if isYAMLFile(filePath) {
		updatedYAML, err := yaml.Marshal(paths)
		if err != nil {
			fmt.Printf("Error serializing YAML for file %s: %v\n", filePath, err)
			return
		}

		err = os.WriteFile(filePath, updatedYAML, 0644)
		if err != nil {
			fmt.Printf("Error writing YAML to file %s: %v\n", filePath, err)
			return
		}
	}

	fmt.Printf("Route added successfully to %s\n", filePath)
}
