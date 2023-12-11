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

## Error wrapping

https://earthly.dev/blog/golang-errors

* Error wrapping can provide additional context about the lineage of an error, in ways similar to a traditional stack-trace
* `fmt.Errorf` with a `%w`
* Wrapping also preserves the original error, which means `errors.Is` and `errors.As` continue to work, regardless of how many times an error has been wrapped
* Call `errors.Unwrap` to return the previous error in the chain
