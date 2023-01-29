package service

import (
	"fmt"
	"time"
	"travel-planner/backend"
	"travel-planner/model"

	"github.com/google/uuid"
)

func ShowRoute (vacationId uint32) ([]model.Plan, error) {
	// Step 1 
	// get all the site list from the backend using GetSitesInVacation
	var sites []model.Site
	var plans []model.Plan
	sites, err := backend.DB.GetSitesInVacation(vacationId)
	// if cannot get site, then just return empty plan
	if err != nil {
		return nil, nil
	}
	fmt.Println("Successfully get sites")
	fmt.Println(sites)

	// Step 2 
	// get the corresponding vacation by vacationId, because we need
	// the vacation Start Date, and End Date for later use to create Activity, and Transportation
	vacation, err := backend.DB.GetSingleVacation(vacationId)
	if err != nil {
		return nil, nil
	}
	fmt.Println("Successfully get the corresponding vacation")
	startDate := vacation.StartDate
	endDate := vacation.EndDate
	fmt.Println(startDate, endDate)


	// Step 3 
	// convert all the sites into Activity, save it to DB
	// call Google Distance Matrix API from a randomly choosing site to all the other site
	// convert it into Transportation, save it to DB
	// afterward, we save each day as one plan
	CreatePlan(sites, vacation)


	// Step 4
	// get all the plan from DB, and return
	plans, err = backend.DB.GetAllPlans(vacationId)
	if err != nil {
		return nil, nil
	}
	return plans, err
}

func CreatePlan (si []model.Site, vacation *model.Vacation) () {
	/*
		numberOfSite := len(sites)
		
		 below while condition why??? because we are updating startDate, and reducing the size of sites
		while (startDate.before(endDate) && numberOfSite > 1)
			
			Convert it into Activity
			save it to DB
			
			add the time by 2
			func (t Time) Add(d Duration) Time
				startDate = startDate.Add(time.Hour + 2)
			
			Google API
			pick the first site in the array	
			origin := site[0]
			destination := site [1:]
			call API
			use a for loop to find the shortest distance one (transportation method by public transportation)

			Convert it into Transportation, might need to change the properties in transportation
			Save it to DB

			check if the time is in certain time range
			if (it is after 12pm && before 2pm)
				fmt.println(lunch time 2 hrs)
				add 2 hrs
			
			
			if (it is after 8pm)
				save to new plan
				add to time to the next day 8pm
	*/
	// initialize all the variables we need
	// date
	startDate := vacation.StartDate
	endDate := vacation.EndDate

	// for planId
	var planId uint32
	dayCount := 1
	idCreatedCount := 0

	// site index
	i := 0
	for ; startDate.Before(endDate) && i < len(si); {
		if idCreatedCount < dayCount{
			planId = uuid.New().ID()
		}
		tempDate := startDate.AddDate(0, 0, 1)
		// Assume the user stay only 2 hrs in each site
		a := &model.Activity{
			Activity_id : uuid.New().ID(),
			Start_time: startDate,
			End_time: startDate.Add(time.Hour * 2),
			Date: startDate,
			Duration: 2,
			Site_id: si[i].Id,
			Plan_id: planId,
		}
		backend.DB.SaveActivity(a)
		startDate = startDate.Add(time.Hour * 2)
		// Save(activity)

		duration, in, err := backend.GetDistanceMatrix(si, i)
		if (err != nil) {
			fmt.Println("cannot get google api matrix")
			return
		}
		t := &model.Transportation {
			Transportation_id: uuid.New().ID(),
			Type: "public_transportation",
			Start_time: startDate,
			End_time: startDate.Add(duration),
			Date: startDate,
			Duration: duration,
			Plan_id: planId,
		}
		backend.DB.SaveTransportation(t)
		startDate = startDate.Add(duration)

		// since we already found the next stop, we swap the i + 1 and index, and continue
		temp := si[in]
		si[in] = si[i + 1]
		si[i + 1] = temp
		// update index !!
		i++;

		// check if it is after 8pm, if it is, that means new day, 
		// save plan, and set new time.
		if (startDate.Hour() > 20) {
			//It is a new plan, SavePlan
			p := &model.Plan{
				Plan_id: planId,
				StartDate: startDate.Day(),
				Vacation_id: vacation.Id,
			}
			backend.DB.SavePlan(p)
			//update our tracking parameter
			startDate = tempDate
			dayCount++;
		}

	}





}