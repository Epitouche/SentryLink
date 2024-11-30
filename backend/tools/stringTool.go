package tools

type stringTool interface {
	RemoveCharFromString(s string) string
}

func RemoveCharFromString(s string) string {
	result := ""
	// remove \ in all the string
	for _, char := range s {
		if char != '\\' {
			result += string(char)
		}
	}
	return result
}
