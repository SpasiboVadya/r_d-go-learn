package reflection

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
)

type Person struct {
	Username string `validate:"required,min=3,max=20"`
	Email    string `validate:"required,email"`
	Age      int    `validate:"gte=18,lte=130"`
}

func RunValidator() {
	validate := validator.New()

	u := Person{
		Username: "jo",
		Email:    "notemail",
		Age:      15,
	}

	err := validate.Struct(u)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			fmt.Printf("Field '%s' failed on '%s' validation\n", e.Field(), e.Tag())
		}
	}
}

func (p Person) ValidatePerson() error {
	if len(p.Username) < 3 || len(p.Username) > 20 {
		return errors.New("username must be between 3 and 20 characters")
	}
	if len(p.Email) == 0 || !isEmail(p.Email) {
		return errors.New("email address is required")
	}
	if p.Age < 18 || p.Age > 130 {
		return errors.New("age must be between 18 and 130")
	}
	return nil
}

func isEmail(s string) bool {
	return len(s) >= 3 && s[len(s)-4:] == ".com" // simplified email check
}
