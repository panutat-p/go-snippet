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

### Stop with context cancellation

* Behave like `AbortController` & `AbortSignal` in JavaScript
* It will cancel all other ongoing operations when one operation returns an error

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
                return Fetch(ctx, url)
            })
        }
    }

    err := g.Wait()
    if err != nil {
        panic(err)
    }
    fmt.Println("✅ Done")
}

func Fetch(ctx context.Context, url string) error {
    req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
    if err != nil {
        fmt.Println("🔴 Failed to create a HTTP request for", url)
        return err
    }
    _, err = http.DefaultClient.Do(req)
    if err != nil {
        fmt.Println("🔴 Failed to GET", url)
        return err
    }
    fmt.Println("🟢 Succeeded to GET", url)
    return nil
}
```

## waitgroup + errgroup

* Expected: `apple` `amazon` `reddit` are success
* Expected: `hello` will fail then `cloudflare` `example` `googlre` will be stopped

```go
func main() {
    var wg = &sync.WaitGroup{}
    wg.Add(2)

    go Run(
        context.WithValue(
            context.Background(), "id", "🦊"),
        wg,
        []string{
            "https://apple.com",
            "https://reddit.com",
            "https://amazon.com",
        },
    )
    go Run(
        context.WithValue(context.Background(), "id", "🐵"),
        wg,
        []string{
            "hello",
            "https://example.com",
            "https://google.com",
            "https://cloudflare.com",
        },
    )

    wg.Wait()
}

func Run(ctx context.Context, wg *sync.WaitGroup, urls []string) {
    id := ctx.Value("id")
    fmt.Println(id, "started")
    defer wg.Done()
    g, ctx := errgroup.WithContext(ctx)

    for _, url := range urls {
        url := url
        select {
        case <-ctx.Done():
            return
        default:
            g.Go(func() error {
                return Fetch(ctx, url)
            })
        }
    }

    err := g.Wait()
    if err != nil {
        fmt.Println(id, "❌", err)
        return
    }
    fmt.Println(id, "✅")
}

func Fetch(ctx context.Context, url string) error {
    id := ctx.Value("id")
    req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
    if err != nil {
        fmt.Println(id, "failed to create request for", url)
        return err
    }
    _, err = http.DefaultClient.Do(req)
    if err != nil {
        fmt.Println(id, "failed to GET", url)
        return err
    }
    fmt.Println(id, "succeeded to GET", url)
    return nil
}
```
