package services

import "homepage-mock-api/internal/models"

type MoodService struct{}

func NewMoodService() *MoodService { return &MoodService{} }

func (s *MoodService) GetMoodCheckIn(clientID string) models.MoodCheckIn {
	_ = clientID
	return models.MoodCheckIn{
		Prompt:             "امروز حالت چطوره؟",
		CTALabel:           "ثبت علائم دیگه امروزت",
		AlreadyLoggedToday: false,
		QuickStates: []models.MoodState{
			{ID: 1, Label: "خیلی بد"},
			{ID: 2, Label: "بد"},
			{ID: 3, Label: "معمولی"},
			{ID: 4, Label: "خوب"},
			{ID: 5, Label: "خیلی خوب"},
		},
	}
}
