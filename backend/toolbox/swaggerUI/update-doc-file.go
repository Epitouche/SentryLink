package swaggerui

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
	"regexp"
	"strings"
)

type SwaggerUpdatedDocFile interface {
	UpdateDocTemplate(filePath string) (string, error)
	RemoveSpecialLines(rawValue string) string
	UpdateDocTemplateWithJSON(filePath, tmpFilePath string) error
}

func UpdateDocTemplate(filePath string) (string, error) {
	fmt.Printf("Processing file: %s\n", filePath)
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filePath, nil, parser.AllErrors)
	if err != nil {
		log.Fatalf("Failed to parse file: %v", err)
	}

	// Iterate through the declarations to find the const `docTemplate`
	for _, decl := range node.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok || genDecl.Tok.String() != "const" {
			continue
		}

		for _, spec := range genDecl.Specs {
			valueSpec, ok := spec.(*ast.ValueSpec)
			if !ok || len(valueSpec.Names) == 0 {
				continue
			}

			// Check if the constant is named `docTemplate`
			if valueSpec.Names[0].Name == "docTemplate" {
				// Extract the value (it will include backticks and the raw string literal)
				rawValue := valueSpec.Values[0].(*ast.BasicLit).Value
				rawValue = strings.Trim(rawValue, "`")
				rawValue = RemoveSpecialLines(rawValue)
				os.WriteFile("tmp.json", []byte(rawValue), 0644)
				return rawValue, nil
			}
		}
	}

	fmt.Println("docTemplate constant not found.")
	return "", nil
}

func RemoveSpecialLines(rawValue string) string {
	re := regexp.MustCompile(`(?m)^\s*"schemes":.*\n`)

	updatedValue := re.ReplaceAllString(rawValue, "")

	re = regexp.MustCompile(`(?m)^\s*"basePath":.*\n`)
	updatedValue = re.ReplaceAllString(updatedValue, "")

	re = regexp.MustCompile(`(?m)^\s*"host":.*\n`)
	updatedValue = re.ReplaceAllString(updatedValue, "")

	re = regexp.MustCompile(`(?m)^\s*"info": \{\s*\n\s*"contact": \{\},\s*\n\s*"description": ".*",\s*\n\s*"title": ".*",\s*\n\s*"version": ".*"\s*\n\s*\},\s*\n`)
	updatedValue = re.ReplaceAllString(updatedValue, "")

	return updatedValue
}

func formatHTTPMethodName(updatedJSON []byte) []byte {
	re := regexp.MustCompile(`(?m)^\s*"POST":.*\n`)
	updatedJSON = re.ReplaceAll(updatedJSON, []byte(`"post": {`))

	re = regexp.MustCompile(`(?m)^\s*"GET":.*\n`)
	updatedJSON = re.ReplaceAll(updatedJSON, []byte(`"get": {`))

	re = regexp.MustCompile(`(?m)^\s*"DELETE":.*\n`)
	updatedJSON = re.ReplaceAll(updatedJSON, []byte(`"delete": {`))

	re = regexp.MustCompile(`(?m)^\s*"PUT":.*\n`)
	updatedJSON = re.ReplaceAll(updatedJSON, []byte(`"put": {`))

	return updatedJSON
}

func UpdateDocTemplateWithJSON(filePath, tmpFilePath string, paths map[string]interface{}) error {
	// Read the content of tmp.json
	tmpContent, err := os.ReadFile(tmpFilePath)

	if err != nil {
		return fmt.Errorf("error reading tmp.json: %w", err)
	}
	updatedJSON, err := json.MarshalIndent(paths, "", "  ")
	if err != nil {
		fmt.Printf("Error serializing JSON for file %s: %v\n", tmpFilePath, err)
		return err
	}
	// make all POST / GET / DELETE / PUT in tmp.json to be lowercase
	updatedJSON = formatHTTPMethodName(updatedJSON)

	err = os.WriteFile(tmpFilePath, updatedJSON, 0644)
	if err != nil {
		fmt.Printf("Error writing JSON to file %s: %v\n", tmpFilePath, err)
		return err
	}

	tmpContent, err = os.ReadFile(tmpFilePath)
	if err != nil {
		return fmt.Errorf("error reading tmp.json: %w", err)
	}

	prefixedContent := fmt.Sprintf(`{
  "schemes": {{ marshal .Schemes }},
  "basePath": "{{.BasePath}}",
  "host": "{{.Host}}",
  "info": {
    "contact": {},
    "description": "{{escape .Description}}",
    "title": "{{.Title}}",
    "version": "{{.Version}}"
  },
%s`, tmpContent[1:])

	// fmt.Println(prefixedContent)

	// Parse the Go file
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		return fmt.Errorf("failed to parse Go file: %w", err)
	}

	// Locate the `docTemplate` constant in the AST
	var found bool
	ast.Inspect(node, func(n ast.Node) bool {
		genDecl, ok := n.(*ast.GenDecl)
		if !ok || genDecl.Tok != token.CONST {
			return true // Continue traversing
		}

		for _, spec := range genDecl.Specs {
			valueSpec, ok := spec.(*ast.ValueSpec)
			if !ok || len(valueSpec.Names) == 0 {
				continue
			}

			if valueSpec.Names[0].Name == "docTemplate" {
				// Update its value
				rawString := fmt.Sprintf("`%s`", prefixedContent)
				valueSpec.Values[0] = &ast.BasicLit{
					Kind:  token.STRING,
					Value: rawString,
				}
				found = true
				return false // Stop traversing
			}
		}
		return true
	})

	if !found {
		return fmt.Errorf("docTemplate constant not found in file: %s", filePath)
	}

	// Write the updated AST back to the file
	var buf bytes.Buffer
	if err := printer.Fprint(&buf, fset, node); err != nil {
		return fmt.Errorf("error printing updated Go file: %w", err)
	}

	err = os.WriteFile(filePath, buf.Bytes(), 0644)
	if err != nil {
		return fmt.Errorf("error writing updated Go file: %w", err)
	}

	fmt.Printf("Successfully updated docTemplate in file: %s\n", filePath)
	return nil
}
