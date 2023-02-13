package evaluators

import (
	"application-evaluator/models"
	"bufio"
	"fmt"
	"go/parser"
	"go/token"
	"golang.org/x/exp/slices"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func EvaluateDatabases1(sourceCodeDir string, DBData []models.DBData, language string) string {
	// According to language used, check for database type and return what database type has been used

	var database string
	var databases []string

	for _, data := range DBData {
		databases = append(databases, data.Database)
	}

	for _, data := range DBData {
		if strings.EqualFold(language, data.Language) {
			err := filepath.Walk(sourceCodeDir, func(path string, info os.FileInfo, err error) error {

				// Open the file
				file, err := os.Open(path)
				if err != nil {
					return err
				}

				if strings.EqualFold(language, "Go") {
					database = checkImportsInGo(file, data)
					for _, v := range databases {
						if database == v {

						}
					}
				} else {
					database = checkImports(file, data)
				}

				//defer file.Close()

				return err
			})

			if err != nil {
				return ""
			}
		}

	}

	return database
}

func EvaluateDatabases(sourceCodeDir string, DBData []models.DBData, language string) string {
	// According to language used, check for database type and return what database type has been used

	var database string
	var databases []string

	for _, data := range DBData {
		databases = append(databases, data.Database)
	}

	filepath.Walk(sourceCodeDir, func(path string, info os.FileInfo, err error) error {
		// Open the file
		file, err := os.Open(path)
		if err != nil {
			return err
		}

		// TODO: optimize
		if !info.IsDir() {
			for _, data := range DBData {
				if strings.EqualFold(language, data.Language) {
					fmt.Println(language, data.Language) // TODO: Optimize
					if strings.EqualFold(language, "Go") {
						database = checkImportsInGo(file, data)
					} else {
						database = checkImports(file, data)
					}

					for _, v := range databases {
						if database == v {
							return filepath.SkipDir
						}
					}
				}
			}
		}

		return err
	})

	return database
}

func checkImportsInGo1(file *os.File, data models.DBData) string {
	var inImports bool
	var database string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "import") {
			inImports = true
		}

		if inImports && strings.HasSuffix(line, ")") {
			inImports = false
		}

		// import ("dvfdgdg)
		// import (
		//"sdfsdjfsd"
		//			"dfsfsfsf" )
		if strings.HasPrefix(line, "import") && strings.HasSuffix(line, ")") {
			for _, keyword := range data.Keywords {
				if strings.Contains(line, keyword) {
					database = data.Database
				}
			}
		}

		if inImports {
			for _, keyword := range data.Keywords {
				if strings.Contains(line, keyword) {
					database = data.Database
				}
			}
		}

	}

	return database
}

// TODO: make this function work - using parseDir()
//func checkImportsInGo2(info os.FileInfo, path string, isDir bool, data models.DBData) string {
//	var database string
//	imports := make([]string, 0)
//
//	if isDir {
//		fileSet := token.NewFileSet()
//		pkgs, err := parser.ParseDir(fileSet, path, nil, parser.ImportsOnly)
//		if err != nil {
//			log.Printf("Could not parse directory: %v", err)
//		}
//
//		for _, pkg := range pkgs {
//			if pkg.Imports == nil {
//				fmt.Printf("No imports found in file %s", pkg.Name)
//				continue
//			}
//
//			for _, importSpec := range pkg.Imports {
//				importPath, _ := strconv.Unquote(importSpec.Path.Value)
//				fmt.Printf("%s imports %s\n", parsedFile.Name, importPath)
//			}
//		}
//	}
//
//	astFile, err := parser.ParseFile(fileSet, file.Name(), nil, parser.ImportsOnly)
//	if err != nil {
//		log.Printf("Could not parse file: %v", err)
//	}
//
//	imports := make([]string, 0)
//	if astFile == nil && len(imports) > 0 {
//		for _, importSpec := range astFile.Imports {
//			imports = append(imports, importSpec.Path.Value[1:len(importSpec.Path.Value)-1])
//		}
//
//		for _, imp := range imports {
//			keywordFound := slices.Contains(data.Keywords, imp)
//			if keywordFound {
//				database = data.Database
//				break
//			}
//		}
//	}
//
//	fmt.Println("database", database)
//	return database
//}

func checkImportsInGo(file *os.File, data models.DBData) string {
	var database string
	imports := make([]string, 0)

	fileSet := token.NewFileSet()
	astFile, err := parser.ParseFile(fileSet, file.Name(), nil, parser.ImportsOnly)
	if err != nil {
		log.Printf("Could not parse file: %v", err)
	}

	if astFile != nil && len(astFile.Imports) > 0 {
		fmt.Println("True")
		for _, importSpec := range astFile.Imports {
			unquotedImport, errUnq := strconv.Unquote(importSpec.Path.Value)
			fmt.Println("UN: ", unquotedImport)
			if errUnq != nil {
				log.Printf("Could not unquote: %v", errUnq)
			}

			imports = append(imports, unquotedImport)
		}
	}

	for _, v := range imports {
		fmt.Println("Import: ", v)
	}

	for _, imp := range imports {
		keywordFound := slices.Contains(data.Keywords, imp)
		if keywordFound {
			database = data.Database
			break
		}
	}

	return database
}

func checkImports(file *os.File, data models.DBData) string {
	var database string

	scanner := bufio.NewScanner(file)

loop:
	for scanner.Scan() {

		line := scanner.Text()

		if strings.HasPrefix(line, "import") {
			for _, keyword := range data.Keywords {
				if strings.Contains(line, keyword) {
					database = data.Database
					break loop
				}
			}
		}
	}

	return database
}
