package rate_limiting

import (
	"fmt"
	"time"
	
	"golang.org/x/time/rate"
)

// https://pkg.go.dev/golang.org/x/time/rate
func TestRateLimit_token_bucket(t *testing.T) {
	limiter := rate.NewLimiter(rate.Limit(100), 100)
	for i := 0; i < 200; i++ {
		if !limiter.Allow() {
			t.Error("Rate limit exceeded")
			continue
		}
		fmt.Print(".")
		time.Sleep(time.Millisecond * 100)
	}
}
