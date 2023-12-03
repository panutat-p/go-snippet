# Echo

https://echo.labstack.com

## Gracefully shutdown

```go
e := echo.New()
e.Use(RequestLogger)
e.GET("/hello", Hello)

go func() {
  err := e.Start(":8080")
  if err != nil && err != http.ErrServerClosed {
    e.Logger.Fatal("Shutting down the server")
  }
}()

quit := make(chan os.Signal, 1)
signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
<-quit
fmt.Println("🟡 Received shutdown signal")
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()
err := e.Shutdown(ctx)
if err != nil {
  e.Logger.Fatal(err)
}
```

## Middleware

https://echo.labstack.com/docs/middleware/logger#new-requestlogger-middleware

```go
func RequestLogger(next echo.HandlerFunc) echo.HandlerFunc {
  return func(c echo.Context) error {
    t1 := time.Now()
    err := next(c)
    latency := time.Since(t1)
    fmt.Printf("%s %s %v \n", c.Request().Method, c.Request().URL.Path, latency)
    return err
  }
}
```

## Reverse Proxy

```go
func NewWebProxy(e *echo.Echo) {
  url1, err := url.Parse("http://localhost:4000")
	if err != nil {
		fmt.Println("🔴 Failed to parse URL")
		panic(err)
	}
	url2, err := url.Parse("http://localhost:4001")
	if err != nil {
		fmt.Println("🔴 Failed to parse URL")
		panic(err)
	}
	targets := []*middleware.ProxyTarget{
		{URL: url1},
		{URL: url2},
	}
	conf := middleware.ProxyConfig{
		Balancer: middleware.NewRoundRobinBalancer(targets),
		Rewrite: map[string]string{
			"/web":   "/index",
			"/web/*": "/index",
		},
	}
	api := e.Group("/web")
	api.Use(middleware.ProxyWithConfig(conf))
}
```
