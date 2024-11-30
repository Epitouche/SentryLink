package swaggerui

func isJSONFile(filePath string) bool {
	return len(filePath) > 5 && filePath[len(filePath)-5:] == ".json"
}

func isYAMLFile(filePath string) bool {
	return len(filePath) > 5 && filePath[len(filePath)-5:] == ".yaml"
}

func isGOFile(filePath string) bool {
	return len(filePath) > 3 && filePath[len(filePath)-3:] == ".go"
}
