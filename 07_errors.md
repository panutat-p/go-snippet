# errros

https://pkg.go.dev/errors

https://tip.golang.org/doc/go1.20#errors

## Return wrapped error using `errors.Join`

```go
var (
  ErrHTTP = errors.New("http")
)

func GetExample() error {
    _, err := http.Get("https://example.invalid")
  if err != nil {
    return errors.Join(ErrHTTP, err)
  }
  return nil
}
```

## Inspect wrapped error using `errors.Is`

```go
func main() {
  err := GetExample()
  if err != nil {
    if errors.Is(err, ErrHTTP) {
      fmt.Println("🟡 Failed to make a request")
      fmt.Println(err)
    } else {
      fmt.Println("🔴 Unexpected error")
      panic(err)
    }
  }
}
```
