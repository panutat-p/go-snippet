# Random

## rand

`rand.Seed` is deprecated: As of Go 1.20 there is no reason to call Seed with a random value.

```go
rand.Seed(time.Now().UnixNano())
```

Generate [0,n)
```go
n := rand.Intn(100)
fmt.Println(n)
```

Generate string of n digits
```go
func RandNumber(n int) string {
    low := int(math.Pow10(n - 1))
    high := int(math.Pow10(n)) - 1
    return fmt.Sprintf("%d", low+rand.Intn(high-low+1))
}
```

Generate string of n characters
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

Generate decimal of XX.YY
```go
import "github.com/shopspring/decimal"

func RandDecimal(n int, p int32) decimal.Decimal {
    low := decimal.NewFromInt(int64(math.Pow10(n - 1)))
    high := decimal.NewFromInt(int64(math.Pow10(n)) - 1)
    diff := high.Sub(low)
    randDecimal := decimal.NewFromFloat(rand.Float64()).Mul(diff).Add(low)
    return randDecimal.Round(p)
}
```
