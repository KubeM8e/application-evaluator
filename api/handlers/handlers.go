package handlers

import (
	"application-evaluator/models"
	"application-evaluator/pkg/evaluators"
	"application-evaluator/pkg/utils"
	"encoding/json"
	"fmt"
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

const (
	techDB     = "Technologies"
	databaseDB = "Databases"
)

func SourceFileEvalHandler(c echo.Context) error {
	return nil
}

func TechnologyEvalHandler(c echo.Context) error {
	return nil
}

func UploadSourceCodeHandler(c echo.Context) error {
	sourceCodeHeader, err := c.FormFile("sourceCode")
	if err != nil {
		log.Printf("Could not retrieve the file %v", err)
	}

	fileName := utils.CopyUploadedFile(sourceCodeHeader)
	sourceCode := utils.UnZipFile(fileName)

	// evaluates tech (language)
	techDataSlice := utils.ReadFromTechDB(techDB)
	tech := evaluators.EvaluateTechnologies(techDataSlice, sourceCode)

	// check if dockerized
	hasDockerized := evaluators.DockerizationEvaluator(sourceCode)

	// check helm configs
	hasServiceYaml, hasDeploymentYaml := evaluators.HelmConfigsEvaluator(sourceCode)

	// evaluate database
	databasesMap := utils.ReadFromDatabaseDB(databaseDB)
	fmt.Println(databasesMap) // TODO: remove
	// here goes evaluate function

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

func CreateTechDataHandler(c echo.Context) error {
	var request []models.TechData

	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		log.Println(err)
	}

	// database operations
	utils.CreateTechDataDB(request)

	return c.JSON(http.StatusOK, &request)
}

func CreateDBDataHandler(c echo.Context) error {
	var request []models.DBData

	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		log.Fatal(err)
	}

	// database operations
	utils.CreateDBDataDB(request)

	return c.JSON(http.StatusOK, &request)
}
