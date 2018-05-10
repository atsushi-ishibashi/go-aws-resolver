package awsresolver

type GetSsmParameterOutput struct {
	Value *string `json:"value"`
}

type GetRdsClusterOutput struct {
	Endpoint       *string `json:"endpoint"`
	MasterUsername *string `json:"master_username"`
	Port           *int64  `json:"port"`
	ReaderEndpoint *string `json:"reader_endpoint"`
}

type GetSqsQueueURLOutput struct {
	QueueURL *string `json:"queue_url"`
}

type GetElastiCacheReplicationGroupOutput struct {
	ConfigurationEndpoint string                  `json:"configuration_endpoint"`
	NodeGroups            []*ElastiCacheNodeGroup `json:"node_groups"`
}

type ElastiCacheNodeGroup struct {
	PrimaryEndpoint  string   `json:"configuration_endpoint,omitempty"`
	ReplicaEndpoints []string `json:"replica_endpoints,omitempty"`
}

type GetSecretsManagerSecretOutput struct {
	Value string `json:"value"`
}

type GetKMSKeyIDOutput struct {
	KeyID string `json:"key_id"`
}
