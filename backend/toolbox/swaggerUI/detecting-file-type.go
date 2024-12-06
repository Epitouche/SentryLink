package swaggerui

type SwaggerDetectedFileType interface {
	IsJSONFile(filePath string) bool
	IsYAMLFile(filePath string) bool
	IsGOFile(filePath string) bool
}

func IsJSONFile(filePath string) bool {
	return len(filePath) > 5 && filePath[len(filePath)-5:] == ".json"
}

func IsYAMLFile(filePath string) bool {
	return len(filePath) > 5 && filePath[len(filePath)-5:] == ".yaml"
}

func IsGOFile(filePath string) bool {
	return len(filePath) > 3 && filePath[len(filePath)-3:] == ".go"
}
