package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/JoaoGumiero/OngMais/config"
	"github.com/JoaoGumiero/OngMais/entities"
	"github.com/JoaoGumiero/OngMais/firebase"
	"github.com/JoaoGumiero/OngMais/utils"
)

// Cron to periodically fetch data from API
func CronStateCityFetch(conf *config.Config) error {
	// Once every 6 months
	rt := utils.NewRealTicker(6 * 30 * 24 * time.Hour)
	defer rt.Stop()

	// New context to manually define a timeout period
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	Client := firebase.InitFirebase(*conf)
	for {
		select {
		case <-rt.C():
			fmt.Println("City/State fetch timer called")
			states, err := FetchStates()
			if err != nil {
				log.Fatalf("Error fetching states: %v", err)
			}
			firebase.StoreStates(states, Client, ctx)

			cities, err := FetchCities()
			if err != nil {
				log.Fatalf("Error fetching cities: %v", err)
			}
			firebase.StoreCities(cities, Client, ctx)
		case <-ctx.Done():
			rt.Stop()
			return nil
		}
	}
}

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
