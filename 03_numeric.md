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
