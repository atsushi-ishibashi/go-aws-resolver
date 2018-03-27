package main

import (
	"flag"

	"github.com/atsushi-ishibashi/go-aws-resolver/server/handler"
	"github.com/labstack/echo"
)

func main() {
	port := flag.String("port", "8989", "port to listen")
	region := flag.String("region", "ap-northeast-1", "aws region")
	flag.Parse()

	e := echo.New()

	_ = handler.NewAPIHandler(*region)

	e.Logger.Fatal(e.Start(":" + *port))
}
