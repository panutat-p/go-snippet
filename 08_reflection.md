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
