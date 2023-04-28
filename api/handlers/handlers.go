package handlers

import (
	"application-evaluator/models"
	"application-evaluator/pkg/utils"
	"encoding/json"
	"github.com/labstack/echo"
	"github.com/streadway/amqp"
	"log"
	"net/http"
)

const (
	UploadSourcecodeQueue = "sourceCodeUploads"
	fileUploadStatus      = "File uploaded successfully"
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

	// uses RabbitMQ for messaging
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Printf("Could not connect to RabbitMQ: %s", err)
	}
	defer conn.Close()

	// creates a channel for publishing messages
	ch, err := conn.Channel()
	if err != nil {
		log.Printf("Could not open channel: %s", err)
	}
	defer ch.Close()

	// declares a queue to publish messages
	queue, err := ch.QueueDeclare(UploadSourcecodeQueue, false, false, false, false, nil)
	if err != nil {
		log.Printf("Could not declare queue: %s", err)
	}

	// publishes source code directory to the channel
	errPub := ch.Publish(
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(sourceCode),
		},
	)

	if errPub != nil {
		log.Printf("Could not publish the message: %s", errPub)
	}

	// response
	response := models.FileUploadResponse{}
	response.UploadStatus = fileUploadStatus

	return c.JSON(http.StatusOK, &response)
}

//func UploadSourceCodeHandler2(c echo.Context) error {
//	sourceCodeHeader, err := c.FormFile("sourceCode")
//	if err != nil {
//		log.Printf("Could not retrieve the file %v", err)
//	}
//
//	fileName := utils.CopyUploadedFile(sourceCodeHeader)
//	sourceCode := utils.UnZipFile(fileName)
//
//	// evaluates tech (language)
//	techDataSlice := utils.ReadFromTechDB(techDB)
//	tech := evaluators.EvaluateTechnologies(techDataSlice, sourceCode)
//
//	// check if dockerized
//	hasDockerized := evaluators.DockerizationEvaluator(sourceCode)
//
//	// check helm configs
//	hasServiceYaml, hasDeploymentYaml := evaluators.HelmConfigsEvaluator(sourceCode)
//
//	// evaluate database
//	databasesSlice := utils.ReadFromDatabaseDB(databaseDB)
//	databaseUsed := evaluators.EvaluateDatabases(sourceCode, databasesSlice, tech)
//
//	// evaluate service dependencies
//	serviceDependencies := utils.ReadFromServiceDB(serviceDB)
//	evaluators.EvaluateServiceDependencies(sourceCode, serviceDependencies, tech)
//	//fmt.Println(isService)
//
//	// response
//	response := models.FileUploadResponse{}
//	response.UploadStatus = fileUploadStatus
//	response.TechnologyUsed = techUsed + tech
//
//	if hasDockerized {
//		response.DockerizationStatus = hasDockerizedMsg
//	} else {
//		response.DockerizationStatus = hasNotDockerizedMsg
//	}
//
//	if hasServiceYaml {
//		response.ServiceYamlStatus = hasServiceYamlMsg
//	} else {
//		response.ServiceYamlStatus = noServiceYamlMsg
//	}
//
//	if hasDeploymentYaml {
//		response.DeploymentYamlStatus = hasDeploymentYamlMsg
//	} else {
//		response.DeploymentYamlStatus = noDeploymentYamlMsg
//	}
//
//	response.Database = databaseUsed
//
//	return c.JSON(http.StatusOK, &response)
//}

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

func CreateServiceDataHandler(c echo.Context) error {
	var request []models.ServiceData

	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		log.Fatal(err)
	}

	// database operations
	utils.CreateServiceDataDB(request)

	return c.JSON(http.StatusOK, &request)
}
