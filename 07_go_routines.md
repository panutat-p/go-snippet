# Go Routines

## errgroup

https://pkg.go.dev/golang.org/x/sync/errgroup

```shell
go get golang.org/x/sync/errgroup
```

### No stop

```go
func main() {
    urls := []string{
        "https://example.com",
        "https://google.com",
        "https://cloudflare.com",
        "_",
        "hello",
        "http://not-exist.com",
        "https://not-exist.com",
    }

    g := &errgroup.Group{}

    for _, url := range urls {
        url := url
        g.Go(func() error {
            return Fetch(url)
        })
    }

    err := g.Wait()
    if err != nil {
        panic(err)
    }
    fmt.Println("✅ Done")
}

func Fetch(url string) error {
    _, err := http.Get(url)
    if err != nil {
        fmt.Println("🔴 Failed to GET", url)
        return err
    }
    fmt.Println("🟢 Succeeded to GET", url)
    return nil
}
```

### Stop when error

* Behave like `Promise.all` in JavaScript

```go
func main() {
    urls := []string{
        "https://example.com",
        "https://google.com",
        "https://cloudflare.com",
        "_",
        "hello",
        "http://not-exist.com",
        "https://not-exist.com",
    }

    g, ctx := errgroup.WithContext(context.Background())

    for _, url := range urls {
        url := url
        select {
        case <-ctx.Done():
            return
        default:
            g.Go(func() error {
                return Fetch(url)
            })
        }
    }

    err := g.Wait()
    if err != nil {
        panic(err)
    }
    fmt.Println("✅ Done")
}

func Fetch(url string) error {
    _, err := http.Get(url)
    if err != nil {
        fmt.Println("🔴 Failed to GET", url)
        return err
    }
    fmt.Println("🟢 Succeeded to GET", url)
    return nil
}
```

## Stop and cancel when error

* Behave like `AbortController` & `AbortSignal` in JavaScript

```go

```
