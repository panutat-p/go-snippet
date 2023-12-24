# Elasticsearch Client

https://github.com/elastic/go-elasticsearch

https://pkg.go.dev/github.com/elastic/go-elasticsearch/v8

https://pkg.go.dev/github.com/elastic/go-elasticsearch/v8/esapi

```shell
go get github.com/elastic/go-elasticsearch/v8@latest
```

```go
import "github.com/elastic/go-elasticsearch/v8"
```

## Blog

https://www.elastic.co/blog/the-go-client-for-elasticsearch-introduction

https://www.elastic.co/blog/the-go-client-for-elasticsearch-configuration-and-customization

https://www.elastic.co/blog/the-go-client-for-elasticsearch-working-with-data

## client

```go
var (
  c *elasticsearch.Client
)
```

```go
func Connect() *elasticsearch.Client {
  conf := elasticsearch.Config{
    Addresses: []string{
      "http://localhost:9200",
    },
    Username: "admin",
    Password: "1234",
  }
  c, err := elasticsearch.NewClient(conf)
  if err != nil {
    fmt.Println("🔴 Failed to NewClient")
    panic(err)
  }
  api := c.Ping
  res, err := api(
    api.WithContext(context.Background()),
  )
  if err != nil {
    fmt.Println("🔴 Failed to Ping, client error")
    panic(err)
  }
  if res.IsError() {
    fmt.Println("🔴 Failed to Ping, Elasticsearch error")
    panic(err)
  }
  return c
}
```

```go
func CreateIndex(ctx context.Context) error {
  api := c.Indices.Create
  res, err := api(
    "fruit",
    api.WithContext(context.Background()),
    api.WithTimeout(5*time.Second),
  )
  if err != nil {
    fmt.Println("🔴 Failed to Create, client error")
    return err
  }
  defer res.Body.Close()
  if res.IsError() {
    fmt.Println("🔴 Failed to Create, Elasticsearch error")
    return err
  }
  b, err := io.ReadAll(res.Body)
  if err != nil {
    return err
  }
  fmt.Println("🟢 Succeeded to Create", string(b))
  return nil
}
```

```go
func Insert(ctx context.Context, doc any) error {
  data, err := json.Marshal(doc)
  if err != nil {
    return err
  }
  api := c.Index
  res, err := api(
    "fruit",
    bytes.NewReader(data),
    api.WithContext(context.Background()),
    api.WithTimeout(5*time.Second),
  )
  if err != nil {
    fmt.Println("🔴 Failed to Index, client error")
    return err
  }
  defer res.Body.Close()
  if res.IsError() {
    fmt.Println("🔴 Failed to Index, Elasticsearch error")
    return err
  }
  b, err := io.ReadAll(res.Body)
  if err != nil {
    return err
  }
  fmt.Println("🟢 Succeeded to Index", string(b))
  return nil
}
```
