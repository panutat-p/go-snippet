# Cache

## go cache

https://github.com/patrickmn/go-cache

```go
import (
     "github.com/patrickmn/go-cache"
)
```

```go
c := cache.New(5*time.Minute, 10*time.Minute)

c.Set("apple", 15, cache.DefaultExpiration)
c.Set("banana", 8, cache.NoExpiration)

v, ok := c.Get("apple")
if ok {
    apple := v.(string)
    fmt.Println(apple)
}
```
