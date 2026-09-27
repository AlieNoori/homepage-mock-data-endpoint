package services

import (
	"sync"

	"homepage-mock-api/internal/models"
)

const MAX_WORKER = 10

type job func()

type pool struct {
	jobs chan job
	wg   sync.WaitGroup
}

func newPool(workers int) *pool {
	if workers < 1 {
		workers = 1
	}

	p := &pool{jobs: make(chan job)}

	p.wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			defer p.wg.Done()
			for j := range p.jobs {
				j()
			}
		}()
	}

	return p
}

func (p *pool) Submit(j job) {
	p.jobs <- j
}

func (p *pool) Close() {
	close(p.jobs)
	p.wg.Wait()
}

type HomeService struct {
	users     *UserService
	streaks   *StreakService
	moods     *MoodService
	programs  *ProgramService
	exercises *ExerciseService
	wellbeing *WellbeingService
	monthly   *MonthlyService
	content   *ContentService

	pool *pool
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

		pool: newPool(MAX_WORKER),
	}
}

func (s *HomeService) Build(clientID string) models.HomeResponse {
	var resp models.HomeResponse

	jobs := []job{
		func() { resp.User = s.users.GetHomeUser(clientID) },
		func() { resp.StreakCalendar = s.streaks.GetStreakCalendar(clientID) },
		func() { resp.MoodCheckIn = s.moods.GetMoodCheckIn(clientID) },
		func() { resp.ActivePrograms = s.programs.GetActivePrograms(clientID) },
		func() { resp.TodayExercise = s.exercises.GetTodayExercises(clientID) },
		func() { resp.SuggestedExercise = s.exercises.GetSuggestedExercise(clientID) },
		func() { resp.WellbeingSummary = s.wellbeing.GetWellbeingSummary(clientID) },
		func() { resp.MonthlyOverview = s.monthly.GetMonthlyOverview(clientID) },
		func() { resp.Recommended = s.content.GetRecommended(clientID) },
	}

	for _, j := range jobs {
		s.pool.Submit(j)
	}

	return resp
}

func (s *HomeService) Close() {
	s.pool.Close()
}
