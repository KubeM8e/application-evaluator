package evaluators

import (
	"application-evaluator/models"
	"application-evaluator/pkg/utils"
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
	microservices := EvaluateServiceDependencies(sourceCode, serviceDependencies, language)

	// response
	response := models.EvaluationResponse{}
	response.Language = language
	response.Microservices = microservices

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

}
