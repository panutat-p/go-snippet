# Custom JSON Marshaller

```go
package main

import (
	"encoding/json"
	"fmt"
)

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

func main() {
	person1 := Person{Name: "John", Age:23, Address: "123 Main St"}
	person2 := Person{Name: "Alice", Age: 35, Address: ""}

	jsonStr1, _ := json.Marshal(person1)
	jsonStr2, _ := json.Marshal(person2)
	
	fmt.Println("person1:", string(jsonStr1))
	fmt.Println("person2:", string(jsonStr2))
}
```

```go
package main

import (
	"encoding/json"
	"fmt"
)

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

func main() {
	person1 := Person{Name: "John", Age: 23, Address: "123 Main St"}
	person2 := Person{Name: "Alice", Age: 35, Address: ""}

	jsonStr1, _ := json.Marshal(person1)
	jsonStr2, _ := json.Marshal(person2)

	fmt.Println("person1:", string(jsonStr1))
	fmt.Println("person2:", string(jsonStr2))
}
```
