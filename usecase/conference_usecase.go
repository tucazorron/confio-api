package usecase

import (
	"tucazorron/confio-api/model"
	"tucazorron/confio-api/repository"
)

type ConferenceUseCase struct {
	repository repository.ConferenceRepository
}

func NewConferenceUsecase(repository repository.ConferenceRepository) ConferenceUseCase {
	return ConferenceUseCase{
		repository: repository,
	}
}

func (uscs *ConferenceUseCase) GetConferences() ([]model.Conference, error) {
	return uscs.repository.GetConferences()
}

func (uscs *ConferenceUseCase) CreateConference(conference model.Conference) (model.Conference, error) {
	id, err := uscs.repository.CreateConference(conference)
	if err != nil {
		return model.Conference{}, err
	}
	conference.ID = id
	return conference, nil
}
