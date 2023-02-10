package models

// API

type CreateTechDataRequest struct {
	Technology string   `json:"technology"`
	Keywords   []string `json:"keywords"`
}

// Database

type Technology struct {
	Keywords []string `json:"keywords"`
}
