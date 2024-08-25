# encoding/json

https://pkg.go.dev/encoding/json

## Print struct as JSON

```go
func PrintJSON(o any) {
    b, err := json.MarshalIndent(o, "", "  ")
    if err != nil {
        panic(err)
    }
    fmt.Println(string(b))
}
```

```go
func PrintJSON(o any) {
    kind := reflect.TypeOf(o).Kind()
    switch kind {
    case reflect.String:
        s := reflect.ValueOf(o).String()
        b := []byte(s)
        if len(b) == 0 {
            fmt.Println("❌  Cannot Unmarshal empty string to JSON")
            return
        }
        var v any
        err := json.Unmarshal([]byte(s), &v)
        if err != nil {
            fmt.Println("❌  Cannot Unmarshal string to JSON, err:", err)
            return
        }
        b, _ = json.MarshalIndent(v, "", "  ")
        fmt.Println(string(b))
    case reflect.Slice:
        if reflect.TypeOf(o).Elem().Kind() == reflect.Uint8 {
            b := reflect.ValueOf(o).Bytes()
            if len(b) == 0 {
                fmt.Println("❌  Cannot Unmarshal empty []byte to JSON")
                return
            }
            var v any
            err := json.Unmarshal(b, &v)
            if err != nil {
                fmt.Println("❌  Cannot Unmarshal []byte to JSON, err:", err)
                return
            }
            b, _ = json.MarshalIndent(v, "", "  ")
            fmt.Println(string(b))
        } else {
            b, err := json.MarshalIndent(o, "", "  ")
            if err != nil {
                fmt.Println("❌  Cannot Marshal to JSON, err:", err)
                return
            }
            fmt.Println(string(b))
        }
    default:
        b, err := json.MarshalIndent(o, "", "  ")
        if err != nil {
            fmt.Println("❌  Cannot Marshal object to JSON, err:", err)
            return
        }
        fmt.Println(string(b))
    }
}
```

## Marshal & Unmarshal

```go
b, err := json.Marshal(o)
if err != nil {
    panic(err)
}
fmt.Println(string(b))
```

```go
var m map[string]any
err := json.Unmarshal([]byte(`{"name": "apple", "price": 100}`), &m)
if err != nil {
    panic(err)
}
fmt.Printf("%+v\n", m)
```

## Custom marshaller: string to number

* `Fruit{ Name:"apple" Price: "100.50" }` to `{"name":"apple","price":100.50}`
* `{"name":"apple","price":100.50}` to `Fruit{ Name:"apple" Price:"100.50" }`

```go
import (
    "encoding/json"
    "fmt"
)

type Fruit struct {
    Name  string `json:"name"`
    Price string `json:"price"`
}

func (f Fruit) MarshalJSON() ([]byte, error) {
    type Alias Fruit
    
    return json.Marshal(&struct {
            Alias
            Price json.Number `json:"price"`
        }{
            Alias: (Alias)(f),
            Price: json.Number(f.Price),
    })
}

func (f *Fruit) UnmarshalJSON(b []byte) error {
    type Alias Fruit
    aux := &struct {
        *Alias
        Price json.Number `json:"price"`
    }{
        Alias: (*Alias)(f),
    }
    if err := json.Unmarshal(b, &aux); err != nil {
        return err
    }
    f.Price = aux.Price.String()
    return nil
}
```

### Alias

```go
type Person struct {
  Name    string `json:"name"`
  Age     int    `json:"age"`
  Address string `json:"address"`
}

func (p Person) MarshalJSON() ([]byte, error) {
  type Alias Person

  if p.Address == "" {
    return json.Marshal(&struct {
      Alias
      Address *string `json:"address"`
    }{
      Alias:   (Alias)(p),
      Address: nil,
    })
  }

  return json.Marshal(&struct {
    Alias
  }{
    Alias: (Alias)(p),
  })
}
```

```go
func main() {
  person1 := Person{Name: "John", Age:23, Address: "123 Main St"}
  person2 := Person{Name: "Alice", Age: 35, Address: ""}
  
  jsonStr1, _ := json.Marshal(person1)
  jsonStr2, _ := json.Marshal(person2)
  
  fmt.Println("person1:", string(jsonStr1))
  fmt.Println("person2:", string(jsonStr2))
}
```

### Without alias

```go
type Person struct {
  Name    string
  Age     int
  Address string
}

func (p Person) MarshalJSON() ([]byte, error) {
  if p.Address == "" {
    return json.Marshal(&struct {
      Name    string   `json:"name"`
      Age     int      `json:"age"`
      Address *string  `json:"address"`
    }{
      Name:    p.Name,
      Age:     p.Age,
      Address: nil,
    })
  }

  return json.Marshal(&struct {
    Name    string `json:"name"`
    Age     int    `json:"age"`
    Address string `json:"address"`
  }{
    Name:    p.Name,
    Age:     p.Age,
    Address: p.Address,
  })
}
```

```go
func main() {
  person1 := Person{Name: "John", Age: 23, Address: "123 Main St"}
  person2 := Person{Name: "Alice", Age: 35, Address: ""}
  
  jsonStr1, _ := json.Marshal(person1)
  jsonStr2, _ := json.Marshal(person2)
  
  fmt.Println("person1:", string(jsonStr1))
  fmt.Println("person2:", string(jsonStr2))
}
```
