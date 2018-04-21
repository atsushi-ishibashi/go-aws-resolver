package awsresolver

type ResolverIface interface {
	GetSsmParameter(name string) (*GetSsmParameterOutput, error)
	GetSqsQueueURL(queue string) (*GetSqsQueueURLOutput, error)
	GetRdsCluster(cluster string) (*GetRdsClusterOutput, error)
	GetElastiCacheReplicationGroup(replicationGroupID string) (*GetElastiCacheReplicationGroupOutput, error)
}
