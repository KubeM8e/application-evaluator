package evaluators

import (
	"application-evaluator/models"
	"testing"
)

func TestEvaluateTechnologies(t *testing.T) {
	techData := []models.TechData{
		{
			Technology: "Go",
			Keywords:   []string{"go", "golang", "go language"},
		},
		{
			Technology: "Python",
			Keywords:   []string{"python", "python language"},
		},
	}

	dir := "./testdata"
	expectedTech := "Go"
	result := EvaluateTechnologies(techData, dir)

	if result != expectedTech {
		t.Errorf("Expected technology to be %s but got %s", expectedTech, result)
	}
}
