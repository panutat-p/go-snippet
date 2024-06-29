# Echo

https://echo.labstack.com

```shell
go get github.com/labstack/echo/v4
```

```go
import (
  echomiddleware "github.com/labstack/echo/v4/middleware"
  "github.com/labstack/echo/v4"
)
```

## Gracefully shutdown

```go
e := echo.New()
e.Use(echomiddleware.Recover())
e.GET("/", func(c echo.Context) error {
  return c.String(http.StatusOK, "Hello, Echo!")
})

go func() {
  if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
    e.Logger.Fatal(err)
  }
}()

var stop = make(chan os.Signal, 1)
signal.Notify(
  stop,
  os.Interrupt,
  syscall.SIGINT,
  syscall.SIGTERM,
)
<-stop
fmt.Println("🟡 Gracefully shutting down")
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()
if err := e.Shutdown(ctx); err != nil {
  fmt.Println("🔴 Failed to Shutdown Echo")
  fmt.Println(err)
}
```

## Middleware

### Request logger

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

### IP address

```go
func IpAddressExtractor(next echo.HandlerFunc) echo.HandlerFunc {
  return func(c echo.Context) error {
    ip := c.RealIP() // get IP address from network layer
    c.Set("client_ip", ip)
    return next(c)
  }
}
```

## Serve static

```go
e := echo.New()
e.File("/index", "public/index.html")
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
  g := e.Group("/web")
  g.Use(middleware.ProxyWithConfig(conf))
}
```
