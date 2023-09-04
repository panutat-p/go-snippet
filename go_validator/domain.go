package go_validator

type Profile struct {
	Username string `validate:"required,alphanum,min=4,max=20"`
	Password string `validate:"required,min=6,max=20"`
}
