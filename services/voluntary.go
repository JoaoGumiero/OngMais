package services

import (
	"context"

	"github.com/JoaoGumiero/OngMais/entities"
	"github.com/JoaoGumiero/OngMais/firebase"
	"github.com/JoaoGumiero/OngMais/utils"
	"github.com/go-playground/validator"
)

type VoluntaryService struct {
	VoluntaryRepository *firebase.VoluntaryRepository
}

func NewVoluntaryService(VoluntaryRepo *firebase.VoluntaryRepository) *VoluntaryService {
	return &VoluntaryService{VoluntaryRepository: VoluntaryRepo}
}

// instance validator
var validate = validator.New()

func (s *VoluntaryService) AddVoluntaryService(voluntary *entities.Voluntary, ctx context.Context) error {
	// Make some validations here
	if err := validate.Struct(voluntary); err != nil {
		return err
	}
	// Include here validation of Email and Phone
	return s.VoluntaryRepository.AddVoluntaryDB(voluntary, ctx)
}

func (s *VoluntaryService) DeleteVoluntaryService(id string, ctx context.Context) error {
	// Include here id validation
	return s.VoluntaryRepository.DeleteVoluntaryDB(id, ctx)
}

func (s *VoluntaryService) UpdateVoluntaryService(id string, updates map[string]interface{}, ctx context.Context) error {
	if err := utils.ValidateUpdatesStruct(updates); err != nil {
		return err
	}
	// Include here validation of Email and Phone
	return s.VoluntaryRepository.UpdateVoluntaryDB(id, ctx, updates)
}

func (s *VoluntaryService) GetVoluntaryService(voluntary *entities.Voluntary, ctx context.Context) error {
	return nil
}

func (s *VoluntaryService) GetVoluntaryByIdService(id string, ctx context.Context) (voluntary *entities.Voluntary, err error) {
	// validate ID
	// if err := validate.; err != nil {
	//	return err
	// }
	return s.VoluntaryRepository.GetVoluntaryByIdDB(id, ctx)
}
