package svc

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/aws/aws-sdk-go/service/secretsmanager/secretsmanageriface"
)

type SecretsManagerClient struct {
	svc secretsmanageriface.SecretsManagerAPI
}

func NewSecretsManagerClient(region string) *SecretsManagerClient {
	return &SecretsManagerClient{
		svc: secretsmanager.New(session.New(), aws.NewConfig().WithRegion(region)),
	}
}

func (c *SecretsManagerClient) GetSecretValue(name string) (*secretsmanager.GetSecretValueOutput, error) {
	return c.svc.GetSecretValue(&secretsmanager.GetSecretValueInput{
		SecretId: aws.String(name),
	})
}
