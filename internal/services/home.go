package services

import "homepage-mock-api/internal/models"

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
	return models.HomeResponse{
		User:              s.users.GetHomeUser(clientID),
		StreakCalendar:    s.streaks.GetStreakCalendar(clientID),
		MoodCheckIn:       s.moods.GetMoodCheckIn(clientID),
		ActivePrograms:    s.programs.GetActivePrograms(clientID),
		TodayExercise:     s.exercises.GetTodayExercises(clientID),
		SuggestedExercise: s.exercises.GetSuggestedExercise(clientID),
		WellbeingSummary:  s.wellbeing.GetWellbeingSummary(clientID),
		MonthlyOverview:   s.monthly.GetMonthlyOverview(clientID),
		Recommended:       s.content.GetRecommended(clientID),
	}
}
