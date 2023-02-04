package backend

import (
	"fmt"
	"travel-planner/model"
)

/*
 Get list of activities within a plan with plan_id = @{planId}, sorted by ascending start_time
 */
func (backend *MySQLBackend) GetActivityFromPlanId(planId uint32) ([]model.Activity, error) {
	var activitiesResult []model.Activity
	result := backend.db.
				Table("Activities").
				Where("plan_id = ?", planId).
				Order("start_time").
				Find(&activitiesResult)
	if result.Error != nil {
		return nil, result.Error
	}

	return activitiesResult, nil
}

func (backend *MySQLBackend) SaveTransportation(transportation *model.Transportation) (bool, error) {
	result := backend.db.Table("Transportations").Create(&transportation)
	if err := result.Error; err != nil {
		return false, err
	}
	fmt.Println("Transportation saved in db")
	return true, nil
}

func (backend *MySQLBackend) SaveActivity(activity *model.Activity) (bool, error) {
	result := backend.db.Table("Activities").Create(&activity)
	if err := result.Error; err != nil {
		return false, err
	}
	fmt.Println("Activity saved in db")
	return true, nil
}
