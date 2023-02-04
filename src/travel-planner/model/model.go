package model

import (
	// "fmt"
	"regexp"
	"strings"
	"time"

	"travel-planner/util/errors"
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
	Id           uint32    `json:"id"`
	Destination  string    `json:"destination"`
	StartDate    time.Time `json:"start_date"`
	EndDate      time.Time `json:"end_date"`
	DurationDays int64     `json:"duration_days"`
	UserId       uint32    `json:"user_id"`
}

type User struct {
	Id       uint32 `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
	Age      int64  `json:"age"`
	Gender   string `json:"gender"`
}

type Site struct {
	Id          uint32  `json:"id"`
	SiteName    string  `json:"site_name"`
	Rating      string  `json:"rating"`
	PhoneNumber string  `json:"phone_number"`
	VacationId  uint32  `json:"vacation_id"`
	Description string  `json:"description"`
	Address     string  `json:"address"`
	Latitude    float32 `json:"latitude"`
	Longitude   float32 `json:"longitude"`
	SiteUrl     string  `json:"site_url"`
	ImageUrl    string  `json:"image_url"`
}

type TripSite struct {
	LocationId string     `json:"location_id"`
	Name       string     `json:"name"`
	Address    AddressObj `json:"address_obj"`
}

type AddressObj struct {
	Street1       string `json:"street1"`
	Street2       string `json:"street2"`
	City          string `json:"city"`
	State         string `json:"state"`
	Country       string `json:"country"`
	Postalcode    string `json:"postalcode"`
	AddressString string `json:"address_string"`
}

type TripDetails struct {
	LocationId    string `json:"location_id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	WebUrl        string `json:"web_url"`
	AddressString string `json:"address_string"`
	Rating        string `json:"rating"`
	Phone         string `json:"phone"`
}

type Plan struct {
	Id           uint32    `json:"id"`
	StartDate    time.Time `json:"start_date"`
	DurationDays int64     `json:"duration_days"`
	VacationId   uint32    `json:"vacation_id"`
}

type Activity struct {
	Id          uint32    `json:"id"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	Date        time.Time `json:"date"`
	DurationHrs float32   `json:"duration_hrs"`
	SiteId      uint32    `json:"site_id"`
	PlanId      uint32    `json:"plan_id"`
}

type Transportation struct {
	Id          uint32    `json:"id"`
	Type        int       `json:"type"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	Date        time.Time `json:"date"`
	DurationHrs float32   `json:"duration_hrs"`
	PlanId      uint32    `json:"plan_id"`
}

type SavePlanRequestBody struct {
	ActivityInfoList       []Activity       `json:"activity_info_list"`
	TransportationInfoList []Transportation `json:"transportation_info_list"`
}

type ShowPlan struct {
	Plan_id                uint32           `json:"planid"`
	ActivityInfoList       []Activity       `json:"activity_info_list"`
	TransportationInfoList []Transportation `json:"transportation_info_list"`
}

type ListOfShowPlan struct {
	ShowPlans   []ShowPlan `json:"show_plan_list"`
	Vacation_id string     `json:"vacation_id"`
}

type ActivitiesList struct {
	ActivityID            uint32    `json:"activity_id"`
	ActivityName          string    `json:"activity_name"`
	ActivityType          string    `json:"activity_type"`
	ActivityDescription   string    `json:"activity_description"`
	ActivityAddress       string    `json:"activity_address"`
	ActivityPhone         string    `json:"activity_phone"`
	ActivityWebsite       string    `json:"activity_website"`
	ActivityImage         string    `json:"activity_image"`
	ActivityStartDatetime time.Time `json:"activity_start_datetime"`
	ActivityEndDatetime   time.Time `json:"activity_end_datetime"`
	ActivityDate          time.Time `json:"activity_date"`
	ActivityDuration      float32   `json:"activity_duration"`
	ActivityStartAddress  string    `json:"activity_start_address"`
	ActivityEndAddress    string    `json:"activity_end_address"`
	ActivityLongitude     float32   `json:"activity_longitude"`
	ActivityLatitude      float32   `json:"activity_latitude"`
}

type TransportationList struct {
	TransportationId           uint32    `json:"transportation_id"`
	TransportationType         int       `json:"transportation_type"`
	TransportationStartTime    time.Time `json:"transportation_start_time"`
	TransportationEndTime      time.Time `json:"transportation_end_time"`
	TransportationDate         time.Time `json:"transportation_date"`
	TransportationDurationHrs  float32   `json:"transportation_duration_hrs"`
	TransportationPlanId       uint32    `json:"transportation_plan_id"`
	TransportationStartAddress string    `json:"transportation_start_address"`
	TransportationEndAddress   string    `json:"transportation_end_address"`
}

type DayInfo struct {
	DayIDX int                  `json:"day_idx"`
	Act    []ActivitiesList     `json:"activities"`
	// Trans  []TransportationList `json:"transportation"`
}

type PlansInfo struct {
	PlanIDX uint32    `json:"plan_idx"`
	Days    []DayInfo `json:"days"`
}

type PlanDetail struct {
	VacationID uint32      `json:"vacation_id"`
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
