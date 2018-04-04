package svc

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
)

type SqsClient struct {
	svc sqsiface.SQSAPI
}

func NewSqsClient(region string) *SqsClient {
	return &SqsClient{
		svc: sqs.New(session.New(), aws.NewConfig().WithRegion(region)),
	}
}

func (c *SqsClient) GetQueueURL(queue string) (*sqs.GetQueueUrlOutput, error) {
	return c.svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: aws.String(queue),
	})
}
