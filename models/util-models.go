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

type FileUploadResponse struct {
	UploadStatus         string `json:"uploadStatus"`
	TechnologyUsed       string `json:"technologyUsed"`
	DockerizationStatus  string `json:"dockerizationStatus"`
	ServiceYamlStatus    string `json:"serviceYamlStatus"`
	DeploymentYamlStatus string `json:"deploymentYamlStatus"`
	Database             string `json:"database"`
}

// Database

type Technology struct {
	Keywords []string `json:"keywords"`
}

type Database struct {
	Language string   `json:"language"`
	Keywords []string `json:"keywords"`
}
