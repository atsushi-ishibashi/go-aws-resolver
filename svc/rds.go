package svc

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/aws/aws-sdk-go/service/rds/rdsiface"
)

type RdsClient struct {
	svc rdsiface.RDSAPI
}

func NewRdsClient(region string) *RdsClient {
	return &RdsClient{
		svc: rds.New(session.New(), aws.NewConfig().WithRegion(region)),
	}
}

func (c *RdsClient) GetCluster(cluster string) (*rds.DBCluster, error) {
	resp, err := c.svc.DescribeDBClusters(&rds.DescribeDBClustersInput{
		DBClusterIdentifier: aws.String(cluster),
	})
	if err != nil {
		return nil, err
	}
	if len(resp.DBClusters) != 1 {
		return nil, fmt.Errorf("rds cluster (%s) found %d, expected 1", cluster, len(resp.DBClusters))
	}
	return resp.DBClusters[0], nil
}
