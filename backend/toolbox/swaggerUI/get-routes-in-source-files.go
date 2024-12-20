package swaggerui

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

type DetectionProcess interface {
	ExtractRouteFromProject(entryFile string, basePath *BasePathInfo) ([]RouteFound, error)
}

type RouteFound struct {
	Method string
	Path   string
}

func extractRoutesFromFile(filePath string, basePath *BasePathInfo) ([]RouteFound, error) {
	var routes []RouteFound

	// Parse the file into an AST
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filePath, nil, parser.AllErrors)
	if err != nil {
		return nil, fmt.Errorf("error parsing file %s: %w", filePath, err)
	}

	// Track group prefixes
	var groupPrefixes []string

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
				// Extract the group prefix
				if len(callExpr.Args) > 0 {
					if arg, ok := callExpr.Args[0].(*ast.BasicLit); ok && arg.Kind == token.STRING {
						prefix := strings.Trim(arg.Value, `"`)
						fmt.Printf("Prefix: %s\n", prefix)
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

					// Combine group prefixes with the route path
					fullPath := strings.Join(append(groupPrefixes, path), "")

					// Add the route to the list
					routes = append(routes, RouteFound{
						Method: method,
						Path:   fullPath,
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

func ExtractRouteFromProject(entryFile string, basePath *BasePathInfo) ([]RouteFound, error) {
	var allFoundRoutes []RouteFound
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
