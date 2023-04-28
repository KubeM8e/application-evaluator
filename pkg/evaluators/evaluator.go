package evaluators

import (
	"application-evaluator/models"
	"application-evaluator/pkg/utils"
	"fmt"
)

const (
	techDB     = "Technologies"
	databaseDB = "Databases"
	serviceDB  = "ServiceDependencies"
)

func EvaluateSourcecode(sourceCode string) {
	// evaluates tech (language)
	techDataSlice := utils.ReadFromTechDB(techDB)
	language := EvaluateTechnologies(techDataSlice, sourceCode)

	// check if dockerized
	hasDockerized := DockerizationEvaluator(sourceCode)

	// check helm configs
	hasServiceYaml, hasDeploymentYaml := HelmConfigsEvaluator(sourceCode)

	// evaluate database
	databasesSlice := utils.ReadFromDatabaseDB(databaseDB)
	databaseUsed := EvaluateDatabases(sourceCode, databasesSlice, language)

	// evaluate service dependencies
	serviceDependencies := utils.ReadFromServiceDB(serviceDB)
	EvaluateServiceDependencies(sourceCode, serviceDependencies, language)
	//fmt.Println(isService)

	// response
	response := models.EvaluationResponse{}
	response.Language = language

	if hasDockerized {
		response.HasDockerized = true
	} else {
		response.HasDockerized = false
	}

	if hasServiceYaml {
		response.HasKubernetesService = true
	} else {
		response.HasKubernetesService = false
	}

	if hasDeploymentYaml {
		response.HasKubernetesDeployment = true
	} else {
		response.HasKubernetesDeployment = false
	}

	response.Database = databaseUsed

	fmt.Println("RESP: ", response)

	//return c.JSON(http.StatusOK, &response)
}
