package validator

import (
	"net/mail"

	"github.com/asaskevich/govalidator"
)

type Validator struct{}

// ValidateStruct is a method to validate a struct using govalidator.
func (v Validator) ValidateStruct(i interface{}) (bool, error) {
	return govalidator.ValidateStruct(i)
}

func (v Validator) ValidateEmail(email string) error {
	_, err := mail.ParseAddress(email)
	return err
}
