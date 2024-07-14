# net/http

https://pkg.go.dev/net/http

```go
res, err := http.Get("http://example.com")
if err != nil {
    panic(err)
}
defer res.Body.Close()
b, err := io.ReadAll(res.Body)
if err != nil {
    panic(err)
}
fmt.Println(string(b))
```

```go
req, err := http.NewRequest("GET", "http://example.com", nil)
req.Header.Add("User-Agent", "demo")
res, err := client.Do(req)
if err != nil {
    panic(err)
}
defer res.Body.Close()
b, err := io.ReadAll(res.Body)
if err != nil {
    panic(err)
}
fmt.Println(string(b))
```

```go
reqBody := []byte(`{"key": "value"}`)
req, err := http.NewRequest("POST", "http://example.com", bytes.NewBuffer(requestBody))
req.Header.Set("Content-Type", "application/json")
res, err := client.Do(req)
if err != nil {
    panic(err)
}
defer res.Body.Close()
b, err := io.ReadAll(res.Body)
if err != nil {
    panic(err)
}
fmt.Println(string(b))
```
