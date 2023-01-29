package service

import (
	"fmt"
	"travel-planner/backend"
	"travel-planner/model"

	//"travel-planner/constants"
	"errors"

	//"golang.org/x/tools/go/analysis/passes/nilness"
)

func CheckUser (userEmail string, password string) (bool, error){
 user, err := backend.DB.ReadUserByEmail(userEmail);

 if err != nil {
	return false, err
 }
 if user.Password == password {
	return true, nil
 }
 return  false, nil
}

func AddUser (user *model.User) (bool, error){
   success, err := backend.DB.SaveUser(user)
   if !success {
	fmt.Println("Failed to save in db")
	return false, err
   }

   return true, nil

}

func CheckUserInfo (userID  uint32) (*model.User, error){
     user, err := backend.DB.ReadUserById(userID)
	 if err != nil {
		return nil, err
	 }

	 if user == nil{
		return nil, errors.New("unable to find app in db")
	 }

	 return user, nil

}

func UpdateUserInfo(id uint32, password, username, gender string, age int64)(bool, error){
	fmt.Println("updateuser")
	
    success, err := backend.DB.UpdateInfo(id, password, username,gender, age)
	
	if err != nil {
		return false, err
	}

	return success, nil
}