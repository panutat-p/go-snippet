# Random

## rand

`rand.Seed` is deprecated: As of Go 1.20 there is no reason to call Seed with a random value.

```go
rand.Seed(time.Now().UnixNano())
```

```go
n := rand.Intn(100)
fmt.Println(n)
```

```go
func RandString(n int) string {
    var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
    b := make([]rune, n)
    for i := range b {
        b[i] = letterRunes[rand.Intn(len(letterRunes))]
    }
    return string(b)
}
```
