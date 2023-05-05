package file

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
)

// AddImportStatement 添加import语句
// 例如: AddImportStatement("main.go", "fmt")
func AddImportStatement(fileName, importName, alias string) error {
	fset := token.NewFileSet()

	node, err := parser.ParseFile(fset, fileName, nil, parser.ParseComments)
	if err != nil {
		return err
	}

	importPresent := false
	for _, decl := range node.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok || genDecl.Tok != token.IMPORT {
			continue
		}
		for _, spec := range genDecl.Specs {
			importSpec := spec.(*ast.ImportSpec)
			if importSpec.Path.Value == fmt.Sprintf(`"%s"`, importName) {
				importPresent = true
				break
			}
		}
		if importPresent {
			break
		}
	}

	if !importPresent {
		importDecl := &ast.GenDecl{
			Tok: token.IMPORT,
			Specs: []ast.Spec{
				&ast.ImportSpec{
					Path: &ast.BasicLit{
						Kind:  token.STRING,
						Value: fmt.Sprintf(`"%s"`, importName),
					},
					Name: &ast.Ident{Name: alias},
				},
			},
		}
		node.Decls = append([]ast.Decl{importDecl}, node.Decls...)
	}

	var buf bytes.Buffer
	if err := format.Node(&buf, fset, node); err != nil {
		return err
	}

	if err := formatFile(fileName, buf.Bytes()); err != nil {
		return err
	}

	if err := delExcessiveImport(fileName); err != nil {
		return err
	}
	return nil
}

func formatFile(fileName string, content []byte) error {
	_, err := parser.ParseFile(token.NewFileSet(), fileName, nil, parser.ParseComments)
	if err != nil {
		return err
	}
	fileSet := token.NewFileSet()
	fileSet.AddFile(fileName, fileSet.Base(), len(content))
	_, err = parser.ParseFile(fileSet, fileName, content, parser.ParseComments)
	if err != nil {
		return err
	}
	formatted, err := format.Source(content)
	if err != nil {
		return err
	}

	err = writeToFile(fileName, formatted)
	if err != nil {
		return err
	}
	return nil
}

func writeToFile(fileName string, content []byte) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(content)
	if err != nil {
		return err
	}
	return nil
}

// DelExcessiveImport 删除多余的import
func delExcessiveImport(fileName string) error {
	// 读取源代码
	src, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	// 解析源代码
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", src, parser.ParseComments)
	if err != nil {
		return err
	}

	// 删除重复导入
	var imports []*ast.ImportSpec
	seen := map[string]bool{}
	for _, i := range f.Imports {
		path := i.Path.Value
		if !seen[path] {
			seen[path] = true
			imports = append(imports, i)
		}
	}
	f.Imports = imports

	// 格式化代码
	var buf bytes.Buffer
	err = format.Node(&buf, fset, f)
	if err != nil {
		return err
	}

	return nil
}
