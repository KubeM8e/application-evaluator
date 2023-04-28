package models

// API

type TechData struct {
	Technology string   `json:"technology"`
	Keywords   []string `json:"keywords"`
}

type DBData struct {
	Database string   `json:"database"`
	Language string   `json:"language"`
	Keywords []string `json:"keywords"`
}

type ServiceData struct {
	Language string   `json:"language"`
	Keywords []string `json:"keywords"`
}

type FileUploadResponse struct {
	UploadStatus string `json:"uploadStatus"`
}

type EvaluationResponse struct {
	Language                string `json:"language"`
	Database                string `json:"database"`
	HasDockerized           bool   `json:"hasDockerized"`
	HasKubernetesService    bool   `json:"hasKubernetesService"`
	HasKubernetesDeployment bool   `json:"hasKubernetesDeployment"`
}

// Database

type Technology struct {
	Keywords []string `json:"keywords"`
}

type Database struct {
	Language string   `json:"language"`
	Keywords []string `json:"keywords"`
}

// OpenAI

type DavinciRequest struct {
	Model       string `json:"model"`
	Prompt      string `json:"prompt"`
	MaxTokens   int    `json:"max_tokens"`
	Temperature int    `json:"temperature"`
}

type DavinciResponse struct {
	Id      string `json:"id"`
	Choices []struct {
		Text string `json:"text"`
	} `json:"choices"`
}

type ServiceDependency struct {
	Microservice string `json:"microservice"`
	DependsOn    string `json:"dependsOn"`
	URL          string `json:"url"`
}
