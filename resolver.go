package awsresolver

import (
	"fmt"

	"github.com/atsushi-ishibashi/go-aws-resolver/svc"
)

type Resolver struct {
	ssmClient            *svc.SsmClient
	sqsClient            *svc.SqsClient
	rdsClient            *svc.RdsClient
	elasticacheClient    *svc.ElastiCacheClient
	secretsManagerClient *svc.SecretsManagerClient
}

func NewResolver(region string) *Resolver {
	return &Resolver{
		ssmClient:            svc.NewSsmClient(region),
		sqsClient:            svc.NewSqsClient(region),
		rdsClient:            svc.NewRdsClient(region),
		elasticacheClient:    svc.NewElastiCacheClient(region),
		secretsManagerClient: svc.NewSecretsManagerClient(region),
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
		Endpoint:       resp.Endpoint,
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

func (r *Resolver) GetElastiCacheReplicationGroup(replicationGroupID string) (*GetElastiCacheReplicationGroupOutput, error) {
	resp, err := r.elasticacheClient.GetReplicationGroup(replicationGroupID)
	if err != nil {
		return nil, err
	}

	flattenEndpoint := func(domain *string, port *int64) string {
		if domain == nil || port == nil {
			return ""
		}
		return fmt.Sprintf("%s:%d", *domain, *port)
	}

	result := &GetElastiCacheReplicationGroupOutput{}
	if ce := resp.ConfigurationEndpoint; ce != nil {
		result.ConfigurationEndpoint = flattenEndpoint(ce.Address, ce.Port)
	}
	replicas := make([]*ElastiCacheNodeGroup, 0, len(resp.NodeGroups))
	for _, ng := range resp.NodeGroups {
		ecng := &ElastiCacheNodeGroup{
			PrimaryEndpoint: flattenEndpoint(ng.PrimaryEndpoint.Address, ng.PrimaryEndpoint.Port),
		}
		rrEndpoints := make([]string, 0, len(ng.NodeGroupMembers))
		for _, ngm := range ng.NodeGroupMembers {
			if *ngm.CurrentRole == "replica" {
				rrEndpoints = append(rrEndpoints, flattenEndpoint(ngm.ReadEndpoint.Address, ngm.ReadEndpoint.Port))
			}
		}
		ecng.ReplicaEndpoints = rrEndpoints
		replicas = append(replicas, ecng)
	}
	result.NodeGroups = replicas
	return result, nil
}

func (r *Resolver) GetSecretsManagerSecret(name string) (*GetSecretsManagerSecretOutput, error) {
	resp, err := r.secretsManagerClient.GetSecretValue(name)
	if err != nil {
		return nil, err
	}
	return &GetSecretsManagerSecretOutput{
		Value: *resp.SecretString,
	}, nil
}
