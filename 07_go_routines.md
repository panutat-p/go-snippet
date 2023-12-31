# Go Routines

## errgroup

https://pkg.go.dev/golang.org/x/sync/errgroup

```shell
go get golang.org/x/sync/errgroup
```

### No stop

* Expected: `apple` `amazon` `reddit` will success
* Expected: `hello` will fail but `cloudflare` `example` `google` will success

```go
func main() {
    var wg = sync.WaitGroup{}
    wg.Add(2)

    go Run(
        context.WithValue(
            context.Background(), "id", "🦊"),
        &wg,
        []string{
            "https://apple.com",
            "https://reddit.com",
            "https://amazon.com",
        },
    )
    go Run(
        context.WithValue(context.Background(), "id", "🐵"),
        &wg,
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
    defer wg.Done()
    id := ctx.Value("id")
    fmt.Println(id, "started")
    var g errgroup.Group

    for _, url := range urls {
        url := url
        g.Go(func() error {
            return Fetch(ctx, url)
        })
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
    _, err := http.DefaultClient.Get(url)
    if err != nil {
        fmt.Println(id, "failed to GET", url)
        return err
    }
    fmt.Println(id, "succeeded to GET", url)
    return nil
}

```

### Stop creating new Go routines when error

* Behave like `Promise.all` in JavaScript
* Expected: `apple` `amazon` `reddit` will success
* Expected: `hello` will fail but `cloudflare` `example` `google` will success

```go

```

### Stop with context cancellation

* Behave like `AbortController` & `AbortSignal` in JavaScript
* It will cancel all other ongoing operations when one operation returns an error
* Expected: `apple` `amazon` `reddit` will success
* Expected: `hello` will fail then `cloudflare` `example` `google` will be stopped

```go
func main() {
    var wg = sync.WaitGroup{}
    wg.Add(2)

    go Run(
        context.WithValue(
            context.Background(), "id", "🦊"),
        &wg,
        []string{
            "https://apple.com",
            "https://reddit.com",
            "https://amazon.com",
        },
    )
    go Run(
        context.WithValue(context.Background(), "id", "🐵"),
        &wg,
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
    defer wg.Done()
    id := ctx.Value("id")
    fmt.Println(id, "started")
    g, ctx := errgroup.WithContext(ctx)

    for _, url := range urls {
        url := url
        g.Go(func() error {
            return Fetch(ctx, url)
        })
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
