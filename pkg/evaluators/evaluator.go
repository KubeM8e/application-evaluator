package evaluators

import (
	"application-evaluator/models"
	"application-evaluator/pkg/project"
	"application-evaluator/pkg/ruleEngine"
	"application-evaluator/pkg/utils"
)

const (
	techDB     = "Technologies"
	databaseDB = "Databases"
	serviceDB  = "ServiceDependencies"
)

func EvaluateSourcecode(projectSource models.ProjectSource) {
	sourceCode := projectSource.Sourcecode
	evalResult := models.EvaluationResponse{}

	// evaluates tech (language)
	techDataSlice := utils.ReadFromTechDB(techDB)
	language := EvaluateTechnologies(techDataSlice, sourceCode)

	// checks if dockerized
	hasDockerized := DockerizationEvaluator(sourceCode)

	// checks helm configs
	hasServiceYaml, hasDeploymentYaml := HelmConfigsEvaluator(sourceCode)

	// evaluates database
	//databasesSlice := utils.ReadFromDatabaseDB(databaseDB)
	//databaseUsed := EvaluateDatabases(sourceCode, databasesSlice, language)

	// evaluates service dependencies
	serviceDependencies := utils.ReadFromServiceDB(serviceDB)
	microservices, databaseUsed, microservicesMap := EvaluateMicroservices(sourceCode, serviceDependencies, language)
	evalResult.Microservices = microservices

	// evalResult

	evalResult.Language = language
	evalResult.Database = databaseUsed

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

	// core rule based system
	for k, microservice := range microservicesMap {
		//microservice.ServiceEvaluation = ruleEngine.EvaluateServiceType(microservice.ServiceEvaluation)
		eval := ruleEngine.EvaluateServiceType(microservice.ServiceEvaluation)
		evalWithKubeConfig := ruleEngine.EvaluateKubeConfigType(eval)
		microservice.ServiceEvaluation = evalWithKubeConfig
		microservicesMap[k] = microservice

	}

	saveEvaluationResults(projectSource.ProjectId, evalResult, microservicesMap)

}

func saveEvaluationResults(projectId string, evaluationResult models.EvaluationResponse, microservicesMap map[string]models.Microservice) {
	projectData := project.GetProject(projectId)

	projectData.EvaluationResult = evaluationResult
	projectData.Microservices = microservicesMap

	project.UpdateProject(projectId, projectData)
}
