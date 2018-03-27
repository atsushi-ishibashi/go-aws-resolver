package main

import (
	"flag"

	"github.com/labstack/echo"
)

func main() {
	port := flag.String("port", "8989", "port to listen")
	region := flag.String("region", "ap-northeast-1", "aws region")
	flag.Parse()

	e := echo.New()

	e.Logger.Fatal(e.Start(":" + *port))
}
