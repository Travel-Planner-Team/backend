package model

import (
	"time"
	//"gorm.io/gorm"
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
	Id       uint32    `json:"id"`
	Email    string  `json:"email"`
	Password string   `json:"password"`
	Username string `json:"username"`
	Age      int64  `json:"age"`
	Gender   string `json:"gender"`
}
// type Vacation struct {
// 	Id       string    `json:"id"`
// 	Destination_city    string  `json:"destication_city"`
// 	State_date string   `json:"state_date"`
// 	End_date string `json:"end_date"`
// 	Duration      int64  `json:"duration"`
// 	User_id   string `json:"user_id"`
// }

type Site struct {
	Id       uint32    `json:"id"`
	Site_name    string  `json:"destication_city"`
	Rating string   `json:"rating"`
	Phone_number string `json:"phone_number"`
	Vacation_id   string `json:"vacation_id"`
	Description string `json:"description"`
	Address string `json:"address"`
}

type TripSite struct{
    Location_id string `json:"location_id"`
    Name string `json:"name"`
	Address_obj Address_obj `json:"address_obj"`

}
type Address_obj struct{
    Street1 string `json:"street1"`
    Street2 string `json:"street2"`
    City string `json:"city"`
    State string `json:"state"`
    Country string `json:"country"`
    Postalcode string `json:"postalcode"`
    Address_string string `json:"address_string"`
}

type TripDetails struct{
    Location_id string `json:"location_id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Web_url string `json:"web_url"`
	Address_string string `json:"address_string"`
    Rating string `json:"rating"`
	Phone string `json:"phone"`
	Latitude uint32 `json:"latitude"`
	Longitude uint32 `json:"longitude"`
}

type Activity struct {
	Activity_id   uint32 `json:"activity_id"`
	Start_time    time.Time `json:"start_time"`
	End_time      time.Time `json:"end_time"`
	Date 		  time.Time `json:"start_date"`
	Duration      time.Duration `json:"duration"`
	Site_id   	  uint32 `json:"site_id"`
	Plan_id       uint32 `json:"plan_id"`
}

type Transportation struct {
	Transportation_id   uint32 `json:"transportation_id"`
	Type          string    `json:"transportation_type"`
	Start_time    time.Time `json:"start_time"`
	End_time      time.Time `json:"end_time"`
	Date 		  time.Time `json:"start_date"`
	Duration      time.Duration `json:"duration"`
	Plan_id       uint32    `json:"plan_id"`
}

type Plan struct {
	Plan_id       uint32    `json:"plan_id"`
	StartDate     int	`json:"startdate"`
	Vacation_id   uint32 `json:"vacation_id"`
}
