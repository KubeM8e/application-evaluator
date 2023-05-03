package evaluators

import (
	"application-evaluator/models"
	"application-evaluator/pkg/project"
	"application-evaluator/pkg/utils"
)

const (
	techDB     = "Technologies"
	databaseDB = "Databases"
	serviceDB  = "ServiceDependencies"
)

func EvaluateSourcecode(projectSource models.ProjectSource) {
	sourceCode := projectSource.Sourcecode

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

	// evalResult
	evalResult := models.EvaluationResponse{}
	evalResult.Language = language
	evalResult.Database = databaseUsed

	// sets microservices
	evalResult.Microservices = microservices

	if hasDockerized {
		evalResult.HasDockerized = true
	} else {
		evalResult.HasDockerized = false
	}

	if hasServiceYaml {
		evalResult.HasKubernetesService = true
	} else {
		evalResult.HasKubernetesService = false
	}

	if hasDeploymentYaml {
		evalResult.HasKubernetesDeployment = true
	} else {
		evalResult.HasKubernetesDeployment = false
	}

	evalResult.Database = databaseUsed

	saveEvaluationResults(projectSource.ProjectId, evalResult)

}

func saveEvaluationResults(projectId string, evaluationResult models.EvaluationResponse) {
	projectData := project.GetProject(projectId)

	projectData.EvaluationResult = evaluationResult

	project.UpdateProject(projectId, projectData)
}
