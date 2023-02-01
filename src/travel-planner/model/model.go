package model

import (
	// "fmt"
	"regexp"
	"strings"
	"time"
	"travel-planner/util/errors"
	// "gorm.io/gorm"
)

type AppStub struct {
	Id          string `json:"id"`
	User        string `json:"user"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Url         string `json:"url"`
	ProductID   string `json:"product_id"`
	PriceID     string `json:"price_id"`
}

type UserStub struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Age      int64  `json:"age"`
	Gender   string `json:"gender"`
}

type Vacation struct {
	Id           string    `json:"id"`
	Destination  string    `json:"destination"`
	StartDate    time.Time `json:"start_date"`
	EndDate      time.Time `json:"end_date"`
	DurationDays int64     `json:"duration_days"`
	UserId       uint32    `json:"user_id"`
}

//	type Model struct {
//		ID uint `jason:"id"` // `gorm:"primary_key jason:"id"`
//		CreatedAt   time.Time  `json:"created_at"`
//		UpdatedAt   time.Time  `json:"updated_at"`
//		DeletedAt   *time.Time `json:"deleted_at"`
//	}
type User struct {
	Id       uint32 `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
	Age      int64  `json:"age"`
	Gender   string `json:"gender"`
}

type Site struct {
	Id           uint32 `json:"id"`
	Site_name    string `json:"site_name"`
	Rating       string `json:"rating"`
	Phone_number string `json:"phone_number"`
	Vacation_id  string `json:"vacation_id"`
	Description  string `json:"description"`
	Address      string `json:"address"`
	Url          string `json:"site_url`
}

type TripSite struct {
	Location_id string      `json:"location_id"`
	Name        string      `json:"name"`
	Address_obj Address_obj `json:"address_obj"`
}

type Address_obj struct {
	Street1        string `json:"street1"`
	Street2        string `json:"street2"`
	City           string `json:"city"`
	State          string `json:"state"`
	Country        string `json:"country"`
	Postalcode     string `json:"postalcode"`
	Address_string string `json:"address_string"`
}

type TripDetails struct {
	Location_id    string `json:"location_id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	Web_url        string `json:"web_url"`
	Address_string string `json:"address_string"`
	Rating         string `json:"rating"`
	Phone          string `json:"phone"`
}

type Plan struct {
	Id            string `json:"id"`
	Start_date    string `json:"start_date"`
	Duration_days int64  `json:"duration_days"`
	VacationId    string `json:"VacationId"`
}

type Activity struct {
	Id        uint32 `json:"id"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Date      string `json:"date"`
	Duration  int64  `json:"duration"`
	Site_id   uint32 `json:"site_id"`
}

type Transportaion struct {
	Id        uint32 `json:"id"`
	Type      string `json:"type"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Date      string `json:"date"`
}

type ActivitiesList struct {
	ActivityID            int    `json:"activity_id"`
	ActivityName          string `json:"activity_name"`
	ActivityType          string `json:"activity_type"`
	ActivityDescription   string `json:"activity_description"`
	ActivityAddress       string `json:"activity_address"`
	ActivityPhone         string `json:"activity_phone"`
	ActivityWebsite       string `json:"activity_website"`
	ActivityImage         string `json:"activity_image"`
	ActivityStartDatetime string `json:"activity_start_datetime"`
	ActivityEndDatetime   string `json:"activity_end_datetime"`
	ActivityDate          string `json:"activity_date`
	ActivityDuration      int64  `json:"activity_duration"`
}

type DaysInfo struct {
	DayIDX int              `json:"day_idx"`
	Act    []ActivitiesList `json:"activities"`
	Trans  []Transportaion  `json:"transportation"`
}

type PlansInfo struct {
	PlanIDX int      `json:"plan_idx"`
	Days    DaysInfo `json:"days"`
}

type PlanDetail struct {
	VacationID int         `json:"vacation_id"`
	Plans      []PlansInfo `json:"plans"`
}

func (user *User) Validate() *errors.RestErr {
	user.Username = strings.TrimSpace(user.Username)
	user.Password = strings.TrimSpace(user.Password)
	user.Email = strings.TrimSpace(user.Email)
	if user.Email == "" {
		return errors.NewBadRequestError("Invalid email address")
	}
	if user.Username == "" || regexp.MustCompile(`^[a-z0-9]$`).MatchString(user.Username) {
		return errors.NewBadRequestError("Invalid username")
	}
	if user.Password == "" {
		return errors.NewBadRequestError("Invalid password")
	}
	return nil
}
