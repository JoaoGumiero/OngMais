package services

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/JoaoGumiero/OngMais/entities"
)

func FetchStates() ([]entities.State, error) {
	url := "https://servicodados.ibge.gov.br/api/v1/localidades/estados"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var states []entities.State
	err = json.Unmarshal(body, &states)
	if err != nil {
		return nil, err
	}

	return states, nil
}

func FetchCities() ([]entities.City, error) {
	url := "https://servicodados.ibge.gov.br/api/v1/localidades/municipios"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var cities []entities.City
	err = json.Unmarshal(body, &cities)
	if err != nil {
		return nil, err
	}

	return cities, nil
}
