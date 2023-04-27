package evaluators

import (
	"application-evaluator/models"
	"application-evaluator/pkg/utils"
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"golang.org/x/exp/slices"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	OpenAIAPIKey  = "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"
	OpenAIPrompt  = "If there are calls between microservices answer yes or no and list down the microservice in point form which calls the other microservice as \"microservice\" and the second microservice as \"depends on\" and the URL that is being called as \"URL\"."
	OpenAIPrompt1 = "Is there a call to another microservice in the following source code? If so which microservice is calling which? Provide the URL too."
)

func EvaluateServiceDependencies(sourceCodeDir string, serviceData []models.ServiceData, language string) {
	var imports []string
	var serviceDependencies []models.ServiceDependency

	filepath.Walk(sourceCodeDir, func(path string, info os.FileInfo, err error) error {
		// Open the file
		file, err := os.Open(path)
		if err != nil {
			return err
		}

		// TODO: optimize
		if !info.IsDir() {
			for _, data := range serviceData {
				if strings.EqualFold(language, data.Language) {
					if strings.EqualFold(language, "Go") {
						imports = utils.CheckImportsInGo(file) // scan import section in go files

						for _, imp := range imports {
							keywordFound := slices.Contains(data.Keywords, imp)
							if keywordFound {
								evaluationResp := passFileToDavinci(file.Name())       // reads the content of the file and pass to text-davinci-003
								serviceDep := extractServiceDependency(evaluationResp) // extracts service dependency data
								serviceDependencies = append(serviceDependencies, serviceDep)
								//break
							}
						}

					} else {
						//imports = checkImports(file)

						scanner := bufio.NewScanner(file)

					loopScan:
						for scanner.Scan() {
							line := scanner.Text()
							for _, keyword := range data.Keywords {
								if strings.Contains(line, keyword) {
									evaluationResp := passFileToDavinci(file.Name())
									serviceDep := extractServiceDependency(evaluationResp)
									serviceDependencies = append(serviceDependencies, serviceDep)
									break loopScan
								}
							}
						}

						//loop:
						//	for _, imp := range imports {
						//		for _, keyword := range data.Keywords {
						//			if strings.Contains(imp, keyword) {
						//				isServiceDependency = true
						//				break loop
						//			}
						//		}
						//	}

					}
				}
			}
		}

		// sort service dependencies
		sortDependencies(serviceDependencies)
		return err
	})

	for i, v := range serviceDependencies {
		fmt.Println("Dep:", i, v)
	}

}

func callDavinciModel(prompt string) string {
	client := resty.New()

	// headers
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = "Bearer " + OpenAIAPIKey

	//body
	data := models.DavinciRequest{
		Model:       "text-davinci-003",
		Prompt:      prompt,
		MaxTokens:   2408,
		Temperature: 0,
	}

	resp, err := client.R().
		SetHeaders(headers).
		SetBody(data).
		Post("https://api.openai.com/v1/completions")

	if err != nil {
		log.Printf("Could not call the OpenAI API %s", err)
	}

	response := models.DavinciResponse{}
	err = json.Unmarshal(resp.Body(), &response)

	return response.Choices[0].Text
}

func passFileToDavinci(fileName string) string {
	fileContent, err := os.ReadFile(fileName)
	if err != nil {
		log.Printf("Could not read the contetnt of the file: %s", err)
	}
	strFileContent := string(fileContent)

	// call the text-davinci-003
	prompt := OpenAIPrompt + strFileContent
	return callDavinciModel(prompt)

}

func extractServiceDependency(text string) models.ServiceDependency {
	var serviceDep models.ServiceDependency
	// Check if the answer is "Yes"
	if strings.Contains(text, "Yes") {

		// Extract the microservice name, dependency, and URL
		re := regexp.MustCompile(`Microservice:\s+(?P<name>[\w\s]+)\s+Depends on:\s+(?P<depends_on>[\w\s]+)\s+URL:\s+(?P<url>http://\S+)`)
		matches := re.FindStringSubmatch(text)

		fmt.Println(len(matches))
		if len(matches) >= 4 {
			serviceDep = models.ServiceDependency{
				Microservice: strings.TrimSpace(matches[re.SubexpIndex("name")]),
				DependsOn:    strings.TrimSpace(matches[re.SubexpIndex("depends_on")]),
				URL:          strings.TrimSpace(matches[re.SubexpIndex("url")]),
			}

		}
	}

	return serviceDep
}

func sortDependencies(dependencies []models.ServiceDependency) {
	// Create a map to store the incoming edges for each microservice
	incomingEdges := make(map[string][]string)

	// Create a map to store the outgoing edges for each microservice
	outgoingEdges := make(map[string][]string)

	// Populate the incoming and outgoing edges maps based on the microservices list
	for _, dep := range dependencies {
		incomingEdges[dep.Microservice] = []string{}
		outgoingEdges[dep.Microservice] = []string{}
	}

	for _, dep := range dependencies {
		incomingEdges[dep.DependsOn] = append(incomingEdges[dep.DependsOn], dep.Microservice)
		outgoingEdges[dep.Microservice] = append(outgoingEdges[dep.Microservice], dep.DependsOn)
	}

	// Create a list to store the microservices in topologically sorted order
	sortedMicroservices := []string{}

	// Create a queue to store microservices with no incoming edges
	queue := []string{}

	// Populate the queue with microservices with no incoming edges
	for name := range incomingEdges {
		if len(incomingEdges[name]) == 0 {
			queue = append(queue, name)
		}
	}

	// Traverse the microservices in the queue, removing incoming edges and adding new microservices with no incoming edges
	for len(queue) > 0 {
		// Get the next microservice from the queue
		microservice := queue[0]
		queue = queue[1:]

		// Add the microservice to the sorted list
		sortedMicroservices = append(sortedMicroservices, microservice)

		// Remove the outgoing edges for the microservice
		for _, dependent := range outgoingEdges[microservice] {
			incomingEdges[dependent] = remove(incomingEdges[dependent], microservice)
			if len(incomingEdges[dependent]) == 0 {
				queue = append(queue, dependent)
			}
		}
	}

	// Print the sorted list of microservices in the correct order
	for i := len(sortedMicroservices) - 1; i >= 0; i-- {
		fmt.Println(sortedMicroservices[i])
	}
}

func remove(slice []string, element string) []string {
	for i, v := range slice {
		if v == element {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}
