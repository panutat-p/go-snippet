package rate_limiting

import (
	"fmt"
	"time"
	
	"golang.org/x/time/rate"
)

// https://medium.com/@isuruvihan/rate-limiting-in-go-controlling-traffic-with-efficiency-6a5ef7444ef8
func TestRateLimit_fixed_window(t *testing.T) {
	limiter := rate.NewLimiter(rate.Limit(100), 1) // Allow 100 requests per second
	for i := 0; i < 200; i++ {
		if !limiter.Allow() {
			fmt.Println("Rate limit exceeded. Request rejected.")
			continue
		}
		// Process the request
		fmt.Println("Request processed successfully.")
		time.Sleep(time.Millisecond * 100) // Simulate request processing time
	}
}

// https://medium.com/@isuruvihan/rate-limiting-in-go-controlling-traffic-with-efficiency-6a5ef7444ef8
func TestRateLimit_token_bucket(t *testing.T) {
	limiter := rate.NewLimiter(rate.Limit(10), 5) // Allow 10 requests per second with a burst of 5
	for i := 0; i < 200; i++ {
		if !limiter.Allow() {
			fmt.Println("Rate limit exceeded. Request rejected.")
			continue
		}
		// Process the request
		fmt.Println("Request processed successfully.")
		time.Sleep(time.Millisecond * 100) // Simulate request processing time
	}
}
