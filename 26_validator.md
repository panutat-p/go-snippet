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

## Print validation failure message

```go
validate := validator.New()
err := validate.Struct(p)
if err != nil {
    var ve validator.ValidationErrors
    if errors.As(err, &ve) {
        fmt.Println("🔴 Validation failed")
        for _, err := range ve {
            fmt.Printf("Field: '%s', Condition: '%s', Actual: '%v'\n", err.Field(), err.Tag(), err.Param())
        }
    } else {
        fmt.Println(err)
    }
}
```

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
fmt.Println(err) // nil
```

## Custom

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
