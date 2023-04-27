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
	UploadStatus         string `json:"uploadStatus"`
	TechnologyUsed       string `json:"technologyUsed"`
	Database             string `json:"database"`
	DockerizationStatus  string `json:"dockerizationStatus"`
	ServiceYamlStatus    string `json:"serviceYamlStatus"`
	DeploymentYamlStatus string `json:"deploymentYamlStatus"`
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
