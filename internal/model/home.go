package models

type HomeResponse struct {
	User              UserSummary             `json:"user"`
	StreakCalendar    []StreakDay             `json:"streak_calendar"`
	MoodCheckIn       MoodCheckIn             `json:"mood_check_in"`
	ActivePrograms    []TherapyProgramSummary `json:"active_programs"`
	TodayTasks        TodayTasksSection       `json:"today_tasks"`
	SuggestedExercise SuggestedExercise       `json:"suggested_exercise"`
	WellbeingSummary  WellbeingSummary        `json:"wellbeing_summary"`
	MonthlyOverview   MonthlyOverview         `json:"monthly_overview"`
	Recommended       RecommendedContent      `json:"recommended_content"`
}

type UserSummary struct {
	ID                string `json:"id"`
	FullName          string `json:"full_name"`
	GreetingTimeOfDay string `json:"greeting_time_of_day"`
	TreatmentDayCount int    `json:"treatment_day_count"`
}

type StreakDay struct {
	Date           string `json:"date"`
	WeekdayLabel   string `json:"weekday_label"`
	DayOfMonth     int    `json:"day_of_month"`
	IsToday        bool   `json:"is_today"`
	MoodLogged     bool   `json:"mood_logged"`
	ExerciseLogged bool   `json:"exercise_logged"`
}

type MoodCheckIn struct {
	Prompt             string      `json:"prompt"`
	CTALabel           string      `json:"cta_label"`
	AlreadyLoggedToday bool        `json:"already_logged_today"`
	QuickStates        []MoodState `json:"quick_states"`
}

type TherapyProgramSummary struct {
	EnrollmentID         string         `json:"enrollment_id"`
	ProgramID            string         `json:"program_id"`
	ProgramName          string         `json:"program_name"`
	TherapistID          string         `json:"therapist_id"`
	TherapistName        string         `json:"therapist_name"`
	SessionsCompleted    int            `json:"sessions_completed"`
	SessionsTotal        int            `json:"sessions_total"`
	TreatmentWeekCurrent int            `json:"treatment_week_current"`
	TreatmentWeekTotal   int            `json:"treatment_week_total"`
	ExercisesCompleted   int            `json:"exercises_completed"`
	ExercisesTotal       int            `json:"exercises_total"`
	ProgressPercent      int            `json:"progress_percent"`
	Stages               []ProgramStage `json:"stages"`
}

type ProgramStage struct {
	Order     int    `json:"order"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
	Current   bool   `json:"current"`
}

type TodayTasksSection struct {
	Title    string      `json:"title"`
	Subtitle string      `json:"subtitle"`
	Tasks    []TodayTask `json:"tasks"`
}

type TodayTask struct {
	AssignmentID    string `json:"assignment_id"`
	ExerciseID      string `json:"exercise_id"`
	ExerciseName    string `json:"exercise_name"`
	TherapistName   string `json:"therapist_name"`
	DurationMinutes int    `json:"duration_minutes"`
	Status          string `json:"status"`
}

type SuggestedExercise struct {
	Headline     string `json:"headline"`
	Prompt       string `json:"prompt"`
	CTALabel     string `json:"cta_label"`
	ExerciseID   string `json:"exercise_id"`
	ExerciseName string `json:"exercise_name"`
}

type WellbeingSummary struct {
	ComparisonPeriodLabel string          `json:"comparison_period_label"`
	Metrics               []MetricSummary `json:"metrics"`
}

type MetricSummary struct {
	Key           string `json:"key"`
	Label         string `json:"label"`
	ValuePercent  int    `json:"value_percent"`
	ChangePercent int    `json:"change_percent"`
}

type MonthlyOverview struct {
	MonthLabel           string `json:"month_label"`
	Year                 int    `json:"year"`
	OverallStatusText    string `json:"overall_status_text"`
	MonthSaverText       string `json:"month_saver_text"`
	BiggestWinText       string `json:"biggest_win_text"`
	ConsistencyBadgeName string `json:"consistency_badge_name"`
	ConsistencyBadgeText string `json:"consistency_badge_text"`
	NextMonthCTA         string `json:"next_month_cta"`
}

type RecommendedContent struct {
	Articles []ArticleSummary `json:"articles"`
	Podcasts []PodcastSummary `json:"podcasts"`
}

type ArticleSummary struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	AuthorName string `json:"author_name"`
	ImageURL   string `json:"image_url"`
}

type PodcastSummary struct {
	ID              string `json:"id"`
	Title           string `json:"title"`
	AuthorName      string `json:"author_name"`
	DurationMinutes int    `json:"duration_minutes"`
	ImageURL        string `json:"image_url"`
}
