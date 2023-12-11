# errros

https://pkg.go.dev/errors

https://tip.golang.org/doc/go1.20#errors

```go
ErrRedis := errors.New("redis")
e1 := errors.New("network error")
err := errors.Join(ErrRedis, e1)
fmt.Println(err)

if errors.Is(err, ErrRedis) {
fmt.Println("This is Redis error")
}
```
