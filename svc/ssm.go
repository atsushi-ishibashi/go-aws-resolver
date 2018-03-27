package svc

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/aws/aws-sdk-go/service/ssm/ssmiface"
)

type SsmClient struct {
	svc ssmiface.SSMAPI
}

func NewSsmClient(region string) *SsmClient {
	return &SsmClient{
		svc: ssm.New(session.New(), aws.NewConfig().WithRegion(region)),
	}
}

func (c *SsmClient) GetParameter(name string) (string, error) {
	resp, err := c.svc.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String(name),
		WithDecryption: aws.Bool(true),
	})
	if err != nil {
		return "", err
	}
	return *resp.Parameter.Value, nil
}
