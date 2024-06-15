# reflect

https://pkg.go.dev/reflect

https://github.com/a8m/reflect-examples

https://lpar.ath0.com/2016/04/20/reflection-go-modifying-struct-values

## Simple reflection

```go
s := "hello"
val := reflect.ValueOf(s)
fmt.Println(val.Type()) // string
fmt.Println(val.Type()) // string
```

```go
type User struct{}
u := User{}
val := reflect.ValueOf(u)
fmt.Println(val.Type()) // main.User
fmt.Println(val.Kind()) // struct
```

```go
func PrintType(x any) {
    t := reflect.TypeOf(x)
    switch t.Kind() {
    case reflect.Int:
        fmt.Println("It's an integer!")
    case reflect.Float64:
        fmt.Println("It's a float!")
    case reflect.String:
        fmt.Println("It's a string!")
    default:
        fmt.Println("I don't recognize the type.")
    }
}
```

## Get tags of a struct

```go
func main() {
    t := reflect.TypeOf(Fruit{})
    for i := 0; i < t.NumField(); i++ {
        field := t.Field(i)
        tag := field.Tag.Get("json")
        fmt.Println(tag)
    }
}

type Fruit struct {
    Name string `json:"name"`
    Price int `json:"price"`
}
```

## Struct diff

```go
package main

import (
    "fmt"
    "reflect"
    "time"
)

func main() {
    fruits := []Fruit{
        {"apple", 15, "Hokkaido", time.Now()},
        {"apple", 15, "Yamanashi", time.Now()},
        {"apple red", 30, "Yamanashi", time.Now()},
    }

    for i := 0; i < len(fruits)-1; i++ {
        m := Diff(fruits[i], fruits[i+1])
        if len(m) > 0 {
            fmt.Printf("👉 Diff of fruit %d & fruit %d:\n", i, i+1)
            for field, value := range m {
                fmt.Printf("   %s: %v\n", field, value)
            }
            fmt.Println()
        }
    }
}

type Fruit struct {
    Name        string
    Price       int
    Factory     string
    Produced_at time.Time `diff:"ignore"`
}

func Diff(a, b any) map[string]any {
    if a == nil || b == nil {
        return nil
    }

    va := reflect.ValueOf(a)
    vb := reflect.ValueOf(b)

    ta := va.Type()

    diff := make(map[string]interface{})
    for i := 0; i < va.NumField(); i++ {
        if ta.Field(i).Tag.Get("diff") == "ignore" {
            continue
        }

        if va.Field(i).Interface() != vb.Field(i).Interface() {
            diff[ta.Field(i).Name] = vb.Field(i).Interface()
        }
    }

    return diff
}
```

## Recursive reflection

```go
func PrintObject(x any) {
    val := reflect.ValueOf(x)
    switch val.Kind() {
    case reflect.Ptr:
        val = val.Elem()
        fmt.Printf("%s: %p -> %v\n", val.Type(), x, val)
        if val.Kind() == reflect.Struct {
            iterateFields(val)
        }
    case reflect.Struct:
        iterateFields(val)
    default:
        fmt.Printf("%s: %v\n", val.Type(), val)
    }
}

func iterateFields(val reflect.Value) {
    for i := 0; i < val.NumField(); i++ {
        f := val.Type().Field(i)
        if f.PkgPath != "" {
            fmt.Printf("%s: %v\n", f.Name, "***")
            continue
        }
        v := val.Field(i)
        fmt.Printf("%s: %v\n", f.Name, v)
        if v.Kind() == reflect.Ptr {
            v = v.Elem()
        }
        if v.Kind() == reflect.Struct {
            iterateFields(v)
            continue
        }
    }
}
```
