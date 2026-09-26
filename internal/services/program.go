package services

import "homepage-mock-api/internal/models"

type ProgramService struct{}

func NewProgramService() *ProgramService { return &ProgramService{} }

func (s *ProgramService) GetActivePrograms(clientID string) []models.TherapyProgramSummary {
	_ = clientID
	return []models.TherapyProgramSummary{
		{
			EnrollmentID:         "enrol_1939",
			ProgramID:            "program_1919",
			ProgramName:          "اختلالات اضطرابی و وسواسی",
			TherapistID:          "user_therapist_1",
			TherapistName:        "دکتر مهرداد حق‌شناس",
			SessionsCompleted:    1,
			SessionsTotal:        8,
			TreatmentWeekCurrent: 3,
			TreatmentWeekTotal:   10,
			ExercisesCompleted:   8,
			ExercisesTotal:       10,
			ProgressPercent:      20,
			Stages: []models.ProgramStage{
				{Order: 1, Name: "ارزیابی", Completed: false, Current: true},
				{Order: 2, Name: "پایدارسازی", Completed: false, Current: false},
				{Order: 3, Name: "تثبیت", Completed: false, Current: false},
				{Order: 4, Name: "بهبود جامع", Completed: false, Current: false},
			},
		},
	}
}
