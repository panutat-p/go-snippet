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

type ElasticMapping struct {
	Mapping struct {
		Properties map[string]struct {
			Type   string `json:"type"`
			Fields any    `json:"fields"`
		} `json:"properties"`
	} `json:"mappings"`
}

func IndexMapping(ctx context.Context, index string) (*ElasticMapping, error) {
    api := c.Indices.GetMapping
    res, err := api(
        api.WithContext(ctx),
        api.WithErrorTrace(),
        api.WithIndex(index),
    )
    if err != nil {
        return nil, err
    }
    defer res.Body.Close()
    if res.IsError() {
        return nil, err
    }
    b, err := io.ReadAll(res.Body)
    if err != nil {
        return nil, err
    }
    var indexMapping map[string]ElasticMapping
	err = json.Unmarshal(b, &indexMapping)
	if err != nil {
        err = errors.New("json")
        err = errors.Join(ErrElasticsearchMerchant, err)
        return nil, err
	}
    mapping := indexMapping[index]
    return &mapping, nil
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

## Bulk insert

```go
import (
    "github.com/dustin/go-humanize"
    "github.com/elastic/go-elasticsearch/v8"
    "github.com/elastic/go-elasticsearch/v8/esapi"
)
```

```go
var (
    count = 1000
    batch = 250

    buf bytes.Buffer

    indexName = "articles"

    numItems   int
    numErrors  int
    numIndexed int
    numBatches int
    currBatch  int
)

func main() {
    log.SetFlags(0)

    log.Printf("\x1b[1mBulk\x1b[0m: documents [%s] batch size [%s]", humanize.Comma(int64(count)), humanize.Comma(int64(batch)))
    log.Println(strings.Repeat("▁", 65))

    es := ConnectElasticsearch()

    docs := GenerateDocs()

    fmt.Print("🟢 Sending batch ")

    if count%batch == 0 {
        numBatches = (count / batch)
    } else {
        numBatches = (count / batch) + 1
    }

    start := time.Now().UTC()

    for i, doc := range docs {
        numItems++
        currBatch = i / batch
        if i == count-1 {
            currBatch++
        }

        meta := []byte(fmt.Sprintf(`{ "index" : { "_id" : "%d" } }%s`, doc.ID, "\n"))
        data, _ := json.Marshal(doc)
        data = append(data, "\n"...)
        buf.Grow(len(meta) + len(data))
        buf.Write(meta)
        buf.Write(data)

        if i > 0 && i%batch == 0 || i == count-1 {
            fmt.Printf("[%d/%d] ", currBatch, numBatches)
            var res *esapi.Response
            res, _ = es.Bulk(bytes.NewReader(buf.Bytes()), es.Bulk.WithIndex(indexName))
            switch res.IsError() {
            case true:
                numErrors += numItems
                var raw map[string]any
                _ = json.NewDecoder(res.Body).Decode(&raw)
                log.Printf("🔴 [%d] %s: %s", res.StatusCode, raw["error"].(map[string]any)["type"], raw["error"].(map[string]any)["reason"])
            case false:
                var blk *bulkResponse
                _ = json.NewDecoder(res.Body).Decode(&blk)
                for _, d := range blk.Items {
                    if d.Index.Status > 201 {
                        numErrors++ // status_code > 200 is error
                        log.Printf("🔴 [%d]: %s: %s: %s: %s", d.Index.Status, d.Index.Error.Type, d.Index.Error.Reason, d.Index.Error.Cause.Type, d.Index.Error.Cause.Reason)
                    } else {
                        numIndexed++
                    }
                }
            }
            res.Body.Close()
            // Reset the buffer and items counter
            buf.Reset()
            numItems = 0
        }
    }

    // Report the results: number of indexed docs, number of errors, duration, indexing rate
    fmt.Print("\n")
    log.Println(strings.Repeat("▔", 65))

    dur := time.Since(start)

    if numErrors > 0 {
        log.Fatalf("Indexed [%s] documents with [%s] errors in %s (%s docs/sec)", humanize.Comma(int64(numIndexed)), humanize.Comma(int64(numErrors)), dur.Truncate(time.Millisecond), humanize.Comma(int64(1000.0/float64(dur/time.Millisecond)*float64(numIndexed))))
    } else {
        log.Printf("Sucessfuly indexed [%s] documents in %s (%s docs/sec)", humanize.Comma(int64(numIndexed)), dur.Truncate(time.Millisecond), humanize.Comma(int64(1000.0/float64(dur/time.Millisecond)*float64(numIndexed))))
    }
}

func GenerateDocs() []*Article {
    var articles []*Article
    names := []string{"Alice", "John", "Mary"}
    for i := 1; i < count+1; i++ {
        articles = append(articles, &Article{
            ID:        i,
            Title:     strings.Join([]string{"Title", strconv.Itoa(i)}, " "),
            Body:      "Lorem ipsum dolor sit amet...",
            Published: time.Now().Round(time.Second).UTC().AddDate(0, 0, i),
            Author: Author{
                FirstName: names[rand.Intn(len(names))],
                LastName:  "Smith",
            },
        })
    }
    log.Printf("→ Generated %s articles", humanize.Comma(int64(len(articles))))
    return articles
}

func ConnectElasticsearch() *elasticsearch.Client {
    es, err := elasticsearch.NewDefaultClient()
    if err != nil {
        log.Fatalf("🔴 Failed to NewDefaultClient: %s", err)
    }
    return es
}

type Article struct {
    ID        int       `json:"id"`
    Title     string    `json:"title"`
    Body      string    `json:"body"`
    Published time.Time `json:"published"`
    Author    Author    `json:"author"`
}

type Author struct {
    FirstName string `json:"first_name"`
    LastName  string `json:"last_name"`
}

type bulkResponse struct {
    Errors bool `json:"errors"`
    Items  []struct {
        Index struct {
            ID     string `json:"_id"`
            Result string `json:"result"`
            Status int    `json:"status"`
            Error  struct {
                Type   string `json:"type"`
                Reason string `json:"reason"`
                Cause  struct {
                    Type   string `json:"type"`
                    Reason string `json:"reason"`
                } `json:"caused_by"`
            } `json:"error"`
        } `json:"index"`
    } `json:"items"`
}
```
