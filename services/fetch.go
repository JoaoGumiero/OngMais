package services

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/JoaoGumiero/OngMais/entities"
)

func FetchStates() ([]entities.SimplifiedState, error) {
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

	// Use a whole state response to catch every data from the IBGE API
	var stateResp []entities.StateAPIResponse
	err = json.Unmarshal(body, &stateResp)
	if err != nil {
		return nil, err
	}

	// Iterates through the responde API and parse/append data to the a simplier API
	var stateList []entities.SimplifiedState
	for _, s := range stateResp {
		// str := strconv.Itoa(s.ID) // Here i could create a function that parse int into UUID, probably gonna do it next time.
		state := entities.SimplifiedState{
			ID:   s.ID,
			Name: s.Nome,
		}
		stateList = append(stateList, state)
	}

	return stateList, nil
}

func FetchCities() ([]entities.SimplifiedCity, error) {
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

	var cityResp []entities.CityAPIResponse
	err = json.Unmarshal(body, &cityResp)
	if err != nil {
		return nil, err
	}

	// Iterates through the responde API and parse/append data to the a simplier API
	var citiesList []entities.SimplifiedCity
	for _, c := range cityResp {
		// str := strconv.Itoa(c.ID)
		// str2 := strconv.Itoa(c.Microrregiao.Mesorregiao.UF.ID) // Here i could create a function that parse int into UUID, probably gonna do it next time.
		city := entities.SimplifiedCity{
			ID:   c.ID,
			Name: c.Nome,
			State: entities.SimplifiedState{
				ID:   c.Microrregiao.Mesorregiao.UF.ID,
				Name: c.Microrregiao.Mesorregiao.UF.Nome,
			},
		}
		citiesList = append(citiesList, city)
	}

	return citiesList, nil
}
