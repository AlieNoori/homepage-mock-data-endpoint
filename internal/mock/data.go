package mock

import "homepage-mock-api/internal/models"

func BuildHomeResponse() models.HomeResponse {
	return models.HomeResponse{
		User: models.UserSummary{
			ID:                "user_42",
			FullName:          "عاطفه",
			GreetingTimeOfDay: "morning",
			TreatmentDayCount: 83,
		},

		StreakCalendar: []models.StreakDay{
			{Date: "2026-04-06", WeekdayLabel: "شنبه", DayOfMonth: 18, IsToday: false, MoodLogged: true, ExerciseLogged: true},
			{Date: "2026-04-07", WeekdayLabel: "یکشنبه", DayOfMonth: 19, IsToday: false, MoodLogged: true, ExerciseLogged: false},
			{Date: "2026-04-08", WeekdayLabel: "دوشنبه", DayOfMonth: 20, IsToday: false, MoodLogged: true, ExerciseLogged: true},
			{Date: "2026-04-09", WeekdayLabel: "سه‌شنبه", DayOfMonth: 21, IsToday: true, MoodLogged: false, ExerciseLogged: false},
			{Date: "2026-04-10", WeekdayLabel: "چهارشنبه", DayOfMonth: 22, IsToday: false, MoodLogged: false, ExerciseLogged: false},
			{Date: "2026-04-11", WeekdayLabel: "پنجشنبه", DayOfMonth: 23, IsToday: false, MoodLogged: false, ExerciseLogged: false},
			{Date: "2026-04-12", WeekdayLabel: "جمعه", DayOfMonth: 24, IsToday: false, MoodLogged: false, ExerciseLogged: false},
		},

		MoodCheckIn: models.MoodCheckIn{
			Prompt:             "امروز حالت چطوره؟",
			CTALabel:           "ثبت علائم دیگه امروزت",
			AlreadyLoggedToday: false,
			QuickStates: []models.MoodState{
				{ID: 1, Label: "خیلی بد"},
				{ID: 2, Label: "بد"},
				{ID: 3, Label: "معمولی"},
				{ID: 4, Label: "خوب"},
				{ID: 5, Label: "خیلی خوب"},
			},
		},

		ActivePrograms: []models.TherapyProgramSummary{
			{
				EnrollmentID:         "enrol_1939",
				ProgramID:            "program_1919",
				ProgramName:          "اختلالات اضطرابی و وسواسی",
				TherapistID:          "user_42",
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
		},

		TodayExercise: models.TodayExerciesSection{
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
		},

		SuggestedExercise: models.SuggestedExercise{
			Headline:     "استرس داری؟",
			Prompt:       "امشب قبل از جلسه با دکتر حق‌شناس، مایلی یه تمرین کوتاه برای کاهش استرست انجام بدیم؟",
			CTALabel:     "مشاهده تمرین",
			ExerciseID:   "exc_2003",
			ExerciseName: "تمرین کاهش استرس سریع",
		},

		WellbeingSummary: models.WellbeingSummary{
			ComparisonPeriodLabel: "در هفته گذشته",
			Metrics: []models.MetricSummary{
				{Key: "focus", Label: "سطح تمرکز", ValuePercent: 65, ChangePercent: 5},
				{Key: "energy", Label: "سطح انرژی", ValuePercent: 36, ChangePercent: 4},
				{Key: "sleep_quality", Label: "کیفیت خواب", ValuePercent: 34, ChangePercent: 15},
				{Key: "stress", Label: "سطح استرس", ValuePercent: 46, ChangePercent: 0},
			},
		},

		MonthlyOverview: models.MonthlyOverview{
			MonthLabel:           "فروردین",
			Year:                 1405,
			OverallStatusText:    "اضطراب و نشخوارهای فکری تو در مجموع این ماه ۲۴٪ کاهش داشت.",
			MonthSaverText:       "در اوج حملات پانیک، ۵ بار از تمرین «تنفس با حباب» استفاده کردی و ضربان قلبت رو به موقع پایدار کردی.",
			BiggestWinText:       "با انجام مرتب تکالیف قبل از خواب دکتر حق‌شناس، کیفیت خوابت نسبت به شروع ماه ۱۴٪ بالاتر اومده.",
			ConsistencyBadgeName: "ثابت قدم",
			ConsistencyBadgeText: "با وجود روزهای سخت، ۴ جلسه درمانی رو کامل رفتی و زنجیره تمریناتت رو قطع نکردی. دمت گرم!",
			NextMonthCTA:         "برای ماه بعد برنامه‌ چیه؟",
		},

		Recommended: models.RecommendedContent{
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
		},
	}
}
