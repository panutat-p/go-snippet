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
        start := time.Now()
        err := next(c)
        latency := time.Since(start)
		fmt.Printf("%s %s %v \n", c.Request().Method, c.Request().URL.Path, latency)
        return err
    }
}
```
