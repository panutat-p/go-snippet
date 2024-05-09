# Cache

## go cache

https://github.com/patrickmn/go-cache

```go
import "github.com/patrickmn/go-cache"
```

```go
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
```

```go
c := cache.New(5*time.Minute, 10*time.Minute)

c.Set("apple", 15, cache.DefaultExpiration)
c.Set("banana", 8, cache.DefaultExpiration)
items := c.Items()
for k, v := range items {
    expiredAt := time.Unix(0, v.Expiration)
    ttl := time.Until(expiredAt)
    fmt.Println("🟢", k, v.Object, ttl)
}
```

```go
c := cache.New(5*time.Minute, 10*time.Minute)

c.Set("apple", 15, cache.DefaultExpiration)
v, err := c.IncrementInt("apple", 1)
if err != nil {
    panic(err)
}
fmt.Println("🟢 apple:", v) // 16
```

```go
c := cache.New(5*time.Minute, 10*time.Minute)

c.Set("apple", 15, cache.DefaultExpiration)
c.Delete("apple")
v, ok = c.Get("apple")
if ok {
    apple := v.(int)
    fmt.Println("🟢 apple:", apple)
} else {
    fmt.Println("🔴 apple not found")
}
```
