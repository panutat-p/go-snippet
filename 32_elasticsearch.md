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
func NewElasticsearchClient() *elasticsearch.Client {
  conf := elasticsearch.Config{
    Addresses: []string{
      "http://localhost:9200",
    },
    Username: "admin",
    Password: "1234",
  }
  c, err := elasticsearch.NewClient(conf)
  if err != nil {
    panic(err)
  }
  api := c.Ping
  res, err := api(
    api.WithContext(context.Background()),
  )
  if err != nil {
    panic(err)
  }
  if res.IsError() {
    panic(err)
  }
  return c
}
```

```go
var c *elasticsearch.Client

func CreateIndex(ctx context.Context, index string) error {
  api := c.Indices.Create
  res, err := api(
    index,
    api.WithContext(context.Background()),
    api.WithTimeout(5*time.Second),
  )
  if err != nil {
    return err
  }
  defer res.Body.Close()
  if res.IsError() {
    return err
  }
  b, err := io.ReadAll(res.Body)
  if err != nil {
    return err
  }
  return nil
}
```

```go
var c *elasticsearch.Client

func Insert(ctx context.Context, index string, doc any) error {
  data, err := json.Marshal(doc)
  if err != nil {
    return err
  }
  api := c.Index
  res, err := api(
    index,
    bytes.NewReader(data),
    api.WithContext(context.Background()),
    api.WithTimeout(5*time.Second),
  )
  if err != nil {
    return err
  }
  defer res.Body.Close()
  if res.IsError() {
    return err
  }
  return nil
}
```
