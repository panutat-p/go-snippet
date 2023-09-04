package go_validator

type Profile struct {
	Username string `validate:"required,alphanum,min=4,max=20"`
	Password string `validate:"required,min=6,max=20"`
}

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
