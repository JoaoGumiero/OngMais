package entities

import "github.com/google/uuid"

type Voluntary struct {
	ID    uuid.UUID       `json:"id" validate:"omitempty,uuid4" firestore:"ID"`
	Name  string          `json:"name" validate:"required" firestore:"Name"`
	Email string          `json:"email" validate:"required,email" firestore:"Email"`
	Phone int             `json:"phone" validate:"required,phone_br" firestore:"Phone"`
	State SimplifiedState `json:"state" validate:"required" firestore:"State"`
	City  SimplifiedCity  `json:"city" validate:"required" firestore:"City"`
}

type SimplifiedState struct {
	ID   int    `json:"id" validate:"required,uuid4" firestore:"ID"`
	Name string `json:"name" validate:"required" firestore:"Name"`
}

// I could place here "Sigla" and "Região"
type SimplifiedCity struct {
	ID    int             `json:"id" validate:"required,uuid4" firestore:"ID"`
	Name  string          `json:"name" validate:"required" firestore:"Name"`
	State SimplifiedState `json:"city_state" validate:"required" firestore:"State"`
}

// IBGE API response structs
type Region struct {
	ID    int    `json:"id"`
	Sigla string `json:"sigla"`
	Nome  string `json:"nome"`
}

// IBGE API response structs
type StateAPIResponse struct {
	ID     int    `json:"id"`
	Sigla  string `json:"sigla"`
	Nome   string `json:"nome"`
	Regiao Region `json:"regiao"`
}

// IBGE API response structs
type CityAPIResponse struct {
	ID           int    `json:"id"`
	Nome         string `json:"nome"`
	Microrregiao struct {
		ID          int    `json:"id"`
		Nome        string `json:"nome"`
		Mesorregiao struct {
			ID   int              `json:"id"`
			Nome string           `json:"nome"`
			UF   StateAPIResponse `json:"UF"`
		} `json:"mesorregiao"`
	} `json:"microrregiao"`
}
