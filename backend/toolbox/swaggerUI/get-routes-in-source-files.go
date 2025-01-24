package swaggerui

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"path/filepath"
	"strings"

	"github.com/Tom-Mendy/SentryLink/toolbox/swaggerUI/schemas"
)

type DetectionProcess interface {
	ExtractRouteFromProject(entryFile string, basePath *BasePathInfo) ([]schemas.RouteFound, error)
}

func extractRoutesFromFile(filePath string, basePath *BasePathInfo) ([]schemas.RouteFound, error) {
	var routes []schemas.RouteFound

	// Parse the file into an AST
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filePath, nil, parser.AllErrors)
	if err != nil {
		return nil, fmt.Errorf("error parsing file %s: %w", filePath, err)
	}

	// Track group prefixes
	groupPrefixes := []string{}
	realPrefix := []string{}

	// Walk through the AST to find Gin route definitions
	ast.Inspect(node, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.AssignStmt:
			// Handle `Group()` calls to track prefixes
			for _, rhs := range x.Rhs {
				callExpr, ok := rhs.(*ast.CallExpr)
				if !ok {
					continue
				}
				selector, ok := callExpr.Fun.(*ast.SelectorExpr)
				if !ok || selector.Sel.Name != "Group" {
					continue
				}
				groupPrefixes = []string{}
				realPrefix = []string{}
				// Extract the group prefix
				if len(callExpr.Args) > 0 {
					if arg, ok := callExpr.Args[0].(*ast.BasicLit); ok && arg.Kind == token.STRING {
						prefix := strings.Trim(arg.Value, `"`)
						fmt.Printf("Prefix: %s\n", prefix)
						realPrefix = append(realPrefix, prefix)
						// routes[len(routes)-1].Prefix = prefix
						if basePath.BasePath != "" && !strings.Contains(prefix, basePath.BasePath) {
							prefix = basePath.BasePath + prefix
						}
						groupPrefixes = append(groupPrefixes, prefix)
					}
				}
			}

		case *ast.CallExpr:
			// Extract routes like `GET`, `POST`, etc.
			selector, ok := x.Fun.(*ast.SelectorExpr)
			if !ok {
				return true
			}
			method := selector.Sel.Name
			if method != "GET" && method != "POST" && method != "PUT" && method != "DELETE" {
				return true
			}

			// Extract the route path
			if len(x.Args) > 0 {
				arg, ok := x.Args[0].(*ast.BasicLit)
				if ok && arg.Kind == token.STRING {
					path := strings.Trim(arg.Value, `"`)

					// Skip ignored routes
					if path == "/" || strings.HasPrefix(path, "/swagger") {
						return true
					}

					handlerArg := x.Args[len(x.Args)-1]

					// Convert the handler AST node to a string
					var handlerNameBuf bytes.Buffer
					if err := printer.Fprint(&handlerNameBuf, fset, handlerArg); err != nil {
						fmt.Printf("Error printing handler: %v\n", err)
						return true
					}
					handlerName := handlerNameBuf.String()

					// Combine group prefixes with the route path
					fullPath := strings.Join(append(groupPrefixes, path), "")
					if len(realPrefix) < 1 {
						break
					}

					// Add the route to the list
					routes = append(routes, schemas.RouteFound{
						Method:      method,
						Path:        fullPath,
						Prefix:      realPrefix[len(realPrefix)-1][1:],
						HandlerName: handlerName,
					})
				}
			}
		}
		return true
	})

	return routes, nil
}

func findAllGoFiles(dir string) ([]string, error) {
	var files []string

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
		}
		if !info.IsDir() && IsGOFile(path) {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

func ExtractRouteFromProject(entryFile string, basePath *BasePathInfo) ([]schemas.RouteFound, error) {
	var allFoundRoutes []schemas.RouteFound
	projectDir := filepath.Dir(entryFile)

	files, err := findAllGoFiles(projectDir)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		routes, err := extractRoutesFromFile(file, basePath)
		if err != nil {
			return nil, err
		}

		allFoundRoutes = append(allFoundRoutes, routes...)
	}
	return allFoundRoutes, nil
}

func ExtractSchemaFromProject(entryDirectory string) []schemas.SchemaValueSchemas {
	schemasValue := []schemas.SchemaValueSchemas{}
	err := filepath.Walk(entryDirectory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if the file is a Go file
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") {
			fmt.Printf("Processing file: %s\n", path)
			getSchemas(path, &schemasValue)
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error walking the directory: %v\n", err)
	}
	// fmt.Printf("Schemas: %++v\n", schemasValue)
	return schemasValue
}

func getSchemas(filePath string, schemasValue *[]schemas.SchemaValueSchemas) {
	// Create a new FileSet to parse the Go file
	fset := token.NewFileSet()

	// Parse the Go file
	node, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		fmt.Printf("Error parsing file %s: %v\n", filePath, err)
		return
	}

	// Traverse the AST to find struct definitions
	ast.Inspect(node, func(n ast.Node) bool {
		// Check if the node is a general declaration (e.g., type, var, const)
		genDecl, ok := n.(*ast.GenDecl)
		if !ok {
			return true
		}

		// Check if the declaration is a type declaration
		if genDecl.Tok != token.TYPE {
			return true
		}

		// Check if the declaration has an @Ignore comment
		if hasIgnoreComment(genDecl.Doc) {
			fmt.Printf("Ignoring group in the file : %s\n", filePath)
			return true
		}

		// Iterate over the specs in the declaration
		for _, spec := range genDecl.Specs {
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}

			// Check if the type is a struct
			structType, ok := typeSpec.Type.(*ast.StructType)
			if !ok {
				continue
			}

			// Print the struct name
			// if hasIgnoreComment(typeSpec.Doc) {
			// 	fmt.Printf("Ignoring struct: %s\n", typeSpec.Name.Name)
			// 	continue
			// }
			fmt.Printf("Found struct: %s\n", typeSpec.Name.Name)

			*schemasValue = append(*schemasValue, schemas.SchemaValueSchemas{
				SchemaName: typeSpec.Name.Name,
			})

			// Print the fields of the struct
			for _, field := range structType.Fields.List {
				fieldName := ""
				if len(field.Names) > 0 {
					fieldName = field.Names[0].Name
				}

				fieldType := fmt.Sprintf("%s", field.Type)
				// fmt.Printf("  Field: %s, Type: %s\n", fieldName, fieldType)
				(*schemasValue)[len(*schemasValue)-1].FieldValues = append((*schemasValue)[len(*schemasValue)-1].FieldValues, schemas.SchemaFieldValues{
					FieldName: fieldName,
					FieldType: fieldType,
				})
			}
		}
		return true
	})
}

// hasIgnoreComment checks if the comment group contains the @Ignore tag
func hasIgnoreComment(comments *ast.CommentGroup) bool {
	if comments == nil {
		return false
	}
	for _, comment := range comments.List {
		if strings.Contains(comment.Text, "@Ignore") {
			return true
		}
	}
	return false
}

func isValidHTTPMethod(method string) bool {
	switch method {
	case "GET", "POST", "PUT", "DELETE", "PATCH", "HEAD", "OPTIONS":
		return true
	default:
		return false
	}
}
