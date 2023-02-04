package service

import (
	"fmt"
	"travel-planner/backend"
	"travel-planner/model"
	"travel-planner/util/errors"
	//"travel-planner/constants"
	// "errors"
	//"golang.org/x/tools/go/analysis/passes/nilness"
)

func CreateUser(user *model.User) (bool, *errors.RestErr) {
	// username existed?
	success, err := backend.DB.FindUser(user)
	if err != nil {
		return false, errors.NewBadRequestError("The database has error")
	}
	if !success {
		return false, errors.NewBadRequestError("The user has already exist")
	}

	// save to db

	success, err = backend.DB.SaveUser(user)
	if err != nil {
		return false, errors.NewInternalServerError("Failed to create user")
	}
	// fmt.Printf("User is added: %s\n", user.Username)
	return true, nil
}

func CheckUser(userEmail string, password string) (bool, string, error) {
	user, err := backend.DB.ReadUserByEmail(userEmail)

	if err != nil {
		return false, "", err
	}

	if user.Password == password {
		return true, fmt.Sprintf("%d", user.Id), nil
	}
	return false, "", nil
}

func CheckUserInfo(userID uint32) (*model.User, *errors.RestErr) {
	user, err := backend.DB.ReadUserById(userID)
	if err != nil {
		return nil, errors.NewInternalServerError("Failed to load user")
	}

	if user == nil {
		return nil, errors.NewBadRequestError("unable to find app in db")
	}

	return user, nil

}

func UpdateUserInfo(id uint32, password, username, gender string, age int64) (bool, error) {
	fmt.Println("updateuser")

	success, err := backend.DB.UpdateInfo(id, password, username, gender, age)

	if err != nil {
		return false, err
	}

	return success, nil
}
