package handler

import "github.com/atsushi-ishibashi/go-aws-resolver/svc"

type APIHandler struct {
	ssmClient *svc.SsmClient
}

func NewAPIHandler(region string) *APIHandler {
	return &APIHandler{
		ssmClient: svc.NewSsmClient(region),
	}
}
