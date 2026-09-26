package services

import "homepage-mock-api/internal/models"

type UserService struct{}

func NewUserService() *UserService { return &UserService{} }

func (s *UserService) GetHomeUser(_clientID string) models.UserSummary {
	return models.UserSummary{
		ID:                "user_42",
		FullName:          "عاطفه",
		GreetingTimeOfDay: "morning",
		TreatmentDayCount: 83,
	}
}
