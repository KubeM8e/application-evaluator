package project

import (
	"application-evaluator/models"
	"testing"
)

func TestCreateProject(t *testing.T) {
	// Set up test data
	testData := models.ProjectData{
		AppName:     "Test Project",
		Description: "This is a test project.",
	}

	// Call the function being tested
	result := CreateProject(testData)

	// Verify the result
	if result.ProjectId == "" {
		t.Errorf("Project ID was not generated.")
	}
}

func TestGetAllProjects(t *testing.T) {
	// Call the function being tested
	result := GetAllProjects()

	// Verify the result
	if len(result) == 0 {
		t.Errorf("No project data was returned.")
	}
}

func TestGetProject(t *testing.T) {
	// Set up test data
	testData := models.ProjectData{
		AppName:     "Test Project",
		Description: "This is a test project.",
	}
	CreateProject(testData)
	projectId := testData.ProjectId

	// Call the function being tested
	result := GetProject(projectId)

	// Verify the result
	if result.ProjectId != projectId {
		t.Errorf("Returned project ID does not match input project ID.")
	}
}

func TestUpdateProject(t *testing.T) {
	// Set up test data
	testData := models.ProjectData{
		AppName:     "Test Project",
		Description: "This is a test project.",
	}
	CreateProject(testData)
	projectId := testData.ProjectId
	testData.Description = "This is an updated test project."

	// Call the function being tested
	isError := UpdateProject(projectId, testData)

	// Verify the result
	if isError {
		t.Errorf("Update failed.")
	}
}
