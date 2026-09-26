package services

import "homepage-mock-api/internal/models"

type WellbeingService struct{}

func NewWellbeingService() *WellbeingService { return &WellbeingService{} }

func (s *WellbeingService) GetWellbeingSummary(clientID string) models.WellbeingSummary {
	_ = clientID
	return models.WellbeingSummary{
		ComparisonPeriodLabel: "در هفته گذشته",
		Metrics: []models.MetricSummary{
			{Key: "focus", Label: "سطح تمرکز", ValuePercent: 65, ChangePercent: 5},
			{Key: "energy", Label: "سطح انرژی", ValuePercent: 36, ChangePercent: 4},
			{Key: "sleep_quality", Label: "کیفیت خواب", ValuePercent: 34, ChangePercent: 15},
			{Key: "stress", Label: "سطح استرس", ValuePercent: 46, ChangePercent: 0},
		},
	}
}
