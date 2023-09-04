package go_validator

type Profile struct {
	Username string `validate:"required,alphanum,min=4,max=20"`
	Password string `validate:"required,min=6,max=20"`
}

type Person struct {
	Name   string `validate:"required"`
	Height string `validate:"required,validHeight"`
}

func ValidateHeight(fl validator.FieldLevel) bool {
	height, err := strconv.Atoi(fl.Field().String())
	if err != nil {
		return false
	}

	minHeight := 30
	maxHeight := 300
	return height >= minHeight && height <= maxHeight
}
