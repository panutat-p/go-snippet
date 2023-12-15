# Elasticsearch Client

https://github.com/elastic/go-elasticsearch

```shell
go get github.com/elastic/go-elasticsearch/v8@latest
```

## Connect

https://www.elastic.co/guide/en/elasticsearch/client/go-api/current/connecting.html

* Starting from version 8.0, Elasticsearch offers security by default with authentication and TLS enabled to use HTTPS
* If your cluster is configured with security explicitly disabled then you can connect via HTTP

```go
conf := elasticsearch.Config{
  Addresses: []string{
    "http://localhost:9200",
  },
  Username: "elastic",
  Password: "password",
}
c, err := elasticsearch.NewTypedClient(conf)
if err != nil {
  panic(err)
}
```

## Index

https://www.elastic.co/guide/en/elasticsearch/client/go-api/current/getting-started-go.html#_operations

```go
type Doc struct {
  Name string `json:"name"`
}
```

```go
doc := Doc{Name: "sample"}
c.Index("my_index").
  Id("1").
  Request(doc).
  Do(context.Background())
```
