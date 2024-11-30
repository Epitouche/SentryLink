package swaggerui

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"

	"github.com/Tom-Mendy/SentryLink/schemas"
)

type SwaggerFile interface {
	ResolvePath(relativePath string) string
	ImpactSwaggerFiles(routes []schemas.Route)
	ProcessFile(filePath string, route schemas.Route)
}

func ResolvePath(relativePath string) string {
	basePath, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return filepath.Join(basePath, relativePath)
}

func ImpactSwaggerFiles(routes []schemas.Route) {
	var filePathOfFiles = []string{
		ResolvePath("docs/docs.go"),
		ResolvePath("docs/swagger.json"),
		ResolvePath("docs/swagger.yaml"),
	}
	for _, route := range routes {
		for _, file := range filePathOfFiles {
			ProcessFile(file, route)
		}
	}
}

func ProcessFile(filePath string, route schemas.Route) {

	fileData, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file %s: %s\n", filePath, err)
		return
	}

	var paths map[string]interface{}
	var yamlPath interface{}

	if IsGOFile(filePath) {
		_, err := UpdateDocTemplate(filePath)
		// tmpFileData, err := os.ReadFile("tmp.json")
		// err = json.Unmarshal(tmpFileData, &paths)
		if err != nil {
			fmt.Printf("Error reading file %s: %s\n", filePath, err)
			return
		}
	} else if IsJSONFile(filePath) {
		err = json.Unmarshal(fileData, &paths)
		if err != nil {
			fmt.Printf("Error unmarshalling JSON file %s: %s\n", filePath, err)
			return
		}
	} else if IsYAMLFile(filePath) {
		err = yaml.Unmarshal(fileData, &yamlPath)
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
	pathsMap[route.Path] = BuildRouteEntry(route)

	if IsGOFile(filePath) {
		_, err := json.MarshalIndent(paths, "", "  ")
		if err != nil {
			fmt.Printf("Error serializing JSON for file %s: %v\n", filePath, err)
			return
		}
		newActualFilePath := "tmp.json"
		err = UpdateDocTemplateWithJSON(filePath, newActualFilePath)
		if err != nil {
			fmt.Printf("Error updating docTemplate in file %s: %v\n", filePath, err)
			return
		}

	} else if IsJSONFile(filePath) {
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
	} else if IsYAMLFile(filePath) {
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
