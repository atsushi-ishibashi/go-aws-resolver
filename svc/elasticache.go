package svc

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elasticache"
	"github.com/aws/aws-sdk-go/service/elasticache/elasticacheiface"
)

type ElastiCacheClient struct {
	svc elasticacheiface.ElastiCacheAPI
}

func NewElastiCacheClient(region string) *ElastiCacheClient {
	return &ElastiCacheClient{
		svc: elasticache.New(session.New(), aws.NewConfig().WithRegion(region)),
	}
}

func (c *ElastiCacheClient) GetReplicationGroup(rgid string) (*elasticache.ReplicationGroup, error) {
	input := &elasticache.DescribeReplicationGroupsInput{
		ReplicationGroupId: aws.String(rgid),
	}
	resp, err := c.svc.DescribeReplicationGroups(input)
	if err != nil {
		return nil, err
	}
	if len(resp.ReplicationGroups) != 1 {
		return nil, fmt.Errorf("num of ReplicationGroups(%s) expected 1, got %d", rgid, len(resp.ReplicationGroups))
	}
	return resp.ReplicationGroups[0], nil
}
