package ruleEngine

import (
	"application-evaluator/models"
	"reflect"
	"testing"
)

func TestEvaluateKubeConfigType(t *testing.T) {
	// set up test data
	evalObj := models.ServiceEvaluation{
		KubeConfigType: []string{"statefulSetBased", "type2"},
	}

	// call function to be tested
	result := EvaluateKubeConfigType(evalObj)

	// check if the function has modified the input parameter
	if !reflect.DeepEqual(evalObj, result) {
		t.Errorf("EvaluateKubeConfigType() modifies input parameter, expected: %v, got: %v", evalObj, result)
	}

}
