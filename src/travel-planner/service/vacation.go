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


