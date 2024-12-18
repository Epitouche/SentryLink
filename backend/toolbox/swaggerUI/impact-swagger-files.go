package swaggerui

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"

	"github.com/Tom-Mendy/SentryLink/schemas"
)

type SwaggerFile interface {
	ResolvePath(relativePath string) string
	ImpactSwaggerFiles(routes []schemas.Route, routesToRemove []string)
	ProcessFile(filePath string, route schemas.Route)
	ExtractExistingRoutes(filePath string) []string
	FindRoutesToRemove(existingRouteInFile []string, currentRoutes []string) []string
}

func FindRoutesToRemove(existingRouteInFile []string, currentRoutes []string) []string {
	routesToRemove := []string{}
	for _, existingRoute := range existingRouteInFile {
		if !strings.Contains(strings.Join(currentRoutes, ""), existingRoute) {
			routesToRemove = append(routesToRemove, existingRoute)
		}
	}
	return routesToRemove
}

func ExtractExistingRoutes(filePath string) []string {
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file %s: %s\n", filePath, err)
		return nil
	}

	var paths map[string]interface{}
	if IsJSONFile(filePath) {
		err = json.Unmarshal(fileData, &paths)
		if err != nil {
			fmt.Printf("Error unmarshalling JSON file %s: %s\n", filePath, err)
			return nil
		}
	}
	existingRoutes := []string{}
	for path := range paths["paths"].(map[string]interface{}) {
		existingRoutes = append(existingRoutes, path)
	}
	fmt.Printf("Existing routes: %++v\n", existingRoutes)
	return existingRoutes
}

func ResolvePath(relativePath string) string {
	basePath, _ := filepath.Abs(filepath.Dir(os.Args[1]))
	return filepath.Join(basePath, relativePath)
}

func removeRoutesFromFile(filePath string, staleRoutes []string) {
	fmt.Printf("Removing the following routes from %s: %++v\n", filePath, staleRoutes)
	// Remove the route and its content from the file (the entire object from the paths object)
	if len(staleRoutes) == 0 {
		return
	}
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file %s: %s\n", filePath, err)
		return
	}

	var paths map[string]interface{}
	if IsJSONFile(filePath) {
		err = json.Unmarshal(fileData, &paths)
		if err != nil {
			fmt.Printf("Error unmarshalling JSON file %s: %s\n", filePath, err)
			return
		}
	} else if IsGOFile(filePath) {
		_, err := UpdateDocTemplate(filePath)
		if err != nil {
			fmt.Printf("Error reading file %s: %s\n", filePath, err)
			return
		}
		newActualFilePath := "tmp.json"
		newFileData, err := os.ReadFile(newActualFilePath)
		err = json.Unmarshal(newFileData, &paths)
		if err != nil {
			fmt.Printf("Error unmarshalling JSON file %s: %s\n", newActualFilePath, err)
			return
		}
		fmt.Printf("JSON TMP FILE READ\n")
		fmt.Printf("%++v\n", paths)
	}

	pathsMap := paths["paths"].(map[string]interface{})
	for _, staleRoute := range staleRoutes {
		delete(pathsMap, staleRoute)
	}
	if IsGOFile(filePath) {
		updatedJSON, err := json.MarshalIndent(paths, "", "  ")
		if err != nil {
			fmt.Printf("Error serializing JSON for file %s: %v\n", filePath, err)
			return
		}
		newActualFilePath := "tmp.json"
		err = os.WriteFile(newActualFilePath, updatedJSON, 0644)
		if err != nil {
			fmt.Printf("Error writing JSON to file %s: %v\n", filePath, err)
			return
		}
		err = UpdateDocTemplateWithJSON(filePath, newActualFilePath, paths)
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
	}
	fmt.Printf("Routes removed successfully from %s\n", filePath)

}

func ImpactSwaggerFiles(routes []schemas.Route, routesToRemove []string) {
	var filePathOfFiles = []string{
		ResolvePath("docs/swagger.json"),
		// ResolvePath("docs/swagger.yaml"),
		ResolvePath("docs/docs.go"),
	}
	for _, file := range filePathOfFiles {
		removeRoutesFromFile(file, routesToRemove)
		for _, route := range routes {
			ProcessFile(file, route)
		}
	}
	os.Remove("tmp.json")
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
		newActualFilePath := "tmp.json"
		newFileData, err := os.ReadFile(newActualFilePath)
		err = json.Unmarshal(newFileData, &paths)
		if err != nil {
			fmt.Printf("Error unmarshalling JSON file %s: %s\n", newActualFilePath, err)
			return
		}
		// err = UpdateDocTemplateWithJSON(filePath, newActualFilePath, paths)
		// if err != nil {
		// 	fmt.Printf("Error updating docTemplate in file %s: %v\n", filePath, err)
		// 	return
		// }
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
		err = UpdateDocTemplateWithJSON(filePath, newActualFilePath, paths)
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
