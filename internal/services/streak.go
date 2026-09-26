package services

import "homepage-mock-api/internal/models"

type StreakService struct{}

func NewStreakService() *StreakService { return &StreakService{} }

func (s *StreakService) GetStreakCalendar(clientID string) []models.StreakDay {
	_ = clientID
	return []models.StreakDay{
		{Date: "2026-04-06", WeekdayLabel: "شنبه", DayOfMonth: 18, IsToday: false, MoodLogged: true, ExerciseLogged: true},
		{Date: "2026-04-07", WeekdayLabel: "یکشنبه", DayOfMonth: 19, IsToday: false, MoodLogged: true, ExerciseLogged: false},
		{Date: "2026-04-08", WeekdayLabel: "دوشنبه", DayOfMonth: 20, IsToday: false, MoodLogged: true, ExerciseLogged: true},
		{Date: "2026-04-09", WeekdayLabel: "سه‌شنبه", DayOfMonth: 21, IsToday: true, MoodLogged: false, ExerciseLogged: false},
		{Date: "2026-04-10", WeekdayLabel: "چهارشنبه", DayOfMonth: 22, IsToday: false, MoodLogged: false, ExerciseLogged: false},
		{Date: "2026-04-11", WeekdayLabel: "پنجشنبه", DayOfMonth: 23, IsToday: false, MoodLogged: false, ExerciseLogged: false},
		{Date: "2026-04-12", WeekdayLabel: "جمعه", DayOfMonth: 24, IsToday: false, MoodLogged: false, ExerciseLogged: false},
	}
}
