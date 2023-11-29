# Go Validator

https://github.com/go-playground/validator

## Built-in

```go
type Profile struct {
  Username string `validate:"required,alphanum,min=4,max=20"`
  Password string `validate:"required,min=6,max=20"`
}

invalidProfile := Profile{
  Username: "user@name", // Contains non-alphanumeric characters
  Password: "short",     // Password is too short
}

val := validator.New()
err := val.Struct(invalidProfile)
if err != nil {
  fmt.Printf("Validation failed\n%s\n", err)
}
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
validate.RegisterValidation("validHeight", ValidateHeight)

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
