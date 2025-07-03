package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// routing
	e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello from Lambda + Echo!")
	})

	// Echo to Lambda
	adapter := echoadapter.New(e)
	lambda.Start(adapter.ProxyWithContext)
}