package ruleEngine

import (
	"application-evaluator/models"
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	"log"
)

const (
	nameKube    = "Kubernetes"
	versionKube = "1.0.0"
)

func EvaluateKubeConfigType(evalObj models.ServiceEvaluation) models.ServiceEvaluation {
	// adds fact to data context
	dataCtx := ast.NewDataContext()

	err := dataCtx.Add("Eval", &evalObj)
	if err != nil {
		log.Printf("Could not add eval object to data context kube: %s", err)
	}

	// RULE
	knowledgeLibrary := ast.NewKnowledgeLibrary()
	ruleBuilder := builder.NewRuleBuilder(knowledgeLibrary)

	fileRes := pkg.NewFileResource("rules/kubernetes.grl")
	err = ruleBuilder.BuildRuleFromResource(nameKube, versionKube, fileRes)
	if err != nil {
		log.Printf("Could not build rules kube: %s", err)
	}

	// get instance of KB from the knowledge library
	knowledgeBase := knowledgeLibrary.NewKnowledgeBaseInstance(nameKube, versionKube)

	// check the KB for the fact given
	en := engine.NewGruleEngine()

	err = en.Execute(dataCtx, knowledgeBase)
	if err != nil {
		log.Printf("Could not execute rule engone kube: %s", err)
	}

	return evalObj

}
