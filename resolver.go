package awsresolver

import "github.com/atsushi-ishibashi/go-aws-resolver/svc"

type Resolver struct {
	ssmClient *svc.SsmClient
	sqsClient *svc.SqsClient
	rdsClient *svc.RdsClient
}

func NewResolver(region string) *Resolver {
	return &Resolver{
		ssmClient: svc.NewSsmClient(region),
		sqsClient: svc.NewSqsClient(region),
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
		MasterUsername: resp.MasterUsername,
		Port:           resp.Port,
		ReaderEndpoint: resp.ReaderEndpoint,
	}, nil
}

func (r *Resolver) GetSqsQueueURL(queue string) (*GetSqsQueueURLOutput, error) {
	resp, err := r.sqsClient.GetQueueURL(queue)
	if err != nil {
		return nil, err
	}
	return &GetSqsQueueURLOutput{
		QueueURL: resp.QueueUrl,
	}, nil
}
