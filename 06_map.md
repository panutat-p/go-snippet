# map

## Declare

```go
var m = make(map[string]any)
```

```go
var m := map[string]any{}
```

## LinkedHashMap

https://pkg.go.dev/container/list

```go
import (
    "container/list"
    "fmt"
    "strings"
)

type LinkedHashMap struct {
    dict map[string]int
    list *list.List
}

func NewLinkedHashMap() *LinkedHashMap {
    return &LinkedHashMap{
        dict: make(map[string]int),
        list: list.New(),
    }
}

func (m *LinkedHashMap) Get(key string) int {
    v, ok := m.dict[key]
    if !ok {
        return -1
    }
    return v
}

func (m *LinkedHashMap) Put(key string, value int) {
    _, ok := m.dict[key]
    if !ok {
        m.list.PushBack(key)
    }
    m.dict[key] = value
}

func (m *LinkedHashMap) Remove(key string) {
    _, ok := m.dict[key]
    if !ok {
        return
    }
    delete(m.dict, key)
    for e := m.list.Front(); e != nil; e = e.Next() {
        if e.Value.(string) == key {
            m.list.Remove(e)
            break
        }
    }
}

func (m *LinkedHashMap) Keys() []string {
    keys := make([]string, 0, m.list.Len())
    for e := m.list.Front(); e != nil; e = e.Next() {
        key := e.Value.(string)
        keys = append(keys, key)
    }
    return keys
}

func (m *LinkedHashMap) Size() int {
    return m.list.Len()
}

func (m *LinkedHashMap) String() string {
    var builder strings.Builder
    builder.WriteString("{")
    for e := m.list.Front(); e != nil; e = e.Next() {
        key := e.Value.(string)
        value := m.dict[key]
        if e.Next() == nil {
            builder.WriteString(fmt.Sprintf("%s: %d", key, value))
        } else {
            builder.WriteString(fmt.Sprintf("%s: %d, ", key, value))
        }
    }
    builder.WriteString("}")
    return builder.String()
}
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
