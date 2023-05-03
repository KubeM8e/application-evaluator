package utils

import (
	"bufio"
	"go/parser"
	"go/token"
	"log"
	"os"
	"strconv"
	"strings"
)

// can be used to Java and other languages except Go

func CheckImportsInJava(file *os.File) []string {
	var imports []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		line := scanner.Text()

		if strings.HasPrefix(line, "import") {
			imports = append(imports, line)
		}
	}

	return imports
}

func CheckImportsInGo(file *os.File) []string {
	//var database string
	imports := make([]string, 0)

	fileSet := token.NewFileSet()
	astFile, err := parser.ParseFile(fileSet, file.Name(), nil, parser.ImportsOnly)
	if err != nil {
		log.Printf("Could not parse file: %v", err)
	}

	if astFile != nil && len(astFile.Imports) > 0 {
		for _, importSpec := range astFile.Imports {
			unquotedImport, errUnq := strconv.Unquote(importSpec.Path.Value)
			if errUnq != nil {
				log.Printf("Could not unquote: %v", errUnq)
			}

			imports = append(imports, unquotedImport)
		}
	}

	return imports
}

func ScanSourceCodeFile(file *os.File) []string {
	var imports []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		line := scanner.Text()

		if strings.HasPrefix(line, "import") {
			imports = append(imports, line)
		}
	}

	return imports

}
