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
r1 := new(big.Rat).SetFloat64(0.1)
r2 := new(big.Rat).SetFloat64(0.2)
got := new(big.Rat).Mul(r1, r2)
fmt.Println(got.FloatString(2)) // 0.02
```

```go
r1 := new(big.Rat).SetFloat64(0.7)
r2 := new(big.Rat).SetFloat64(0.1)
got := new(big.Rat).Quo(r1, r2)
fmt.Println(got.FloatString(2)) // 7.00
```

## Big rational + currentcy rounding

```go
r := new(big.Rat).SetFloat64(13.514)
fmt.Println(r.FloatString(3))
f, _ := r.Float64()
fmt.Printf("%.2f\n", f) // 13.51
```

```go
r := new(big.Rat).SetFloat64(13.515)
fmt.Println(r.FloatString(3))
f, _ := r.Float64()
fmt.Printf("%.2f\n", f) // 13.52
```
