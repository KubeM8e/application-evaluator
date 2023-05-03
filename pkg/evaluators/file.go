package evaluators

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	gitHubWorkflowFolder = "workflows"
	dockerizationCommand = "docker/build-push-action@v1"
	serviceKind          = "kind: Service"
	deploymentKind       = "kind: Deployment"
)

func DockerizationEvaluator2(sourceCodeDir string) bool {
	var hasDockerized bool

	// walks through all the directories and subdirectories
	err := filepath.Walk(sourceCodeDir, func(path string, info os.FileInfo, err error) error {
		file, err := os.Open(path)
		if err != nil {
			return err
		}

		defer file.Close()

		// Read the file line by line
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			if strings.Contains(line, dockerizationCommand) {
				hasDockerized = true
			}
		}

		return err
	})

	if err != nil {
		log.Println(err)
	}

	return hasDockerized
}

func DockerizationEvaluator(sourceCodeDir string) bool {

	var workflowsPath string
	var hasDockerized bool

	err := filepath.Walk(sourceCodeDir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			if info.Name() == gitHubWorkflowFolder {
				workflowsPath = path
			} else {
				hasDockerized = false
			}
		}
		return err
	})
	if err != nil {
		log.Println(err)
	}

	err = filepath.Walk(workflowsPath, func(path string, info os.FileInfo, err error) error {
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			if strings.Contains(line, dockerizationCommand) {
				hasDockerized = true
			}
		}
		return err
	})

	if err != nil {
		log.Println(err)
	}

	return hasDockerized
}

func HelmConfigsEvaluator(sourceCodeDir string) (bool, bool) {
	var hasServiceYaml bool
	var hasDeploymentYaml bool

	err := filepath.Walk(sourceCodeDir, func(path string, info os.FileInfo, err error) error {
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()

			// service
			if strings.Contains(line, serviceKind) {
				hasServiceYaml = true
			}

			// deployment
			if strings.Contains(line, deploymentKind) {
				hasDeploymentYaml = true
			}
		}
		return err
	})

	if err != nil {
		log.Println(err)
	}

	return hasServiceYaml, hasDeploymentYaml

}
