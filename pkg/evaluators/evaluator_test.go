package evaluators

import (
	"application-evaluator/models"
	"application-evaluator/pkg/project"
	"testing"
)

func TestEvaluateSourcecode(t *testing.T) {
	// create a dummy project source
	projectSource := models.ProjectSource{
		ProjectId:  "project-123",
		Sourcecode: "package main\n\nfunc main() {}",
	}

	// call the function being tested
	EvaluateSourcecode(projectSource)

	// verify that the evaluation results were saved to the project
	projectData := project.GetProject("project-123")
	if projectData.EvaluationResult.Language != "Go" {
		t.Errorf("Expected language to be 'Go', but got %s", projectData.EvaluationResult.Language)
	}

}

func TestSaveEvaluationResults(t *testing.T) {
	// create a dummy evaluation result and microservices map
	evaluationResult := models.EvaluationResponse{
		Language: "Go",
	}
	microservicesMap := map[string]models.Microservice{
		"service-123": {
			ServiceEvaluation: models.ServiceEvaluation{
				ServiceType: "web",
			},
		},
	}

	// create a dummy project
	projectData := models.ProjectData{
		ProjectId: "project-123",
	}

	// call the function being tested
	SaveEvaluationResults(projectData.ProjectId, evaluationResult, microservicesMap)

	// verify that the evaluation results and microservices map were saved to the project
	projectData = project.GetProject("project-123")
	if projectData.EvaluationResult.Language != "Go" {
		t.Errorf("Expected language to be 'Go', but got %s", projectData.EvaluationResult.Language)
	}
	if len(projectData.Microservices) != 1 {
		t.Errorf("Expected 1 microservice, but got %d", len(projectData.Microservices))
	}
	if projectData.Microservices["service-123"].ServiceEvaluation.ServiceType != "web" {
		t.Errorf("Expected service type to be 'web', but got %s", projectData.Microservices["service-123"].ServiceEvaluation.ServiceType)
	}
}
