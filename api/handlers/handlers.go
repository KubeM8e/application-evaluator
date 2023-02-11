package handlers

import (
	"application-evaluator/models"
	"application-evaluator/pkg/evaluators"
	"application-evaluator/pkg/utils"
	"encoding/json"
	"github.com/labstack/echo"
	"log"
	"net/http"
)

const (
	fileUploadStatus     = "File uploaded successfully"
	techUsed             = "The technology used is "
	hasDockerizedMsg     = "The application has been dockerized"
	hasNotDockerizedMsg  = "The application has not been dockerized"
	hasServiceYamlMsg    = "The application contains kubernetes manifest of kind service"
	noServiceYamlMsg     = "The application does not contain kubernetes manifest of kind service"
	hasDeploymentYamlMsg = "The application contains kubernetes manifest of kind deployment"
	noDeploymentYamlMsg  = "The application does not contain kubernetes manifest of kind deployment"
)

func CreateTechDataHandler(c echo.Context) error {
	var request []models.CreateTechDataRequest

	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		log.Fatal(err)
	}

	// database operations
	utils.ConnectMongoDB(request)

	return c.JSON(http.StatusOK, &request)
}

func SourceFileEvalHandler(c echo.Context) error {
	return nil
}

func TechnologyEvalHandler(c echo.Context) error {
	return nil
}

func UploadSourceCodeHandler(c echo.Context) error {
	sourceCodeHeader, err := c.FormFile("sourceCode")
	if err != nil {
		log.Fatalf("Could not retrieve the file %v", err)
	}

	fileName := utils.CopyUploadedFile(sourceCodeHeader)
	sourceCode := utils.UnZipFile(fileName)

	// evaluates tech
	technologiesMap := utils.ReadFromDB()
	tech := evaluators.EvaluateTechnologies(technologiesMap, sourceCode)

	// check if dockerized
	hasDockerized := evaluators.DockerizationEvaluator(sourceCode)

	// check helm configs
	hasServiceYaml, hasDeploymentYaml := evaluators.HelmConfigsEvaluator(sourceCode)

	// response
	response := models.FileUploadResponse{}
	response.UploadStatus = fileUploadStatus
	response.TechnologyUsed = techUsed + tech

	if hasDockerized {
		response.DockerizationStatus = hasDockerizedMsg
	} else {
		response.DockerizationStatus = hasNotDockerizedMsg
	}

	if hasServiceYaml {
		response.ServiceYamlStatus = hasServiceYamlMsg
	} else {
		response.ServiceYamlStatus = noServiceYamlMsg
	}

	if hasDeploymentYaml {
		response.DeploymentYamlStatus = hasDeploymentYamlMsg
	} else {
		response.DeploymentYamlStatus = noDeploymentYamlMsg
	}

	return c.JSON(http.StatusOK, &response)
}
