package utils

import (
	"encoding/json"

	"github.com/go-playground/validator"
)

type VoluntaryUpdate struct {
	Email string `validate:"omitempty,email"`
	Phone string `validate:"omitempty,phone_br"`
	Name  string `validate:"omitempty,min=2,max=100"`
}

func ValidateUpdatesStruct(updates map[string]interface{}) error {
	var updateData VoluntaryUpdate
	data, err := json.Marshal(updates)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, &updateData); err != nil {
		return err
	}

	validate := validator.New()
	if err := validate.Struct(updateData); err != nil {
		return err
	}
	return nil
}
