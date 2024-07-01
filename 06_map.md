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

## Ordered map

```go
import (
    "fmt"
    "strings"
)

type OrderedMap struct {
    dict map[string]int
    list []string
}

func NewOrderedMap() *OrderedMap {
    return &OrderedMap{
        dict: make(map[string]int),
        list: make([]string, 0),
    }
}

func (m *OrderedMap) Has(key string) bool {
    _, ok := m.dict[key]
    return ok
}

func (m *OrderedMap) Get(key string) int {
    v, ok := m.dict[key]
    if !ok {
        return 0
    }
    return v
}

// Set ignore if key does not exist
func (m *OrderedMap) Set(key string, value int) {
    _, ok := m.dict[key]
    if !ok {
        return
    }
    m.dict[key] = value
}

// Put append key to list, and put key-value to map
// remove existing key if already exists
func (m *OrderedMap) Put(key string, value int) {
    m.Remove(key)
    m.list = append(m.list, key)
    m.dict[key] = value
}

func (m *OrderedMap) Remove(key string) {
    _, ok := m.dict[key]
    if !ok {
        return
    }
    delete(m.dict, key)
    for i, e := range m.list {
        if e == key {
            m.list = append(m.list[:i], m.list[i+1:]...)
            break
        }
    }
}

func (m *OrderedMap) Keys() []string {
    return m.list
}

func (m *OrderedMap) Size() int {
    return len(m.list)
}

func (m *OrderedMap) String() string {
    var sb strings.Builder
    sb.WriteString("{")
    for i, key := range m.list {
        if i == len(m.list)-1 {
            sb.WriteString(fmt.Sprintf("%s: %d", key, m.dict[key]))
        } else {
            sb.WriteString(fmt.Sprintf("%s: %d, ", key, m.dict[key]))
        }
    }
    sb.WriteString("}")
    return sb.String()
}

func (m *OrderedMap) GoString() string {
    return m.String()
}
```
