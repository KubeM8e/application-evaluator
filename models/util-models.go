package models

// API

type CreateTechDataRequest struct {
	Technology string   `json:"technology"`
	Keywords   []string `json:"keywords"`
}

type FileUploadResponse struct {
	UploadStatus        string `json:"uploadStatus"`
	TechnologyUsed      string `json:"technologyUsed"`
	DockerizationStatus string `json:"dockerizationStatus"`
}

// Database

type Technology struct {
	Keywords []string `json:"keywords"`
}
