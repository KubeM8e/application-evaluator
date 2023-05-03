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

type AppData struct {
	Id            string                  `json:"id,omitempty" yaml:"id" bson:"id"`
	AppName       string                  `json:"appName,omitempty" yaml:"appName" bson:"appName,omitempty"`
	Description   string                  `json:"description,omitempty"yaml:"description" bson:"description,omitempty"`
	ImageURL      string                  `json:"imageUrl,omitempty"yaml:"imageURL" bson:"imageURL,omitempty"`
	Version       string                  `json:"version,omitempty" yaml:"version" bson:"version,omitempty"`
	HostName      string                  `json:"hostName,omitempty" yaml:"hostName" bson:"hostName,omitempty"`
	ClusterURL    string                  `json:"clusterURL,omitempty" yaml:"clusterURL" bson:"clusterURL,omitempty"`
	ClusterIPs    []string                `json:"clusterIPs,omitempty" yaml:"clusterIPs" bson:"clusterIPs,omitempty"`
	Microservices map[string]Microservice `json:"microservices,omitempty" yaml:"microservices" bson:"microservices,omitempty"`
	Monitoring    bool                    `json:"monitoring,omitempty" yaml:"monitoring" bson:"monitoring,omitempty"`
}

type Microservice struct {
	ServiceName   string                `json:"serviceName,omitempty" yaml:"serviceName" bson:"serviceName,omitempty"`
	Configs       []string              `json:"configs,omitempty" yaml:"configs" bson:"configs,omitempty"`
	AvgReplicas   int                   `json:"avgReplicas,omitempty" yaml:"avgReplicas" bson:"avgReplicas,omitempty"`
	MinReplicas   int                   `json:"minReplicas,omitempty" yaml:"minReplicas" bson:"minReplicas,omitempty"`
	MaxReplicas   int                   `json:"maxReplicas,omitempty" yaml:"maxReplicas" bson:"maxReplicas,omitempty"`
	MaxCPU        string                `json:"maxCPU,omitempty" yaml:"maxCPU" bson:"maxCPU,omitempty"`
	MaxMemory     string                `json:"maxMemory,omitempty" yaml:"maxMemory"bson:"maxMemory,omitempty"`
	DockerImage   string                `json:"dockerImage,omitempty" yaml:"dockerImage"bson:"dockerImage",omitempty`
	ContainerPort int                   `json:"containerPort,omitempty" yaml:"containerPort"bson:"containerPort,omitempty"`
	Envs          map[string]EnvRequest `json:"envs,omitempty" yaml:"envs"bson:"envs,omitempty"`
}

type EnvRequest struct {
	Name  string `json:"name,omitempty" yaml:"name"bson:"name,omitempty"`
	Value string `json:"value,omitempty" yaml:"value"bson:"value,omitempty"`
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
