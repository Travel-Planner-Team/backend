package backend

import (
	"fmt"
	"travel-planner/constants"
	"travel-planner/model"
	"travel-planner/util"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *MySQLBackend

type MySQLBackend struct {
	db *gorm.DB
}

func InitMySQLBackend(config *util.MySQLInfo) {
	endpoint, username, password := config.Endpoint, config.Password, config.Username
	dsn := username + ":" + password + "@tcp(" + endpoint + ")/" + constants.MYSQL_DBNAME + "?" + constants.MYSQL_CONFIG
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = &MySQLBackend{db}
}

func (backend *MySQLBackend) exampleQueryFunc() error {
	var users []model.User
	result := backend.db.Table("Users").Find(&users)
	fmt.Println(users)
	fmt.Println(result.RowsAffected)
	return result.Error
}
func (backend *MySQLBackend) ReadUserByEmail(userEmail string) (*model.User, error) {
	var user model.User
	result := backend.db.Table("Users").Where("email = ?",userEmail).Find(&user)
	if err := result.Error; err != nil {
		return nil, err
	}
	fmt.Println("User find in db")
	return &user, nil
}
func (backend *MySQLBackend) SaveUser (user *model.User) (bool, error) {
	result := backend.db.Create(&user)
	if err := result.Error; err != nil{
		return false, err
	}
	fmt.Println("User saved in db")
	return true, nil
}

func (backend *MySQLBackend) ReadUserById (userId string)(*model.User, error){
	var user model.User
	result := backend.db.Table("Users").First(&user, "userId")
	if err := result.Error; err != nil {
		return nil, err
	}
	return &user, nil
}
// update interface has no return value in gorm?
func (backend *MySQLBackend) UpdateInfo (id, password, username,gender string, age int64)(bool, error){
	var user model.User
	result := backend.db.Table("Users").First(&user, "userId")
	if result.Error != nil{
		return false, result.Error
	}
    backend.db.Model(&user).Updates(model.User{Password: password, Username: username, Gender: gender, Age:age})
	
    return true, nil



}
	






