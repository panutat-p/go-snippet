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

## Marshal to null (override)

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

## Marshal to null (declare tags in method)

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
