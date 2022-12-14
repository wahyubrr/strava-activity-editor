package service

type StravaService interface {
	GetAllActivities(page int, per_page int)
	GetOneActivities(activityId string)
	GetOneGear(gearId string)
}
