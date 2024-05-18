# Go Routines

## Semaphore

https://levelup.gitconnected.com/go-concurrency-pattern-semaphore-9587d45f058d

```go
var (
    c atomic.Uint32
)

type Fruit struct {
    ID    uint32
    Name  string
    Price int
}

func main() {
    var wg sync.WaitGroup
    fruits := GenerateFruits(100)
    semaphore := make(chan struct{}, 10)
    wg.Add(len(fruits))
    for _, f := range fruits {
        f := f
        go func() {
            semaphore <- struct{}{}
            defer wg.Done()
            defer func() { <-semaphore }()
            PrintFruit(f)
        }()
    }
    wg.Wait()
}

func GenerateFruits(count int) []Fruit {
    var fruits []Fruit
    for i := 0; i < count; i++ {
        c.Add(1)
        f := Fruit{
            ID:    c.Load(),
            Name:  RandStringRunes(6),
            Price: rand.Intn(100),
        }
        fruits = append(fruits, f)
    }
    return fruits
}

func RandStringRunes(n int) string {
    var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
    b := make([]rune, n)
    for i := range b {
        b[i] = letterRunes[rand.Intn(len(letterRunes))]
    }
    return string(b)
}

func PrintFruit(f Fruit) {
    time.Sleep(1 * time.Second)
    fmt.Printf("%+v\n", f)
}
```

## errgroup

https://pkg.go.dev/golang.org/x/sync/errgroup

```shell
go get golang.org/x/sync/errgroup
```

### Limits the number of active goroutines in the group

```go
var c atomic.Uint32

func main() {
  var group errgroup.Group
  group.SetLimit(2)
  fruits := GenerateFruits(100)
  for _, f := range fruits {
    f := f
    group.Go(func() error {
      PrintFruit(f)
      return nil
    })
  }
  err := group.Wait()
  if err != nil {
    fmt.Println("🔴 err:", err)
  }
}

func GenerateFruits(count int) []Fruit {
  var fruits []Fruit
  for i := 0; i < count; i++ {
    c.Add(1)
    f := Fruit{
      ID:    c.Load(),
      Name:  RandStringRunes(6),
      Price: rand.Intn(100),
    }
    fruits = append(fruits, f)
  }
  return fruits
}

func PrintFruit(f Fruit) {
  time.Sleep(1 * time.Second)
  fmt.Printf("%+v\n", f)
}
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
