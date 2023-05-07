package ruleEngine

import (
	"application-evaluator/models"
	"testing"
)

func TestEvaluateServiceType(t *testing.T) {
	// create a sample ServiceEvaluation object
	evalObj := models.ServiceEvaluation{
		ServiceType: "statefulService",
	}

	// call the EvaluateServiceType function with the sample object
	result := EvaluateServiceType(evalObj)

	// check that the evaluation field has been updated
	if result.ServiceType != "statefulService" {
		t.Errorf("Expected evaluation to be 'statefulService', got %s", result.ServiceType)
	}

	// check that the ServiceEvaluation object has not been modified in place
	if &evalObj == &result {
		t.Error("Expected EvaluateServiceType to return a new ServiceEvaluation object")
	}
}
