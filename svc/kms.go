package svc

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/aws/aws-sdk-go/service/kms/kmsiface"
)

type KMSClient struct {
	svc kmsiface.KMSAPI
}

func NewKMSClient(region string) *KMSClient {
	return &KMSClient{
		svc: kms.New(session.New(), aws.NewConfig().WithRegion(region)),
	}
}

func (c *KMSClient) GetKey(alias string) (*kms.DescribeKeyOutput, error) {
	return c.svc.DescribeKey(&kms.DescribeKeyInput{
		KeyId: aws.String(alias),
	})
}
