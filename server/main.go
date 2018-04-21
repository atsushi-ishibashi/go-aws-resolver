package main

import (
	"flag"

	"github.com/atsushi-ishibashi/go-aws-resolver/server/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	port := flag.String("port", "8989", "port to listen")
	region := flag.String("region", "ap-northeast-1", "aws region")
	flag.Parse()

	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Recover())

	hdl := handler.NewAPIHandler(*region)
	e.GET("/rds/:cluster", hdl.RdsCluster)
	e.GET("/ssm/parameter", hdl.SsmParameter)
	e.GET("/sqs/:queue/url", hdl.SqsQueueURL)
	e.GET("/elasticache/:rg_id", hdl.ElastiCacheReplicationGroup)

	e.Logger.Fatal(e.Start(":" + *port))
}
