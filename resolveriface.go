package awsresolver

type ResolverIface interface {
	GetSsmParameter(name string) (*GetSsmParameterOutput, error)
	GetRdsCluster(cluster string) (*GetRdsClusterOutput, error)
}
