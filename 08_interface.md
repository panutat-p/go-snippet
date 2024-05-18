# Interface

https://go.dev/tour/methods/9

## PGO

https://tip.golang.org/doc/pgo

Go 1.22, interface method calls are better optimized.

## Stringer

```go
type Fruit struct {
    Name string
    Price int
}

func (f Fruit) String() string{
    return "Fruit.String()"
}

func (f Fruit) GoString() string{
    return "Fruit.GoString()"
}
```
