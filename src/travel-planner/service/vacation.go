package service

import (
	"errors"
	"travel-planner/backend"
	"travel-planner/model"
)

func GetVacationsInfo() ([]model.Vacation, error) {
	vacations, err := backend.DB.GetVacations()
	if err != nil {
		return nil, err
	}

	if vacations == nil || len(vacations) == 0 {
		return nil, errors.New("empty or invalid vacations, check the Database")
	}
	return vacations, nil
}

func AddVacation(vacation *model.Vacation) (bool, error) {
	success, err := backend.DB.SaveVacation(vacation)
	return success, err
}

func GetActivitiesInfoFromPlanId(planId uint32) ([]model.Activity, error) {
	activities, err := backend.DB.GetActivityFromPlanId(planId)
	if err != nil {
		return nil, err
	}

	if activities == nil || len(activities) == 0 {
		return nil, errors.New("empty or invalid vacations, check the Database")
	}
	return activities, nil
}

func SaveVacationPlan(plan model.Plan) (error) {
	err := backend.DB.SaveVacationPlanToSQL(plan)
	return err
}

func GetPlanInfoFromVactionId(vacationId uint32) ([]model.Plan, error) {
	plans, err := backend.DB.GetPlanFromVacationId(vacationId)
	if err != nil {
		return nil, err
	}

	if plans== nil || len(plans) == 0 {
		return nil, errors.New("empty or invalid plan, check the Database")
	}
	return plans, nil
}

func GetTransportationFromPlanId(planId uint32) ([]model.Transportaion, error) {
	transportations, err := backend.DB.GetTransportationFromPlanId(planId)
	if err != nil {
		return nil, err
	}

	if transportations == nil || len(transportations) == 0 {
		return nil, errors.New("empty or invalid vacations, check the Database")
	}
	return transportations, nil
}

// func GetSitesFromActivityId(activityId uint32) ([]model.Activity, error) {
// 	sites, err := backend.DB.GetSitesFromActivityId(activityId)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if sites == nil || len(sites) == 0 {
// 		return nil, errors.New("empty or invalid vacations, check the Database")
// 	}
// 	return sites, nil
// }


func GetSiteFromSiteyId(siteId uint32) (*model.Site, error) {
	site, err := backend.DB.GetSiteFromSiteId(siteId)
	if err != nil {
		return nil, err
	}

	if site == nil {
		return nil, errors.New("empty or invalid vacations, check the Database")
	}
	return site, nil
}

