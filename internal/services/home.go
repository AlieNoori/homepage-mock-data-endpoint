package services

import (
	"sync"

	"homepage-mock-api/internal/models"
)

type HomeService struct {
	users     *UserService
	streaks   *StreakService
	moods     *MoodService
	programs  *ProgramService
	exercises *ExerciseService
	wellbeing *WellbeingService
	monthly   *MonthlyService
	content   *ContentService
}

func NewHomeService() *HomeService {
	return &HomeService{
		users:     NewUserService(),
		streaks:   NewStreakService(),
		moods:     NewMoodService(),
		programs:  NewProgramService(),
		exercises: NewExerciseService(),
		wellbeing: NewWellbeingService(),
		monthly:   NewMonthlyService(),
		content:   NewContentService(),
	}
}

func (s *HomeService) Build(clientID string) models.HomeResponse {
	var resp models.HomeResponse

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		resp.User = s.users.GetHomeUser(clientID)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		resp.StreakCalendar = s.streaks.GetStreakCalendar(clientID)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		resp.MoodCheckIn = s.moods.GetMoodCheckIn(clientID)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		resp.ActivePrograms = s.programs.GetActivePrograms(clientID)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		resp.TodayExercise = s.exercises.GetTodayExercises(clientID)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		resp.SuggestedExercise = s.exercises.GetSuggestedExercise(clientID)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		resp.WellbeingSummary = s.wellbeing.GetWellbeingSummary(clientID)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		resp.MonthlyOverview = s.monthly.GetMonthlyOverview(clientID)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		resp.Recommended = s.content.GetRecommended(clientID)
	}()

	return resp
}
