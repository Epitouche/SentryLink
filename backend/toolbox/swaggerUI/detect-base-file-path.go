package swaggerui

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
	"strings"
)

type BasePath interface {
	DetectBasePathFromProject(entryFile string) (*BasePathInfo, error)
}

type BasePathInfo struct {
	BasePath string
}

func detectBasePathFromFile(filePath string) (*BasePathInfo, error) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filePath, nil, parser.AllErrors)
	if err != nil {
		return nil, fmt.Errorf("error parsing file %s: %w", filePath, err)
	}

	var basePath string
	ast.Inspect(node, func(n ast.Node) bool {
		assignStmt, ok := n.(*ast.AssignStmt)
		if !ok {
			return true
		}

		for _, lhs := range assignStmt.Lhs {
			selectorExpr, ok := lhs.(*ast.SelectorExpr)
			if !ok {
				continue
			}

			// Detect `docs.SwaggerInfo.BasePath`
			xIdent, ok := selectorExpr.X.(*ast.SelectorExpr)
			if ok && xIdent.Sel.Name == "SwaggerInfo" && selectorExpr.Sel.Name == "BasePath" {
				if len(assignStmt.Rhs) > 0 {
					if basicLit, ok := assignStmt.Rhs[0].(*ast.BasicLit); ok && basicLit.Kind == token.STRING {
						basePath = strings.Trim(basicLit.Value, `"`)
					}
				}
			}
		}
		return true
	})

	if basePath != "" {
		return &BasePathInfo{
			BasePath: basePath,
		}, nil
	}
	return nil, nil
}

func DetectBasePathFromProject(entryFile string) (*BasePathInfo, error) {
	projectDir := filepath.Dir(entryFile)
	files, err := findAllGoFiles(projectDir)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		basePathInfo, err := detectBasePathFromFile(file)
		if err != nil {
			return nil, err
		}
		if basePathInfo != nil {
			return basePathInfo, nil
		}
	}
	return nil, fmt.Errorf("BasePath not found")
}
