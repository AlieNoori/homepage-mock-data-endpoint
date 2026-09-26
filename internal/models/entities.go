package models

import "time"

type UserRole string

const (
	RoleClient    UserRole = "client"
	RoleTherapist UserRole = "therapist"
	RoleAdmin     UserRole = "admin"
)

type User struct {
	ID           string    `json:"id"`
	Role         UserRole  `json:"role"`
	FullName     string    `json:"full_name"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	Timezone     string    `json:"timezone"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type TherapistProfile struct {
	UserID        string   `json:"user_id"`
	LicenseNumber *string  `json:"license_number,omitempty"`
	Specialties   []string `json:"specialties,omitempty"`
	Bio           *string  `json:"bio,omitempty"`
}

type PatientProfile struct {
	UserID             string     `json:"user_id"`
	DateOfBirth        *time.Time `json:"date_of_birth,omitempty"`
	PrimaryTherapistID *string    `json:"primary_therapist_id,omitempty"`
	OnboardingNotes    *string    `json:"onboarding_notes,omitempty"`
}

type TherapyProgram struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	Description   *string   `json:"description,omitempty"`
	DurationWeeks int16     `json:"duration_weeks"`
	CreatedBy     string    `json:"created_by"`
	CreatedAt     time.Time `json:"created_at"`
}

type EnrollmentStatus string

const (
	EnrollmentActive    EnrollmentStatus = "active"
	EnrollmentPaused    EnrollmentStatus = "paused"
	EnrollmentCompleted EnrollmentStatus = "completed"
	EnrollmentCancelled EnrollmentStatus = "cancelled"
)

type ProgramEnrollment struct {
	ID          string           `json:"id"`
	ClientID    string           `json:"client_id"`
	TherapistID string           `json:"therapist_id"`
	ProgramID   string           `json:"program_id"`
	StartDate   time.Time        `json:"start_date"`
	EndDate     *time.Time       `json:"end_date,omitempty"`
	Status      EnrollmentStatus `json:"status"`
	CreatedAt   time.Time        `json:"created_at"`
}

type SessionStatus string

const (
	SessionScheduled SessionStatus = "scheduled"
	SessionCompleted SessionStatus = "completed"
	SessionNoShow    SessionStatus = "no_show"
	SessionCancelled SessionStatus = "cancelled"
)

type TherapySession struct {
	ID              string        `json:"id"`
	EnrollmentID    string        `json:"enrollment_id"`
	TherapistID     string        `json:"therapist_id"`
	ScheduledAt     time.Time     `json:"scheduled_at"`
	DurationMinutes int16         `json:"duration_minutes"`
	Status          SessionStatus `json:"status"`
	TherapistNotes  *string       `json:"therapist_notes,omitempty"`
	ClientTakeaway  *string       `json:"client_takeaway,omitempty"`
	CreatedAt       time.Time     `db:"created_at" json:"created_at"`
}

type Exercise struct {
	ID                     string  `json:"id"`
	ProgramID              *string `json:"program_id,omitempty"`
	Name                   string  `json:"name"`
	Category               *string `json:"category,omitempty"`
	Description            *string `json:"description,omitempty"`
	DefaultDurationMinutes *int16  `json:"default_duration_minutes,omitempty"`
	CreatedBy              *string `json:"created_by,omitempty"`
}

type AssignmentStatus string

const (
	AssignmentPending   AssignmentStatus = "pending"
	AssignmentCompleted AssignmentStatus = "completed"
	AssignmentSkipped   AssignmentStatus = "skipped"
	AssignmentExpired   AssignmentStatus = "expired"
)

type ExerciseAssignment struct {
	ID            string           `json:"id"`
	EnrollmentID  string           `json:"enrollment_id"`
	ExerciseID    string           `json:"exercise_id"`
	ScheduledDate time.Time        `json:"scheduled_date"`
	Status        AssignmentStatus `json:"status"`
}

type ExerciseLog struct {
	ID                   string    `json:"id"`
	AssignmentID         string    `json:"assignment_id"`
	CompletedAt          time.Time `json:"completed_at"`
	DurationMinutes      *int16    `json:"duration_minutes,omitempty"`
	PerceivedHelpfulness *int16    `json:"perceived_helpfulness,omitempty"`
	Notes                *string   `json:"notes,omitempty"`
}

type MoodState struct {
	ID    int    `json:"id"`
	Label string `json:"label"`
}

type MoodLog struct {
	ID                  string    `json:"id"`
	ClientID            string    `json:"client_id"`
	MoodStateID         *int16    `json:"mood_state_id,omitempty"`
	DetailedDescription *string   `json:"detailed_description,omitempty"`
	Tags                []string  `json:"tags,omitempty"`
	LoggedAt            time.Time `json:"logged_at"`
}

type DailyMetric struct {
	ID           string    `json:"id"`
	ClientID     string    `json:"client_id"`
	MetricDate   time.Time `json:"metric_date"`
	StressLevel  *int16    `json:"stress_level,omitempty"`
	FocusLevel   *int16    `json:"focus_level,omitempty"`
	EnergyLevel  *int16    `json:"energy_level,omitempty"`
	SleepQuality *int16    `json:"sleep_quality,omitempty"`
	SleepHours   *float64  `json:"sleep_hours,omitempty"`
}

type ContentStatus string

const (
	ContentDraft     ContentStatus = "draft"
	ContentPublished ContentStatus = "published"
	ContentArchived  ContentStatus = "archived"
)

type Article struct {
	ID              string        `json:"id"`
	Slug            string        `json:"slug"`
	Title           string        `json:"title"`
	Excerpt         *string       `json:"excerpt,omitempty"`
	Body            string        `json:"body"`
	AuthorID        *string       `json:"author_id,omitempty"`
	Category        *string       `json:"category,omitempty"`
	Tags            []string      `json:"tags,omitempty"`
	ImageURL        *string       `json:"image_url,omitempty"`
	ReadTimeMinutes *int16        `json:"read_time_minutes,omitempty"`
	Status          ContentStatus `json:"status"`
	PublishedAt     *time.Time    `json:"published_at,omitempty"`
	CreatedAt       time.Time     `json:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at"`
}

type Podcast struct {
	ID              string        `json:"id"`
	Slug            string        `json:"slug"`
	Title           string        `json:"title"`
	Description     *string       `json:"description,omitempty"`
	AudioURL        string        `json:"audio_url"`
	AuthorID        *string       `json:"author_id,omitempty"`
	AuthorName      string        `json:"author_name"`
	EpisodeNumber   *int          `json:"episode_number,omitempty"`
	DurationMinutes int16         `json:"duration_minutes"`
	ImageURL        *string       `json:"image_url,omitempty"`
	Tags            []string      `json:"tags,omitempty"`
	Status          ContentStatus `json:"status"`
	PublishedAt     *time.Time    `json:"published_at,omitempty"`
	CreatedAt       time.Time     `json:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at"`
}
