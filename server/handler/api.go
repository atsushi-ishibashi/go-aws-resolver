package handler

import (
	"net/http"

	awsresolver "github.com/atsushi-ishibashi/go-aws-resolver"
	"github.com/labstack/echo"
)

//APIHandler handler
type APIHandler struct {
	resolver awsresolver.ResolverIface
}

//NewAPIHandler NewFunction for APIHandler
func NewAPIHandler(region string) *APIHandler {
	return &APIHandler{
		resolver: awsresolver.NewResolver(region),
	}
}

//RDSCluster GET /rds/:cluster, retrun GetRdsClusterOutput
func (h *APIHandler) RdsCluster(c echo.Context) error {
	resp, err := h.resolver.GetRdsCluster(c.Param("cluster"))
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}

//SsmParameter GET /ssm/parameter?name=/hoge/fuga, retrun GetSsmParameterOutput
func (h *APIHandler) SsmParameter(c echo.Context) error {
	resp, err := h.resolver.GetSsmParameter(c.QueryParam("name"))
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}

//SqsQueueURL GET /sqs/:queue/url, retrun GetSqsQueueURLOutput
func (h *APIHandler) SqsQueueURL(c echo.Context) error {
	resp, err := h.resolver.GetSqsQueueURL(c.Param("queue"))
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}

//ElastiCacheReplicationGroup GET /elasticache/:rg_id, retrun GetElastiCacheReplicationGroupOutput
func (h *APIHandler) ElastiCacheReplicationGroup(c echo.Context) error {
	resp, err := h.resolver.GetElastiCacheReplicationGroup(c.Param("rg_id"))
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}
