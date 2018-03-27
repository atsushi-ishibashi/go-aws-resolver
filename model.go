package awsresolver

type GetSsmParameterOutput struct {
	Value *string
}

type GetRdsClusterOutput struct {
	DatabaseName   *string
	Endpoint       *string
	MasterUsername *string
	Port           *int64
	ReaderEndpoint *string
}
