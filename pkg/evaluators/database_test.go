package evaluators

import (
	"application-evaluator/models"
	"errors"
	"io/ioutil"
	"os"
	"testing"
)

func TestEvaluateDB(t *testing.T) {
	testDBData := []models.DBData{
		{
			Language: "Go",
			Keywords: []string{"mysql", "postgres"},
			Database: "RDBMS",
		},
		{
			Language: "Java",
			Keywords: []string{"mongodb", "cassandra"},
			Database: "NoSQL",
		},
	}

	testCases := []struct {
		fileContent string
		expectedDB  string
		expectedErr error
		language    string
	}{
		{
			fileContent: "package main\n\nimport (\n\"database/sql\"\n\"fmt\"\n)\n\nfunc main() {\nfmt.Println(\"Hello, World!\")\n}\n",
			expectedDB:  "",
			expectedErr: nil,
			language:    "Go",
		},
		{
			fileContent: "package main;\n\nimport java.util.Scanner;\n\npublic class Main {\n\n  public static void main(String[] args) {\n\n    Scanner input = new Scanner(System.in);\n    System.out.print(\"Enter your name: \");\n    String name = input.nextLine();\n    System.out.println(\"Hello, \" + name);\n  }\n}\n",
			expectedDB:  "NoSQL",
			expectedErr: nil,
			language:    "Java",
		},
		{
			fileContent: "invalid file content",
			expectedDB:  "",
			expectedErr: errors.New("Could not scan file"),
			language:    "Python",
		},
	}

	for _, tc := range testCases {
		f, err := ioutil.TempFile("", "testfile")
		if err != nil {
			t.Fatalf("Failed to create temporary file: %v", err)
		}
		defer os.Remove(f.Name())

		_, err = f.WriteString(tc.fileContent)
		if err != nil {
			t.Fatalf("Failed to write file content: %v", err)
		}

		gotDB := EvaluateDB(f, testDBData, tc.language)

		if gotDB != tc.expectedDB {
			t.Errorf("Expected database %s, but got %s", tc.expectedDB, gotDB)
		}

	}
}
