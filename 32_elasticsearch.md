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
type Fruit struct {
  Name  string `json:"name"`
  Price int    `json:"price"`
}

func main() {
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

  c.Indices.Create("fruit")

  fruit := Fruit{
    Name:  "apple",
    Price: 15,
  }
  data, err := json.Marshal(fruit)
  if err != nil {
    panic(err)
  }
  c.Index("fruit", bytes.NewReader(data))
}
```

## esapi

```go

```
