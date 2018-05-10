package awsresolver

type ResolverIface interface {
	GetSsmParameter(name string) (*GetSsmParameterOutput, error)
	GetSqsQueueURL(queue string) (*GetSqsQueueURLOutput, error)
	GetRdsCluster(cluster string) (*GetRdsClusterOutput, error)
	GetElastiCacheReplicationGroup(replicationGroupID string) (*GetElastiCacheReplicationGroupOutput, error)
	GetSecretsManagerSecret(name string) (*GetSecretsManagerSecretOutput, error)
	GetKMSKeyID(alias string) (*GetKMSKeyIDOutput, error)
}
