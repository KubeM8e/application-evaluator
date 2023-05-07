package evaluators

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDockerizationEvaluator(t *testing.T) {
	// Test directory with Dockerfile and GitHub Workflows
	dirWithDockerfile := "testdata/dockerized"
	// Test directory without Dockerfile
	dirWithoutDockerfile := "testdata/not_dockerized"

	tests := []struct {
		name     string
		dir      string
		expected bool
	}{
		{"Directory with Dockerfile", dirWithDockerfile, true},
		{"Directory without Dockerfile", dirWithoutDockerfile, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := DockerizationEvaluator(test.dir)
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestHelmConfigsEvaluator(t *testing.T) {
	// Test directory with service.yaml and deployment.yaml
	dirWithHelmConfigs := "testdata/helm_configs"
	// Test directory without helm configs
	dirWithoutHelmConfigs := "testdata/no_helm_configs"

	tests := []struct {
		name                   string
		dir                    string
		expectedServiceYaml    bool
		expectedDeploymentYaml bool
	}{
		{"Directory with service and deployment yaml", dirWithHelmConfigs, true, true},
		{"Directory without helm configs", dirWithoutHelmConfigs, false, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actualService, actualDeployment := HelmConfigsEvaluator(test.dir)
			assert.Equal(t, test.expectedServiceYaml, actualService)
			assert.Equal(t, test.expectedDeploymentYaml, actualDeployment)
		})
	}
}
