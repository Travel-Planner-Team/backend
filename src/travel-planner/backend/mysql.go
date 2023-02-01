package backend

import (
	"fmt"
	"travel-planner/constants"
	"travel-planner/model"
	"travel-planner/util"

	//"travel_planner/handler"

	"errors"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *MySQLBackend
)

type MySQLBackend struct {
	db *gorm.DB
}

func InitMySQLBackend(config *util.MySQLInfo) {

	endpoint, username, password := config.Endpoint, config.Username, config.Password

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
		username, password, endpoint, constants.MYSQL_DBNAME, constants.MYSQL_CONFIG)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = &MySQLBackend{db}
}

func (backend *MySQLBackend) ExampleQueryFunc() error {
	var users []model.User
	result := backend.db.Table("Users").Find(&users)
	fmt.Println(users)
	fmt.Println(result.RowsAffected)
	return result.Error

}

func (backend *MySQLBackend) ReadUserByEmail(userEmail string) (*model.User, error) {
	var user model.User
	result := backend.db.Table("Users").Where("email = ?", userEmail).Find(&user)
	if err := result.Error; err != nil {
		return nil, err
	}
	if result.RowsAffected != 0 {
		return &user, nil
	}

	return nil, errors.New("The email has not been registed before.")
}

func (backend *MySQLBackend) ReadUserById(userId uint32) (*model.User, error) {
	var user model.User
	result := backend.db.Table("Users").First(&user, userId)
	if err := result.Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// update interface has no return value in gorm?
func (backend *MySQLBackend) UpdateInfo(id uint32, password, username, gender string, age int64) (bool, error) {
	var user model.User
	result := backend.db.Table("Users").First(&user, id)

	if result.Error != nil {
		fmt.Printf("error for update in db %v\n", result.Error)
		return false, result.Error
	}
	fmt.Printf("userID:%v\n", user.Id)
	fmt.Println(age)
	backend.db.Table("Users").Model(&user).Select("Password", "Username", "Gender", "Age").
		Updates(model.User{Password: password, Username: username, Gender: gender, Age: age})
	fmt.Printf("usersAge:%v\n", user.Age)
	return true, nil
}

func (backend *MySQLBackend) GetSitesInVacation(vacationId uint32) ([]model.Site, error) {
	var sites []model.Site
	result := backend.db.Table("Sites").Where("vacation_id = ?", vacationId).Find(&sites)
	if result.Error != nil {
		fmt.Println("Failed to get sites from db")
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		fmt.Printf("No sites record in vacation %v\n", vacationId)
		return nil, nil
	}
	return sites, nil
}

func (backend *MySQLBackend) GetVacations() ([]model.Vacation, error) {
	var vacations []model.Vacation
	result := backend.db.Table("Vacations").Find(&vacations)
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

func (backend *MySQLBackend) GetActivityFromPlanId(planId uint32) ([]model.Activity, error) {
	var activities []model.Activity
	result := backend.db.Table("Activities").Where("plan_id = ?", planId).Find(&activities)
	fmt.Print(activities, result)
	if result.Error != nil {
		return nil, result.Error
	}
	return activities, nil
}

func (backend *MySQLBackend) GetTransportationFromPlanId(planId uint32) ([]model.Transportaion, error) {
	var transportations []model.Transportaion
	result := backend.db.Table("Transportations").Where("plan_id = ?", planId).Find(&transportations)
	fmt.Print(transportations, result)
	if result.Error != nil {
		return nil, result.Error
	}
	return transportations, nil
}

func (backend *MySQLBackend) GetSiteFromSiteId(siteId uint32) (*model.Site, error) {
	var site *model.Site
	result := backend.db.Table("Sites").Where("id = ?", siteId).Find(&site)
	fmt.Print(site, result)
	if result.Error != nil {
		return  nil, result.Error
	}
	return site, nil
}

func (backend *MySQLBackend) GetPlanFromVacationId(vacationId uint32) ([]model.Plan, error) {
	var plans []model.Plan
	result := backend.db.Table("Plans").Where("vacation_id = ?", vacationId).Find(&plans)
	fmt.Print(plans, result)
	if result.Error != nil {
		return nil, result.Error
	}
	return plans, nil
}

func (backend *MySQLBackend) SaveSites(sites []model.Site) (bool, error) {
	var count = 0
	for _, item := range sites {
		result := backend.db.Table("Sites").Create(&item)

		if result.Error != nil || result.RowsAffected == 0 {
			fmt.Printf("Faild to save site %v\n", item.Site_name)
		}
		count++
	}
	if count == 0 {
		return false, errors.New("Failed to save all the sites")
	}
	return true, nil
}

func (backend *MySQLBackend) SaveSingleSite(site model.Site) (bool, error) {

	result := backend.db.Table("Sites").Create(&site)

	if result.Error != nil || result.RowsAffected == 0 {
		fmt.Printf("Faild to save site %v\n", site.Site_name)
	}

	return true, nil
}

func (backend *MySQLBackend) SaveVacationPlanToSQL(plan model.Plan) (error) {
	result := backend.db.Table("Plans").Create(&plan)
	if result.Error != nil || result.RowsAffected == 0 {
		fmt.Printf("Faild to save plan %v\n", plan.Id)
	}
	return nil
}

func (backend *MySQLBackend) AddVacationIdToSite(siteID uint32, vacationID string)(bool, error){
	var site model.Site
	result := backend.db.Table("Sites").First(&site, siteID)

	if result.Error != nil{
		fmt.Printf("error for update in db %v\n",result.Error)
		return false, result.Error
	}
	fmt.Printf("siteID:%v\n", siteID)
	fmt.Printf("vacationID:%v\n", vacationID)
    backend.db.Table("Sites").Model(&site).Select("vacation_id").Updates(model.Site{Vacation_id: vacationID})

    return true, nil
}
