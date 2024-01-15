# Generics

## `comparable`

* booleans, numbers, strings, pointers, channels, arrays of comparable types

```go
type Map[K comparable, V comparable] map[K]V
m := make(Map[string, int])
m["a"] += 1
m["a"] += 1
m["a"] += 1
m["b"] += 1
m["c"] += 1
fmt.Println(m)
```

## `constraints`

* Integer | Float | ~string

```go
import (
  "strconv"
  
  "golang.org/x/exp/constraints"
)

func Min[T constraints.Ordered](a, b T) T {
  if a < b {
    return a
  }
  return b
}

func Max[T constraints.Ordered](a, b T) T {
  if a > b {
    return a
  }
  return b
}
```
