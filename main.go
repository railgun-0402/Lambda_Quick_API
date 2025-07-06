package main

func main() {
	// e := echo.New()

	// // routing
	// e.GET("/", func(c echo.Context) error {
    // 	return c.String(http.StatusOK, "Hello from Lambda root!")
	// })
	// e.GET("/hello", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello from Lambda + Echo!")
	// })

	// // Echo to Lambda
	// adapter := echoadapter.New(e)
	// lambda.Start(adapter.ProxyWithContext)
	PutItem()
}