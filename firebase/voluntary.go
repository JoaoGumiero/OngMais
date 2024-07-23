package firebase

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/JoaoGumiero/OngMais/entities"
)

type VoluntaryRepository struct {
	client *firestore.Client
}

func NewUserRepository(client *firestore.Client) *VoluntaryRepository {
	return &VoluntaryRepository{client: client}
}

func (r *VoluntaryRepository) AddVoluntaryDB(voluntary *entities.Voluntary, ctx context.Context) error {
	_, _, err := r.client.Collection("voluntaries").Add(ctx, voluntary)
	if err != nil {
		return err
	}
	return nil

	// voluntary.ID, err = uuid.Parse(docRef.ID) // Optionally update the ID with the Firestore generated ID
	// if err != nil {
	// http.Error(w, "Failed parse str to uuid: "+err.Error(), http.StatusInternalServerError)
	// return
	// }
}

func (r *VoluntaryRepository) DeleteVoluntaryDB(id string, ctx context.Context) error {
	_, err := r.client.Collection("voluntaries").Doc(id).Delete(ctx)
	return err
}

func (r *VoluntaryRepository) UpdateVoluntaryDB(id string, ctx context.Context, updates map[string]interface{}) error {
	_, err := r.client.Collection("voluntaries").Doc(id).Set(ctx, updates, firestore.MergeAll)
	return err
}

// Implementation if we need later to retrieve all voluntaries
func (r *VoluntaryRepository) GetVoluntaryDB() error {
	return nil
}

func (r *VoluntaryRepository) GetVoluntaryByIdDB(id string, ctx context.Context) (voluntary *entities.Voluntary, err error) {
	doc, err := r.client.Collection("voluntaries").Doc(id).Get(ctx)
	if err != nil {
		return nil, err
	}
	doc.DataTo(&voluntary)
	return voluntary, nil
}
