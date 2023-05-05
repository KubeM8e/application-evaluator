package models

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

type SuccessfulResponse struct {
	Status string `json:"status"`
}

type EvaluationResponse struct {
	Language                string   `json:"language,omitempty" bson:"language,omitempty"`
	Database                string   `json:"database,omitempty" bson:"database,omitempty"`
	HasDockerized           bool     `json:"hasDockerized" bson:"hasDockerized"`
	HasKubernetesService    bool     `json:"hasKubernetesService" bson:"hasKubernetesService"`
	HasKubernetesDeployment bool     `json:"hasKubernetesDeployment" bson:"hasKubernetesDeployment"`
	Microservices           []string `json:"microservices,omitempty" bson:"microservices,omitempty"`
}

type ProjectData struct {
	ProjectId        string                  `json:"projectId,omitempty" yaml:"projectId" bson:"projectId,omitempty"`
	AppName          string                  `json:"appName,omitempty" yaml:"appName" bson:"appName,omitempty"`
	Description      string                  `json:"description,omitempty" yaml:"description" bson:"description,omitempty"`
	ImageURL         string                  `json:"imageUrl,omitempty" yaml:"imageURL" bson:"imageURL,omitempty"`
	Version          string                  `json:"version,omitempty" yaml:"version" bson:"version,omitempty"`
	HostName         string                  `json:"hostName,omitempty" yaml:"hostName" bson:"hostName,omitempty"`
	ClusterURL       string                  `json:"clusterURL,omitempty" yaml:"clusterURL" bson:"clusterURL,omitempty"`
	ClusterIPs       []string                `json:"clusterIPs,omitempty" yaml:"clusterIPs" bson:"clusterIPs,omitempty"`
	Microservices    map[string]Microservice `json:"microservices,omitempty" yaml:"microservices" bson:"microservices,omitempty"`
	Monitoring       bool                    `json:"monitoring,omitempty" yaml:"monitoring" bson:"monitoring,omitempty"`
	EvaluationResult EvaluationResponse      `json:"evaluationResult,omitempty" yaml:"evaluationResult" bson:"evaluationResult,omitempty"`
}

type Microservice struct {
	ServiceName       string                `json:"serviceName,omitempty" yaml:"serviceName" bson:"serviceName,omitempty"`
	KubeConfigType    []string              `json:"kubeConfigType,omitempty" yaml:"kubeConfigType" bson:"kubeConfigType,omitempty"`
	Configs           []string              `json:"configs,omitempty" yaml:"configs" bson:"configs,omitempty"`
	AvgReplicas       int                   `json:"avgReplicas,omitempty" yaml:"avgReplicas" bson:"avgReplicas,omitempty"`
	MinReplicas       int                   `json:"minReplicas,omitempty" yaml:"minReplicas" bson:"minReplicas,omitempty"`
	MaxReplicas       int                   `json:"maxReplicas,omitempty" yaml:"maxReplicas" bson:"maxReplicas,omitempty"`
	MaxCPU            string                `json:"maxCPU,omitempty" yaml:"maxCPU" bson:"maxCPU,omitempty"`
	MaxMemory         string                `json:"maxMemory,omitempty" yaml:"maxMemory" bson:"maxMemory,omitempty"`
	DockerImage       string                `json:"dockerImage,omitempty" yaml:"dockerImage" bson:"dockerImage,omitempty"`
	ContainerPort     int                   `json:"containerPort,omitempty" yaml:"containerPort" bson:"containerPort,omitempty"`
	Envs              map[string]EnvRequest `json:"envs,omitempty" yaml:"envs" bson:"envs,omitempty"`
	ServiceEvaluation ServiceEvaluation     `json:"serviceEvaluation,omitempty" yaml:"serviceEvaluation" bson:"serviceEvaluation,omitempty"`
}

type EnvRequest struct {
	Name  string `json:"name,omitempty" yaml:"name" bson:"name,omitempty"`
	Value string `json:"value,omitempty" yaml:"value" bson:"value,omitempty"`
}

type ServiceEvaluation struct {
	HtmlCssJsFiles             bool          `json:"htmlCssJsFiles,omitempty" yaml:"htmlCssJsFiles" bson:"htmlCssJsFiles,omitempty"`
	Php                        bool          `json:"php,omitempty" yaml:"php" bson:"php,omitempty"`
	Ruby                       bool          `json:"ruby,omitempty" yaml:"ruby" bson:"ruby,omitempty"`
	Angular                    bool          `json:"angular,omitempty" yaml:"angular" bson:"angular,omitempty"`
	URLRouting                 bool          `json:"URLRouting,omitempty" yaml:"URLRouting" bson:"URLRouting,omitempty"`
	HandleHTTPRequest          bool          `json:"handleHTTPRequest,omitempty" yaml:"handleHTTPRequest" bson:"handleHTTPRequest,omitempty"`
	DB                         bool          `json:"db,omitempty" yaml:"DB" bson:"DB,omitempty"`
	FileSystems                bool          `json:"fileSystems,omitempty" yaml:"fileSystems" bson:"fileSystems,omitempty"`
	Caching                    bool          `json:"caching,omitempty" yaml:"caching" bson:"caching,omitempty"`
	SessionManagement          bool          `json:"sessionManagement,omitempty" yaml:"sessionManagement" bson:"sessionManagement,omitempty"`
	MessageQueues              MessageQueues `json:"messageQueues,omitempty" yaml:"messageQueues" bson:"messageQueues,omitempty"`
	DeliverMsgs                bool          `json:"deliverMsgs,omitempty" yaml:"deliverMsgs" bson:"deliverMsgs,omitempty"`
	IntegrateMsgPlatforms      bool          `json:"integrateMsgPlatforms,omitempty" yaml:"integrateMsgPlatforms" bson:"integrateMsgPlatforms,omitempty"`
	MessageFormats             bool          `json:"messageFormats,omitempty" yaml:"messageFormats" bson:"messageFormats,omitempty"`
	DataTransfer               bool          `json:"dataTransfer,omitempty" yaml:"dataTransfer" bson:"dataTransfer,omitempty"`
	BackupFrequency            bool          `json:"backupFrequency,omitempty" yaml:"backupFrequency" bson:"backupFrequency,omitempty"`
	BackupDestination          bool          `json:"backupDestination,omitempty" yaml:"backupDestination" bson:"backupDestination,omitempty"`
	DataTransform              bool          `json:"dataTransform,omitempty" yaml:"dataTransform" bson:"dataTransform,omitempty"`
	CollectData                bool          `json:"collectData,omitempty" yaml:"collectData" bson:"collectData,omitempty"`
	GenerateReports            bool          `json:"generateReports,omitempty" yaml:"generateReports" bson:"generateReports,omitempty"`
	ReportVisualization        bool          `json:"reportVisualization,omitempty" yaml:"reportVisualization" bson:"reportVisualization,omitempty"`
	IntegrateMonitoringTools   bool          `json:"integrateMonitoringTools,omitempty" yaml:"integrateMonitoringTools" bson:"integrateMonitoringTools,omitempty"`
	AnalyzeLogs                bool          `json:"analyzeLogs,omitempty" yaml:"analyzeLogs" bson:"analyzeLogs,omitempty"`
	CollectPerformanceMetrics  bool          `json:"collectPerformanceMetrics,omitempty" yaml:"collectPerformanceMetrics" bson:"collectPerformanceMetrics,omitempty"`
	SendAlertsAndNotifications bool          `json:"sendAlertsAndNotifications,omitempty" yaml:"sendAlertsAndNotifications" bson:"sendAlertsAndNotifications,omitempty"`
	HandleNetworkCommunication bool          `json:"handleNetworkCommunication,omitempty" yaml:"handleNetworkCommunication" bson:"handleNetworkCommunication,omitempty"`
	NetworkProtocolLibraries   bool          `json:"networkProtocolLibraries,omitempty" yaml:"networkProtocolLibraries" bson:"networkProtocolLibraries,omitempty"`
	Routing                    bool          `json:"routing,omitempty" yaml:"routing" bson:"routing,omitempty"`
	LoadBalancing              bool          `json:"loadBalancing,omitempty" yaml:"loadBalancing" bson:"loadBalancing,omitempty"`
	FirewallRules              bool          `json:"firewallRules,omitempty" yaml:"firewallRules" bson:"firewallRules,omitempty"`
}

type MessageQueues struct {
	SendAndReceiveMsgs bool `json:"sendAndReceiveMsgs,omitempty" yaml:"sendAndReceiveMsgs" bson:"sendAndReceiveMsgs,omitempty"`
	StoreMsgs          bool `json:"storeMsgs,omitempty" yaml:"storeMsgs" bson:"storeMsgs,omitempty"`
}
