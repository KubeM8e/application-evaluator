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

	return c.JSON(http.StatusOK, "File uploaded successfully. The technology used is "+tech)
}
