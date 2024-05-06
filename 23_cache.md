# Cache

## go cache

https://github.com/patrickmn/go-cache

```go
import (
    "fmt"
    "time"

    "github.com/patrickmn/go-cache"
)

func main() {
    c := cache.New(5*time.Minute, 10*time.Minute)

    c.Set("apple", 15, cache.DefaultExpiration)

    v, ok := c.Get("apple")
    if ok {
        apple := v.(int)
        fmt.Println("🟢 apple:", apple)
    }

    v, exp, ok := c.GetWithExpiration("apple")
    if ok {
        apple := v.(int)
        ttl := time.Until(exp)
        fmt.Println("🟢 apple:", apple, "TTL:", ttl) // TTL ~ 4m59.999958973s
    }

    v, err := c.IncrementInt("apple", 1)
    if err != nil {
        panic(err)
    }
    fmt.Println("🟢 apple:", v) // 16
}
```
