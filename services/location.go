package services

import (
	"context"

	"github.com/JoaoGumiero/OngMais/entities"
	"github.com/JoaoGumiero/OngMais/firebase"
)

type LocationService struct {
	LocationRepository *firebase.LocationRepository
}

func NewLocationService(LocationRepo *firebase.LocationRepository) *LocationService {
	return &LocationService{LocationRepository: LocationRepo}
}

func (s *LocationService) GetStatesService(ctx context.Context) (states []entities.SimplifiedState, err error) {
	return s.LocationRepository.GetStatesDB(ctx)
}

func (s *LocationService) GetCitiesService(ctx context.Context) (cities []entities.SimplifiedCity, err error) {
	return s.LocationRepository.GetCitiesDB(ctx)
}
