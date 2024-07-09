package entities

type State struct {
	ID   string `json:"id" validate:"required,uuid4"`
	Name string `json:"name" validate:"required"`
}

type City struct {
	ID    string `json:"id" validate:"required,uuid4"`
	Name  string `json:"name" validate:"required"`
	State State  `json:"state_city" validate:"required"`
}

type SubjectInterest struct {
	ID   string `json:"id" validate:"required,uuid4"`
	Name string `json:"name" validate:"required"`
}

type Voluntary struct {
	ID              string          `json:"id" validate:"omitempty,uuid4"`
	Name            string          `json:"name" validate:"required"`
	Email           string          `json:"email" validate:"required,email"`
	Phone           int             `json:"phone" validate:"required,phone_br"`
	State           State           `json:"state" validate:"required"`
	City            City            `json:"city" validate:"required"`
	SubjectInterest SubjectInterest `json:"Subject_Interest" validate:"required"`
}
