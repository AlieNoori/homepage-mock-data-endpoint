package services

import "homepage-mock-api/internal/models"

type MonthlyService struct{}

func NewMonthlyService() *MonthlyService { return &MonthlyService{} }

func (s *MonthlyService) GetMonthlyOverview(clientID string) models.MonthlyOverview {
	_ = clientID
	return models.MonthlyOverview{
		MonthLabel:           "فروردین",
		Year:                 1405,
		OverallStatusText:    "اضطراب و نشخوارهای فکری تو در مجموع این ماه ۲۴٪ کاهش داشت.",
		MonthSaverText:       "در اوج حملات پانیک، ۵ بار از تمرین «تنفس با حباب» استفاده کردی و ضربان قلبت رو به موقع پایدار کردی.",
		BiggestWinText:       "با انجام مرتب تکالیف قبل از خواب دکتر حق‌شناس، کیفیت خوابت نسبت به شروع ماه ۱۴٪ بالاتر اومده.",
		ConsistencyBadgeName: "ثابت قدم",
		ConsistencyBadgeText: "با وجود روزهای سخت، ۴ جلسه درمانی رو کامل رفتی و زنجیره تمریناتت رو قطع نکردی. دمت گرم!",
		NextMonthCTA:         "برای ماه بعد برنامه‌ چیه؟",
	}
}
