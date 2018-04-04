package awsresolver

type GetSsmParameterOutput struct {
	Value *string `json:"value"`
}

type GetRdsClusterOutput struct {
	Endpoint       *string `json:"endpoint"`
	MasterUsername *string `json:"master_username"`
	Port           *int64  `json:"port"`
	ReaderEndpoint *string `json:"reader_endpoint"`
}

type GetSqsQueueURLOutput struct {
	QueueURL *string `json:"queue_url"`
}
