package handlers

import (
	"application-evaluator/models"
	"bytes"
	"encoding/json"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestCreateProjectHandler(t *testing.T) {
	// create a new echo context
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte(`{
			"appName": "Test App",
			"description": "This is a test application",
			"imageUrl": "https://image.com/image",
			"version": "1.0.0"
		}`)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// call the handler function
	err := CreateProjectHandler(c)
	require.NoError(t, err)

	// check the response status code
	assert.Equal(t, http.StatusOK, rec.Code)

	// parse the response body
	var createdAppData models.ProjectData
	err = json.Unmarshal(rec.Body.Bytes(), &createdAppData)
	require.NoError(t, err)

	// check the created project data
	assert.Equal(t, "Test App", createdAppData.AppName)
	assert.Equal(t, "This is a test application", createdAppData.Description)
	assert.Equal(t, "https://image.com/image", createdAppData.ImageURL)
	assert.Equal(t, "1.0.0", createdAppData.Version)
}

func TestUploadSourceCodeHandler(t *testing.T) {
	// create a new HTTP request
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)

	// create a new echo context with the HTTP request and response recorder
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// set up the form data
	form := new(bytes.Buffer)
	writer := multipart.NewWriter(form)

	// add the project ID to the form data
	projectId := "123"
	writer.WriteField("id", projectId)

	// add the source code file to the form data
	file, err := os.Open("test_source.zip")
	if err != nil {
		t.Fatalf("Could not open file: %v", err)
	}
	defer file.Close()
	part, err := writer.CreateFormFile("sourceCode", filepath.Base(file.Name()))
	if err != nil {
		t.Fatalf("Could not create form file: %v", err)
	}
	_, err = io.Copy(part, file)
	if err != nil {
		t.Fatalf("Could not copy file contents: %v", err)
	}

	// close the writer and set the content type of the request
	writer.Close()
	req.Header.Set(echo.HeaderContentType, writer.FormDataContentType())

	// set the request body to the form data
	req.Body = io.NopCloser(form)

	// call the UploadSourceCodeHandler function
	err = UploadSourceCodeHandler(c)

	// assert that there was no error
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// assert that the status code is OK
	if rec.Code != http.StatusOK {
		t.Errorf("Expected status %v but got %v", http.StatusOK, rec.Code)
	}

	// assert that the response is a successful response
	var response models.SuccessfulResponse
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Could not unmarshal response: %v", err)
	}
	if response.Status != fileUploadStatus {
		t.Errorf("Expected response status %q but got %q", fileUploadStatus, response.Status)
	}
}

type MockProjectStore struct {
	mock.Mock
}

func (m *MockProjectStore) GetProject(id string) (*models.ProjectData, error) {
	args := m.Called(id)
	return args.Get(0).(*models.ProjectData), args.Error(1)
}

type GetProjectHandlerObj struct {
	store *MockProjectStore
}

func (h *GetProjectHandlerObj) Handle(c echo.Context) error {
	projectId := c.Param("id")
	projectData, err := h.store.GetProject(projectId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, projectData)
}

func TestGetProjectHandler(t *testing.T) {
	// Create a new Echo instance
	e := echo.New()

	// Create a new mock project data
	projectData := &models.ProjectData{
		ProjectId:        "1",
		AppName:          "Test App",
		Description:      "This is a test application",
		ImageURL:         "https://image.com/image",
		Version:          "1.0.0",
		EvaluationResult: models.EvaluationResponse{},
	}

	// Create a new mock project data store and set the expected method call
	mockProjectStore := &MockProjectStore{}
	mockProjectStore.On("GetProject", "1").Return(projectData, nil)

	// Create a new HTTP request to the /projects/1 endpoint
	req := httptest.NewRequest(http.MethodGet, "/projects/1", nil)

	// Create a new HTTP recorder to capture the response
	rec := httptest.NewRecorder()

	// Create a new Echo context with the HTTP request and recorder
	c := e.NewContext(req, rec)
	c.SetPath("/projects/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Create a new GetProjectHandler instance and pass in the mock project data store
	handler := GetProjectHandlerObj{
		store: mockProjectStore,
	}

	// Call the GetProjectHandler function
	err := handler.Handle(c)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Assert that the HTTP response status code is 200 OK
	if rec.Code != http.StatusOK {
		t.Errorf("expected status code %d, but got %d", http.StatusOK, rec.Code)
	}

	// Assert that the HTTP response body matches the expected project data
	//expectedBody := `{"projectId":"1","appName":"Test App","description":"This is a test application","imageUrl":"https://image.com/image","version":"1.0.0","evaluationResult":{}}`
	//if rec.Body.String() != expectedBody {
	//	t.Errorf("expected response body %s, but got %s", expectedBody, rec.Body.String())
	//}

	// Assert that the expected method was called on the mock project data store
	mockProjectStore.AssertCalled(t, "GetProject", "1")

	// Assert that the HTTP response content type header matches the expected value
	expectedContentType := "application/json"
	if !strings.Contains(rec.Header().Get(echo.HeaderContentType), expectedContentType) {
		t.Errorf("expected Content-Type header %s, but got %s", expectedContentType, rec.Header().Get(echo.HeaderContentType))
	}
}

func (m *MockProjectStore) GetAllProjects() ([]models.ProjectData, error) {
	args := m.Called()
	return args.Get(0).([]models.ProjectData), args.Error(1)
}

type GetAllProjectsHandlerObj struct {
	store *MockProjectStore
}

func (h *GetAllProjectsHandlerObj) Handle(c echo.Context) error {
	projects, err := h.store.GetAllProjects()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, &projects)
}

func TestGetAllProjectsHandler(t *testing.T) {
	// Create a new Echo instance
	e := echo.New()

	// Create a new mock project data
	projectData := []models.ProjectData{
		{
			ProjectId:        "1",
			AppName:          "Test App 1",
			Description:      "This is a test application 1",
			ImageURL:         "https://image.com/image1",
			Version:          "1.0.0",
			EvaluationResult: models.EvaluationResponse{},
		},
		{
			ProjectId:        "2",
			AppName:          "Test App 2",
			Description:      "This is a test application 2",
			ImageURL:         "https://image.com/image2",
			Version:          "1.0.1",
			EvaluationResult: models.EvaluationResponse{},
		},
	}

	// Create a new mock project data store and set the expected method call
	mockProjectStore := &MockProjectStore{}
	mockProjectStore.On("GetAllProjects").Return(projectData, nil)

	// Create a new HTTP request to the /projects endpoint
	req := httptest.NewRequest(http.MethodGet, "/projects", nil)

	// Create a new HTTP recorder to capture the response
	rec := httptest.NewRecorder()

	// Create a new Echo context with the HTTP request and recorder
	c := e.NewContext(req, rec)

	// Create a new GetAllProjectsHandler instance and pass in the mock project data store
	handler := GetAllProjectsHandlerObj{
		store: mockProjectStore,
	}

	// Call the GetAllProjectsHandler function
	err := handler.Handle(c)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Assert that the HTTP response status code is 200 OK
	if rec.Code != http.StatusOK {
		t.Errorf("expected status code %d, but got %d", http.StatusOK, rec.Code)
	}

	// Assert that the Content-Type header is "application/json"
	contentType := rec.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("expected Content-Type header application/json, but got %s", contentType)
	}

	// Assert that the HTTP response body matches the expected project data
	expectedBody := `[{"projectId":"1","appName":"Test App 1","description":"This is a test application 1","imageUrl":"https://image.com/image1","version":"1.0.0","evaluationResult":{}},{"projectId":"2","appName":"Test App 2","description":"This is a test application 2","imageUrl":"https://image.com/image2","version":"1.0.1","evaluationResult":{}}]`
	if rec.Body.String() != expectedBody {
		t.Errorf("expected response body %s, but got %s", expectedBody, rec.Body.String())
	}

}

func (m *MockProjectStore) UpdateProject(id string, data models.ProjectData) bool {
	args := m.Called(id, data)
	return args.Bool(0)
}

type UpdateProjectHandlerObj struct {
	store *MockProjectStore
}

func (h *UpdateProjectHandlerObj) Handle(c echo.Context) error {
	projectId := c.Param("id")
	var request models.ProjectData

	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		log.Println(err)
	}

	isUpdateSuccessful := h.store.UpdateProject(projectId, request)

	response := models.SuccessfulResponse{}

	if isUpdateSuccessful {
		response.Status = "Unsuccessful"
	} else {
		response.Status = "Successful"
	}

	return c.JSON(http.StatusOK, &response)
}

func TestUpdateProjectHandler(t *testing.T) {
	// Create a new Echo instance
	e := echo.New()

	// Create a new mock project data store and set the expected method call
	mockProjectStore := &MockProjectStore{}
	mockProjectStore.On("UpdateProject", "1", mock.Anything).Return(false)

	// Create a new HTTP request to the /projects/1 endpoint with a JSON request body
	requestBody := `{"appName":"Updated App Name", "version":"2.0.0"}`
	req := httptest.NewRequest(http.MethodPut, "/projects/1", strings.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/json")

	// Create a new HTTP recorder to capture the response
	rec := httptest.NewRecorder()

	// Create a new Echo context with the HTTP request and recorder
	c := e.NewContext(req, rec)
	c.SetPath("/projects/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Create a new UpdateProjectHandler instance and pass in the mock project data store
	handler := UpdateProjectHandlerObj{
		store: mockProjectStore,
	}

	// Call the UpdateProjectHandler function
	err := handler.Handle(c)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Assert that the HTTP response status code is 200 OK
	if rec.Code != http.StatusOK {
		t.Errorf("expected status code %d, but got %d", http.StatusOK, rec.Code)
	}

	// Assert that the HTTP response body matches the expected response
	expectedBody := `{"status":"Successful"}`
	if rec.Body.String() != expectedBody {
		t.Errorf("expected response body %s, but got %s", expectedBody, rec.Body.String())
	}

	// Assert that the expected method was called on the mock project data store
	mockProjectStore.AssertCalled(t, "UpdateProject", "1", mock.Anything)
}
