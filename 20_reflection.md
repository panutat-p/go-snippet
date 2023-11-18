# Reflection

https://pkg.go.dev/reflect

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

## Recursive reflection

```go
func PrintObject(x any) {
	t := reflect.TypeOf(x)
	switch t.Kind() {
	case reflect.Ptr:
		val := reflect.ValueOf(x).Elem()
		fmt.Printf("%s: %p -> %v\n", t, x, val)
		if val.Kind() == reflect.Struct {
			iterateFields(val)
		}
	case reflect.Struct:
		val := reflect.ValueOf(x)
		iterateFields(val)
	default:
		fmt.Printf("%s: %v\n", t, x)
	}
}

func iterateFields(val reflect.Value) {
	for i := 0; i < val.NumField(); i++ {
		t := val.Type()
		f := t.Field(i)
		if f.PkgPath != "" {
			fmt.Printf("%s: %v\n", f.Name, "***")
			continue
		}
		v := val.Field(i)
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		fmt.Printf("%s: %v\n", f.Name, v)
		if v.Kind() == reflect.Struct {
			iterateFields(v)
			continue
		}
	}
}
```
