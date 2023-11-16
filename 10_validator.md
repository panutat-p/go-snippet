# Struct Validator

https://pkg.go.dev/github.com/go-playground/validator/v10

## go-playground/validator/v10

https://pkg.go.dev/github.com/go-playground/validator/v10

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
