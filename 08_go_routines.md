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

### Stop but not cancel

* Behave like `Promise.all` in JavaScript but `errgroup.Group` waits for all goroutines to complete before it returns, even if one of them has returned an error.
* Expected: `apple` `amazon` `reddit` will success
* Expected: `hello` will fail then `cloudflare` `example` `google` will not be created

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

    ctx, cancel := context.WithCancel(ctx)
    defer cancel()
    var g errgroup.Group

    for _, url := range urls {
        time.Sleep(1 * time.Second) // remove cooperative behavior
        url := url
        select {
        case <-ctx.Done():
            fmt.Println(id, "⚠️ context cancelled")
            return
        default:
            g.Go(func() error {
                err := Fetch(ctx, url)
                if err != nil {
                    cancel()
                    return err
                }
                return nil
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
    _, err := http.DefaultClient.Get(url)
    if err != nil {
        fmt.Println(id, "failed to GET", url)
        return err
    }
    fmt.Println(id, "succeeded to GET", url)
    return nil
}
```

### Stop with context cancellation

* Behave like `AbortController` & `AbortSignal` in JavaScript
* When a routine return an error, `errgroup` will cancel the cantext
* The running routines may exit or not, depending on whether they respect the context or not
* For example, `io` does not respect the context, but `net/http` does

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
