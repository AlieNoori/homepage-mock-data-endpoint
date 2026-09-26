package services

import "homepage-mock-api/internal/models"

type ContentService struct{}

func NewContentService() *ContentService { return &ContentService{} }

func (s *ContentService) GetRecommended(clientID string) models.RecommendedContent {
	_ = clientID
	return models.RecommendedContent{
		Articles: []models.ArticleSummary{
			{
				ID:         "article_10",
				Title:      "وقتی بدنت زودتر از خودت نگران می‌شه",
				AuthorName: "دکتر مریم فتاحی",
				ImageURL:   "https://cdn.example.com/articles/body-anxiety.jpg",
			},
			{
				ID:         "article_11",
				Title:      "چطور به قضاوت دیگران باج ندیم؟",
				AuthorName: "دکتر امید رضازاده",
				ImageURL:   "https://cdn.example.com/articles/judgement.jpg",
			},
		},
		Podcasts: []models.PodcastSummary{
			{
				ID:              "podcast_1",
				Title:           "دکمهٔ اضطراری برای افکار تکراری",
				AuthorName:      "دکتر محمد سنایی",
				DurationMinutes: 32,
				ImageURL:        "https://cdn.example.com/podcasts/rumination.jpg",
			},
			{
				ID:              "podcast_2",
				Title:           "اگرهای بی‌پایان",
				AuthorName:      "دکتر ستاره میرزایی",
				DurationMinutes: 20,
				ImageURL:        "https://cdn.example.com/podcasts/what-ifs.jpg",
			},
		},
	}
}
