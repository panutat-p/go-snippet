# Atomic Counter

https://pkg.go.dev/sync/atomic

* For low-level applications, synchronization is better done with channels or the facilities of the `sync` package
* 32-bit platforms have some bugs, consider using the more ergonomic and less error-prone:
  * Use `Int64.Add` instead of `AddInt64`
  * Use `UInt64.Add` instead of `AddUInt64`
  * ...

```go
import "sync/atomic"
```

## Uint64

```go
count := atomic.Uint64{}

go func() {
  for {
    c := count.Add(1)
    fmt.Println("🟢 c:", c)
    time.Sleep(100 * time.Millisecond)
  }
}()

go func() {
  for {
    c := count.Add(1)
    fmt.Println("🔵 c:", c)
    time.Sleep(100 * time.Millisecond)
  }
}()

go func() {
  for {
    c := count.Add(1)
    fmt.Println("🟠 c:", c)
    time.Sleep(100 * time.Millisecond)
  }
}()

time.Sleep(5*time.Second)
```
