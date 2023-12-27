# error

## Panic & Recover

* `recover` is a built-in function that regains control of a panicking Go routine
* ⚠️ `recover` needs to be called in the same Go routine where the panic occurred
* `recover` is only useful inside deferred functions

```go
func main() {
    defer Recover()
    Run()
}

func Recover() {
    r := recover()
    if r != nil {
        fmt.Println("Panic recovered, r:", r)
    }
}

func Run() {
    err := errors.New("app")
    panic(err)
}
```

## errors

https://pkg.go.dev/errors

https://tip.golang.org/doc/go1.20#errors

* Wrapping also preserves the original error
  * which means `errors.Is` and `errors.As` continue to work
  * regardless of how many times an error has been wrapped
* ⚠️ `errors.Unwrap` does not work with `errors.Join`

```go
var (
    ErrHTTP = errors.New("http")
)

func main() {
    err := GetExample()
    if err != nil {
        if errors.Is(err, ErrHTTP) {
            fmt.Println("🟡 Failed to make a request")
            fmt.Println(err)
        }
    } else {
        fmt.Println("🔴 Unexpected error")
        panic(err)
    }
}

func GetExample() error {
    _, err := http.Get("https://example.invalid")
    if err != nil {
        return errors.Join(ErrHTTP, err)
    }
    return nil
}
```

## fmt.Errorf

https://earthly.dev/blog/golang-errors

* Use `fmt.Errorf` with `%w` to wrap an error
* `errors.Unwrap` does not change its argument
* `errors.Unwrap` returns the original error, ignoring string format

```go
err := errors.New("failed to get user")
wrappedErr := fmt.Errorf("http: %w", err)
fmt.Println(wrappedErr)
fmt.Println(errors.Unwrap(wrappedErr))
```

## `errors.As`

```go
var pathError *fs.PathError
_, err := os.Open("non-existing")
if errors.As(err, &pathError) {
    fmt.Println("err:", err)
    fmt.Println("pathError:", pathError)
    // err and pathError now have the same value
}
```
