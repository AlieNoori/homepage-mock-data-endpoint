package services

import "homepage-mock-api/internal/models"

type ExerciseService struct{}

func NewExerciseService() *ExerciseService { return &ExerciseService{} }

func (s *ExerciseService) GetTodayExercises(clientID string) models.TodayExerciesSection {
	_ = clientID
	return models.TodayExerciesSection{
		Title:    "تمرین‌های درخواستی درمانگران",
		Subtitle: "حتما تا قبل جلسه بعدی انجام شوند",
		Exercise: []models.TodayExercies{
			{
				AssignmentID:    "asn_11228",
				ExerciseID:      "exc_2001",
				ExerciseName:    "یادداشت محرک‌های استرس",
				TherapistName:   "دکتر مهرداد حق‌شناس",
				DurationMinutes: 5,
				Status:          "completed",
			},
			{
				AssignmentID:    "asn_11229",
				ExerciseID:      "exc_2002",
				ExerciseName:    "تمرین تنفس",
				TherapistName:   "دکتر مهرداد حق‌شناس",
				DurationMinutes: 3,
				Status:          "pending",
			},
		},
	}
}

func (s *ExerciseService) GetSuggestedExercise(clientID string) models.SuggestedExercise {
	_ = clientID
	return models.SuggestedExercise{
		Headline:     "استرس داری؟",
		Prompt:       "امشب قبل از جلسه با دکتر حق‌شناس، مایلی یه تمرین کوتاه برای کاهش استرست انجام بدیم؟",
		CTALabel:     "مشاهده تمرین",
		ExerciseID:   "exc_2003",
		ExerciseName: "تمرین کاهش استرس سریع",
	}
}
