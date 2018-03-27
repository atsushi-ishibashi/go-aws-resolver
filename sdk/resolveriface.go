package awsresolver

type ResolverIface interface {
	GetSsmParameter(name string) (string, error)
}
