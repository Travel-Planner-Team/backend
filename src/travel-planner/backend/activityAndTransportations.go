package backend

import (
	"fmt"
	"travel-planner/model"
)

func (backend *MySQLBackend) GetActivityFromPlanId(plan_id uint32) ([]model.Activity, error) {
	var activities []model.Activity
	result := backend.db.Table("Activities").Where("plan_id = ?", plan_id).Order("start_time").Find(&activities)
	if result.Error != nil {
		return nil, result.Error
	}
	return activities, nil
}

func (backend *MySQLBackend) GetRoutes(sites []uint32) (int32, []model.Activity, []model.Transportation) {
	var activities []model.Activity
	result := backend.db.Table("Activities").Find(&activities)
	fmt.Println(activities, result)
	if result.Error != nil {
		return -1, nil, nil
	}
	var tranportations []model.Transportation
	result = backend.db.Table("Transportations").Find(&tranportations)
	fmt.Println(activities, result)
	if result.Error != nil {
		return -1, activities, nil
	}
	return 8, activities, tranportations
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
