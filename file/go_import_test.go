package file

import "testing"

func TestAddImportStatement(t *testing.T) {
	fileName := "go_import.go"
	importName := "fmt1"
	// alias := "f"
	alias := ""
	AddImportStatement(fileName, importName, alias)
}
