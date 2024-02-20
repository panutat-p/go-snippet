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
func RandNumber(n int) string {
    rand.Seed(time.Now().UnixNano())
    low := int(math.Pow10(n - 1))
    high := int(math.Pow10(n)) - 1
    return fmt.Sprintf("%d", low+rand.Intn(high-low+1))
}
```

```go
func RandAlphabet(n int) string {
    var alphabet = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
    b := make([]rune, n)
    for i := range b {
        b[i] = alphabet[rand.Intn(len(alphabet))]
    }
    return string(b)
}
```
