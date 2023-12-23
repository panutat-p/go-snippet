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

## Marshal

```go
type Person struct {
  Name  string `json:"name"`
  Age   int    `json:"age"`
  City  string `json:"city"`
  Email string `json:"email"`
}

func main() {
  person := Person{
    Name:  "John Doe",
    Age:   30,
    City:  "New York",
    Email: "john@gmail.com",
  }

  b, err := json.Marshal(person)
  if err != nil {
    panic(err)
  }
  fmt.Println(string(b))
}
```

```go
func PrintStruct(o any) {
  b, err := json.MarshalIndent(o, "", "  ")
  if err != nil {
    panic(err)
  }
  fmt.Println(string(b))
}
```

## Unmarshal

```go
type Person struct {
  Name  string `json:"name"`
  Age   int    `json:"age"`
  City  string `json:"city"`
  Email string `json:"email"`
}

func main() {
  s := `{"name":"John Doe","age":30,"city":"New York","email":"john@gmail.com"}`
  b := []byte(s)

  var person Person
  err := json.Unmarshal(b, &person)
  if err != nil {
    panic(err)
  }

  fmt.Printf("%+v\n", person)
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
