package models

// RabbitMQ

type ProjectSource struct {
	ProjectId  string `json:"projectId"`
	Sourcecode string `json:"sourcecode"`
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
