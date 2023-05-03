package handlers

import (
	"application-evaluator/models"
	"application-evaluator/pkg/project"
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

func CreateProjectHandler(c echo.Context) error {
	request := models.ProjectData{}
	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		log.Printf("Could not retrieve project data: %s", err)
	}

	// creates project
	createdAppData := project.CreateProject(request)

	return c.JSON(http.StatusOK, &createdAppData)
}

func UploadSourceCodeHandler(c echo.Context) error {
	sourceCodeHeader, err := c.FormFile("sourceCode")
	if err != nil {
		log.Printf("Could not retrieve the file %v", err)
	}

	projectId := c.FormValue("id")
	if err != nil {
		log.Printf("Could not retrieve the id %v", err)
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

	// publishes source code directory to the channel - evaluate() function will consume it
	source := models.ProjectSource{
		ProjectId:  projectId,
		Sourcecode: sourceCode,
	}

	sourceJson, err := json.Marshal(source)
	if err != nil {
		// Handle JSON serialization error
	}
	errPub := ch.Publish(
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        sourceJson,
		},
	)

	if errPub != nil {
		log.Printf("Could not publish the message: %s", errPub)
	}

	// response
	response := models.SuccessfulResponse{}
	response.Status = fileUploadStatus

	return c.JSON(http.StatusOK, &response)
}

func GetAllProjectsHandler(c echo.Context) error {
	projects := project.GetAllProjects()

	return c.JSON(http.StatusOK, projects)
}

func GetProjectHandler(c echo.Context) error {
	projectId := c.Param("id")
	projectData := project.GetProject(projectId)

	return c.JSON(http.StatusOK, &projectData)
}

func UpdateProjectHandler(c echo.Context) error {
	projectId := c.Param("id")
	var request models.ProjectData

	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		log.Println(err)
	}

	isUpdateSuccessful := project.UpdateProject(projectId, request)

	response := models.SuccessfulResponse{}

	if isUpdateSuccessful {
		response.Status = "Unsuccessful"
	} else {
		response.Status = "Successful"
	}

	return c.JSON(http.StatusOK, &response)

}

// Utility API handlers

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

// Old version

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
//	//log.Println(isService)
//
//	// response
//	response := models.SuccessfulResponse{}
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
