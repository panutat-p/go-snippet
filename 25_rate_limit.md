# Rate Limiting

https://pkg.go.dev/golang.org/x/time/rate

* Token bucket algorithm
* Bucket size `b` with initially full (Burst limit)
* Refilled at rate r tokens per second

```go
import "golang.org/x/time/rate"
```

```go
limiter := rate.NewLimiter(rate.Limit(10), 10) // rate 10, burst 10
for i := 0; i < 200; i++ {
  if !limiter.Allow() {
    fmt.Println("🔴")
    fmt.Println("token:", limiter.Tokens())
    time.Sleep(time.Millisecond * 100) // fill 1 token
    continue
  }
  fmt.Print(".")
  time.Sleep(time.Millisecond * 50) // consume 20 tokens per second
}
```
