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
	Language                string   `json:"language"`
	Database                string   `json:"database"`
	HasDockerized           bool     `json:"hasDockerized"`
	HasKubernetesService    bool     `json:"hasKubernetesService"`
	HasKubernetesDeployment bool     `json:"hasKubernetesDeployment"`
	Microservices           []string `json:"microservices"`
}

type ConfigurationRequest struct {
	AppName       string                  `json:"appName" yaml:"appName"`
	Version       string                  `json:"version" yaml:"version"`
	HostName      string                  `json:"hostName" yaml:"hostName"`
	ClusterURL    string                  `json:"clusterURL" yaml:"clusterURL"`
	ClusterIPs    []string                `json:"clusterIPs" yaml:"clusterIPs"`
	Microservices map[string]Microservice `json:"microservices" yaml:"microservices"`
	Monitoring    bool                    `json:"monitoring" yaml:"monitoring"`
}

type Microservice struct {
	ServiceName   string                `json:"serviceName" yaml:"serviceName"`
	Configs       []string              `json:"configs" yaml:"configs"`
	AvgReplicas   int                   `json:"avgReplicas" yaml:"avgReplicas"`
	MinReplicas   int                   `json:"minReplicas" yaml:"minReplicas"`
	MaxReplicas   int                   `json:"maxReplicas" yaml:"maxReplicas"`
	MaxCPU        string                `json:"maxCPU" yaml:"maxCPU"`
	MaxMemory     string                `json:"maxMemory" yaml:"maxMemory"`
	DockerImage   string                `json:"dockerImage" yaml:"dockerImage"`
	ContainerPort int                   `json:"containerPort" yaml:"containerPort"`
	Envs          map[string]EnvRequest `json:"envs" yaml:"envs"`
}

type EnvRequest struct {
	Name  string `json:"name" yaml:"name"`
	Value string `json:"value" yaml:"value"`
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
