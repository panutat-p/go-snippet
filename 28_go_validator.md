# Go Validator

https://github.com/go-playground/validator

## Built-in

```go
type Profile struct {
    Username string `validate:"required,alphanum,min=4,max=20"`
    Password string `validate:"required,min=6,max=20"`
}

p := Profile{
    Username: "user@name",
    Password: "short",
}

validate := validator.New()
err := validate.Struct(p)
if err != nil {
    fmt.Println("🔴 Validation failed")
    fmt.Println(err)
    // Key: 'Profile.Username' Error:Field validation for 'Username' failed on the 'alphanum' tag
    // Key: 'Profile.Password' Error:Field validation for 'Password' failed on the 'min' tag
}
```

## Custom valivalidation failure message

https://docs.gofiber.io/guide/validation

https://github.com/go-playground/validator/blob/master/_examples/translations/main.go

```go
type Person struct {
    Name   string `validate:"required"`
    Email *string `validate:"omitnil,email"`
}

p := Person{
    Name: "John",
    Email: nil,
}

validate := validator.New()
err := validate.Struct(p)
if err != nil {
    var errv validator.ValidationErrors
    if errors.As(err, &errv) {
        fmt.Println("🔴 Validation failed")
        for _, e := range errv {
            fmt.Printf("Field: '%s', Condition: '%s', Actual: '%v'\n", e.Field(), e.Tag(), e.Param())
        }
    } else {
        fmt.Println(err)
    }
}
```

## Custom validator tag

```go
type Person struct {
    Name   string `validate:"required"`
    Height int    `validate:"required,validHeight"`
}

func ValidateHeight(fl validator.FieldLevel) bool {
    minHeight := 30
    maxHeight := 300
    height := fl.Field().Int()
    return height >= int64(minHeight) && height <= int64(maxHeight)
}
```

```go
validate := validator.New()
err := validate.RegisterValidation("validHeight", ValidateHeight)
if err != nil {
    panic(err)
}

person := Person{
    Name:   "John",
    Height: 5, // Height is out of the valid range
}
err := validate.Struct(person)
if err != nil {
    fmt.Println("🔴 Validation failed")
    fmt.Println(err) // Key: 'Person.Height' Error:Field validation for 'Height' failed on the 'validHeight' tag
}
```

## Register types

```go
package main

import (
    "database/sql"
    "database/sql/driver"
    "reflect"

    "github.com/go-playground/validator/v10"
)

func main() {
    validate := validator.New()
    validate.RegisterCustomTypeFunc(
        ValidateValuer,
        sql.NullString{},
        sql.NullInt64{},
        sql.NullInt32{},
        sql.NullInt16{},
        sql.NullBool{},
        sql.NullFloat64{},
        sql.NullFloat32{},
    )
}

func ValidateValuer(field reflect.Value) any {
    val, ok := field.Interface().(driver.Valuer)
    if !ok {
        return nil
    }
    v, err := val.Value()
    if err != nil {
        return nil
    }
    return v
}
```
