package backend

import (
	"fmt"
	"travel-planner/model"
)

func (backend *MySQLBackend) GetVacations(userId uint32) ([]model.Vacation, error) {
	var vacations []model.Vacation
	result := backend.db.Table("Vacations").Where("user_id = ?", userId).Find(&vacations)
	fmt.Println(vacations, result)
	if result.Error != nil {
		return nil, result.Error
	}
	return vacations, nil
}

func (backend *MySQLBackend) SaveVacation(vacation *model.Vacation) (bool, error) {
	result := backend.db.Table("Vacations").Create(&vacation)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (backend *MySQLBackend) GetSingleVacation(Id uint32) (*model.Vacation, error) {
	var vacation model.Vacation

	result := backend.db.Table("Vacations").Where("id = ?", Id).Find(&vacation)
	fmt.Println(vacation, result)
	if result.Error != nil {
		fmt.Println("Failed to get vacation from db")
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		fmt.Printf("No vacation record in vacation %v\n", Id)
		return nil, nil
	}
	return &vacation, nil
}

func (backend *MySQLBackend) GetTransportationFromPlanId(planId uint32) ([]model.Transportation, error) {
	var transportations []model.Transportation
	result := backend.db.Table("Transportations").Where("plan_id = ?", planId).Order("start_time").Find(&transportations)
	if result.Error != nil {
		return nil, result.Error
	}
	return transportations, nil
}

func (backend *MySQLBackend) GetSiteFromSiteId(siteId uint32) (*model.Site, error) {
	var site *model.Site
	result := backend.db.Table("Sites").Where("id = ?", siteId).Find(&site)
	if result.Error != nil {
		return nil, result.Error
	}
	return site, nil
}
