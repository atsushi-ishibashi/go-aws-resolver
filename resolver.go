package awsresolver

import "github.com/atsushi-ishibashi/go-aws-resolver/svc"

type Resolver struct {
	ssmClient *svc.SsmClient
	rdsClient *svc.RdsClient
}

func NewResolver(region string) *Resolver {
	return &Resolver{
		ssmClient: svc.NewSsmClient(region),
		rdsClient: svc.NewRdsClient(region),
	}
}

func (r *Resolver) GetSsmParameter(name string) (*GetSsmParameterOutput, error) {
	resp, err := r.ssmClient.GetParameter(name)
	if err != nil {
		return nil, err
	}
	return &GetSsmParameterOutput{
		Value: resp.Parameter.Value,
	}, nil
}

func (r *Resolver) GetRdsCluster(cluster string) (*GetRdsClusterOutput, error) {
	resp, err := r.rdsClient.GetCluster(cluster)
	if err != nil {
		return nil, err
	}
	return &GetRdsClusterOutput{
		DatabaseName:   resp.DatabaseName,
		Endpoint:       resp.Endpoint,
		MasterUsername: resp.MasterUsername,
		Port:           resp.Port,
		ReaderEndpoint: resp.ReaderEndpoint,
	}, nil
}
