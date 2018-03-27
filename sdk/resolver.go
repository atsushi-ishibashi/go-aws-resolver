package awsresolver

import "github.com/atsushi-ishibashi/go-aws-resolver/svc"

type Resolver struct {
	ssmClient *svc.SsmClient
}

func NewResolver(region string) *Resolver {
	return &Resolver{
		ssmClient: svc.NewSsmClient(region),
	}
}

func (r *Resolver) GetSsmParameter(name string) (string, error) {
	return r.ssmClient.GetParameter(name)
}
