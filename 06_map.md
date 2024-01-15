# map

## Declare

```go
var m = make(map[string]any)
```

```go
var m = make(map[string]any, 5)
```

```go
var m := map[string]any{}
```

## Use slice as map key

```go
import (
    "fmt"
    "math/rand"
    "strconv"
    "strings"
)

func main() {
    m := make(map[string]int)
    k1 := Key{0}.String()
    m[k1] = rand.Intn(100)
    k2 := Key{0, 1}.String()
    m[k2] = rand.Intn(100)
    k3 := Key{0, 1, 2}.String()
    m[k3] = rand.Intn(100)
    fmt.Println(m)
}

type Key []int

func (k Key) String() string {
    strs := make([]string, len(k))
    for i, v := range k {
        strs[i] = strconv.Itoa(v)
    }
    return strings.Join(strs, ",")
}
```
