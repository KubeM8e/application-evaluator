package evaluators

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

const (
	gitHubFolder         = ".github"
	gitHubWorkflowFolder = "workflows"
	dockerizationCommand = "docker/build-push-action@v1"
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
			if strings.Contains(dockerizationCommand, line) {
				hasDockerized = true
			}
		}

		return err
	})

	if err != nil {
		fmt.Println(err)
	}

	return hasDockerized
}

func DockerizationEvaluator(sourceCodeDir string) bool {

	var workflowsPath string
	var hasDockerized bool

	err := filepath.Walk(sourceCodeDir, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			if info.Name() == gitHubWorkflowFolder {
				workflowsPath = path
			} else {
				hasDockerized = false
			}
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
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
			if strings.Contains(dockerizationCommand, line) {
				hasDockerized = true
			}
		}
		return nil
	})

	if err != nil {
		fmt.Println(err)
	}

	return hasDockerized
}
