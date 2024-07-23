package firebase

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/JoaoGumiero/OngMais/entities"
	"google.golang.org/api/iterator"
)

type LocationRepository struct {
	client *firestore.Client
}

func NewLocationrRepository(client *firestore.Client) *LocationRepository {
	return &LocationRepository{client: client}
}

func (r *LocationRepository) GetStatesDB(ctx context.Context) (states []entities.SimplifiedState, err error) {
	iter := r.client.Collection("br-states").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		var state entities.SimplifiedState
		if err := doc.DataTo(&state); err != nil {
			return nil, err
		}
		states = append(states, state)
	}
	return states, nil
}

func (r *LocationRepository) GetCitiesDB(ctx context.Context) (cities []entities.SimplifiedCity, err error) {
	iter := r.client.Collection("br-cities").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		var city entities.SimplifiedCity
		if err := doc.DataTo(&city); err != nil {
			return nil, err
		}
		cities = append(cities, city)
	}
	return cities, nil

}
