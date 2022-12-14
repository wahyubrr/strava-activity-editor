package service

import "fmt"

type StravaServiceImpl struct {
	URI string
}

func NewStravaService(uri string) StravaService {
	return &StravaServiceImpl{
		URI: uri,
	}
}

func (service *StravaServiceImpl) GetAllActivities(page int, per_page int) {
	fmt.Println("service: get activities", page, per_page)
}

func (service *StravaServiceImpl) GetOneActivities(activityId string) {
	fmt.Println("service: activityId", activityId)
}

func (service *StravaServiceImpl) GetOneGear(gearId string) {
	fmt.Println("service: gearId", gearId)
}
