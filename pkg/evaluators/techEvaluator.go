package evaluators

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func EvaluateTechnologies(techMap map[string][]string, sourceCodeDir string) string {
	var techFrequency []string

	// Walk through all the files in the directory
	err := filepath.Walk(sourceCodeDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Open the file
		file, err := os.Open(path)
		if err != nil {
			return err
		}

		defer file.Close()

		// Read the file line by line
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()

			// Check if any of the keywords are present in the line
			for tech, keywords := range techMap {
				for _, word := range keywords {
					if contains(line, word) {
						techFrequency = append(techFrequency, tech)
					}
				}
			}
		}
		return nil
	})

	if err != nil {
		log.Fatalf("Could not scan the folder: %s", err)
	}

	// check the highest frequency tech
	sourcecodeTech := FindTech(techFrequency)

	return sourcecodeTech

}

// Function to check if a string contains another string
func contains(s, substr string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(substr))
}

func FindTech(techSlice []string) string {
	techFrequency := make(map[string]int)

	for _, str := range techSlice {
		techFrequency[str]++
	}

	mostFrequentStr := ""
	mostFrequency := 0
	for str, frequency := range techFrequency {
		if frequency > mostFrequency {
			mostFrequency = frequency
			mostFrequentStr = str
		}
	}

	return mostFrequentStr

}
