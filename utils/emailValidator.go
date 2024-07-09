package utils

import (
	"net/mail"
)

type EmailValidator struct {
	Email string
}

// Standard lib email validator
func (v EmailValidator) IsValid() bool {
	_, err := mail.ParseAddress(v.Email)
	return err == nil
}
