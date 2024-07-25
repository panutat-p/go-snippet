# httptest

## Mock Server

```go
func NewMockServer() *httptest.Server {
    handler := http.NewServeMux()
    handler.HandleFunc("/fruits", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(`["Apple", "Banana", "Cherry"]`))
    })
    return httptest.NewServer(handler)
}
```

```go
type Fruit struct {
    Name  string `json:"name"`
    Price int    `json:"price"`
}

func NewMockServer2() *httptest.Server {
    handler := http.NewServeMux()
    handler.HandleFunc("/fruits", func(w http.ResponseWriter, r *http.Request) {
        fruits := []Fruit{
            {Name: "Apple", Price: 30},
            {Name: "Banana", Price: 12},
            {Name: "Cherry", Price: 5},
        }
        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(fruits)
    })
    return httptest.NewServer(handler)
}
```
