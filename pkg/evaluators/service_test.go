package evaluators

import (
	"application-evaluator/models"
	"os"
	"reflect"
	"testing"
)

func TestEvaluateServiceDependencies(t *testing.T) {
	file, err := os.Open("./testdata/test.go")
	if err != nil {
		t.Errorf("Failed to open file: %v", err)
	}
	defer file.Close()

	serviceData := []models.ServiceData{{Language: "Go", Keywords: []string{"database"}}}

	serviceDependencies := EvaluateServiceDependencies(file, serviceData, "Go", []string{}, []models.ServiceDependency{})

	if len(serviceDependencies) != 1 {
		t.Errorf("Expected 1 service dependency, but got %d", len(serviceDependencies))
	}

	expectedServiceDependency := models.ServiceDependency{}

	if !reflect.DeepEqual(serviceDependencies[0], expectedServiceDependency) {
		t.Errorf("Expected %v, but got %v", expectedServiceDependency, serviceDependencies[0])
	}
}

func TestSortAllMicroservices(t *testing.T) {
	allMicroservices := []string{"user", "order", "payment", "auth"}
	sortedMicroservices := []string{"auth", "user", "order"}

	result := sortAllMicroservices(allMicroservices, sortedMicroservices)

	expected := []string{"payment", "auth", "user", "order"}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestEvaluateMicroservices(t *testing.T) {
	sourceCodeDir := "./testdata"
	serviceData := []models.ServiceData{
		{Language: "Go", Keywords: []string{"database"}},
		{Language: "Go", Keywords: []string{"payment"}},
	}
	language := "Go"

	result, _, _ := EvaluateMicroservices(sourceCodeDir, serviceData, language)

	expected := []string{"auth", "payment", "userService", "orderService"}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
