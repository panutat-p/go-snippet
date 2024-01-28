# Numeric

## Float

```go
f1 := 0.1
f2 := 0.2
got := f1 * f2
fmt.Println(got) // 0.020000000000000004
```

```go
f1 := 0.7
f2 := 0.1
got := f1 / f2
fmt.Println(got) // 6.999999999999999
```

## Big rational

```go
import "math/big"
```

```go
r1, _ := new(big.Rat).SetString("0.1")
r2, _ := new(big.Rat).SetString("0.2")
r3 := new(big.Rat).Mul(r1, r2)
fmt.Println(r3.FloatString(2)) // 0.02
```

```go
r1, _ := new(big.Rat).SetString("0.7")
r2, _ := new(big.Rat).SetString("0.1")
r3 := new(big.Rat).Mul(r1, r2)
fmt.Println(r3.FloatString(2)) // 7.00
```
## Big rational + standard rounding

```go
r, _ := new(big.Rat).SetString("13.514")
f, _ := r.Float64()
fmt.Printf("%.2f\n", f) // 13.51
```

```go
r, _ := new(big.Rat).SetString("13.515")
f, _ := r.Float64()
fmt.Printf("%.2f\n", f) // 13.52
```
