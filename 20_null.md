# null

https://github.com/guregu/null

## string

```go
payload := `{"name": "apple"}`
type Fruit struct {
	Name    string    `json:"name"`
}
var f Fruit
err := json.Unmarshal([]byte(payload), &f)
if err != nil {
    panic(err)
}
fmt.Printf("👉%+v\n", f)
```
