package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
	"travel-planner/model"
	"travel-planner/service"

	"github.com/pborman/uuid"
)

func GetVacationsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request: /vacation")
	w.Header().Set("Content-Type", "application/json")

	vacations, err := service.GetVacationsInfo()
	if err != nil {
		http.Error(w, "Fail to read vacation info from backend", http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(vacations)
	if err != nil {
		http.Error(w, "Fail to parse vacations list into JSON", http.StatusInternalServerError)
		return
	}
	w.Write(js)
}

func SaveVacationsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request: /vacation/init")
	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
	var vacation model.Vacation
	fmt.Println(r.Body)
	if err := decoder.Decode(&vacation); err != nil {
		fmt.Println(err)
		http.Error(w, "Cannot decode vacation input", http.StatusBadRequest)
		return
	}

	vacation.Id = uuid.New()
	success, err := service.AddVacation(&vacation)
	if err != nil || !success {
		fmt.Println(err)
		http.Error(w, "Unable to save", http.StatusInternalServerError)
	}

	js, err := json.Marshal(vacation)
	if err != nil {
		http.Error(w, "Fail to save vacation into DB", http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Vacation saved: " + fmt.Sprint(vacation.Id)))
	w.Write(js)
}

func InitPlanHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Plan had been init"))
}

// func MakeRouteForVacation(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Received request: /vacation/{vacation_id}/plan/routes")
// 	w.Write([]byte("Potential Routes Sent"))
// }

func GetVacationPlanHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request: /vacation/{vacation_id}/plan")
	vacationID := r.URL.Query().Get("vacation_id")
	fmt.Printf("vacationID: %v\n", vacationID)
	w.Header().Set("Content_Type", "application/json")
	// get plans
	intId ,_:=strconv.ParseInt(vacationID, 0, 64)
	pasedId := uint32(intId)
	plans,err := service.GetPlanInfoFromVactionId(pasedId)
	// Marshal the activities to JSON
	jsonData, err := json.Marshal(plans)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Write the JSON data to the response
	w.Write(jsonData)
	
	// plan details： activities + transportations
	// get each slice of plans
	for i := 0; i < len(plans); i++ {
		plan := &plans[i]
		fmt.Println(plan.Id)
		intId ,_:=strconv.ParseInt(plan.Id, 0, 64)
		pasedId := uint32(intId)

		// get []activities
		activities,err := service.GetActivitiesInfoFromPlanId(pasedId)
		jsonData, err := json.Marshal(activities)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(jsonData)

		// get []transportations
		transportations,err := service.GetTransportationFromPlanId(pasedId)
		jsonData, err = json.Marshal(transportations)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(jsonData)
	}
	// Create a slice of activities
	// activities := []model.Activity{
	// 	{Id: 1, StartTime: time.Now(), EndTime: time.Now().Add(time.Hour), Date: time.Now(), Duration: 3600, Site_id: 100},
	// 	{Id: 2, StartTime: time.Now().Add(time.Hour * 2), EndTime: time.Now().Add(time.Hour * 3), Date: time.Now(), Duration: 3600, Site_id: 200},
	// 	{Id: 3, StartTime: time.Now().Add(time.Hour * 4), EndTime: time.Now().Add(time.Hour * 5), Date: time.Now(), Duration: 3600, Site_id: 300},
	// }
}

func SaveActivitiesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request: /vacation/{vacation_id}/plan/{plan_id}/save")
	vacationId := r.Context().Value("vacation_id")
	plan_id := r.Context().Value("plan_id")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Activities saved: " + fmt.Sprint(vacationId) + fmt.Sprint(plan_id)))
}

func InitVacationPlanHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request: /vacation/{vacation_id}/plan/init")
	// Decode the request body into a Plan struct
	dateString := r.FormValue("start_date")
	startDate, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	duration, err := strconv.ParseInt(r.FormValue("duration"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	vacation_id := r.FormValue("vacation_id")
	plan := model.Plan{
		Id:            uuid.New(),
		Start_date:    startDate,
		Duration_days: duration,
		VacationId:    vacation_id,
	}
	// err := json.NewDecoder(r.Body).Decode(&plan)
	// if err != nil {
	// 	http.Error(w, "Error decoding request body", http.StatusBadRequest)
	// 	return
	// }

	// Save the plan to the database
	err = service.SaveVacationPlan(plan)
	if err != nil {
		http.Error(w, "Error saving plan to database", http.StatusInternalServerError)
		return
	}

	// Marshal the plan to JSON
	jsonData, err := json.Marshal(plan)
	if err != nil {
		http.Error(w, "Error marshalling JSON", http.StatusInternalServerError)
		return
	}

	// Set the content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON data to the response
	w.Write(jsonData)
}

func MakeRouteForVacation(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request: /vacation/{vacation_id}/plan/routes")
	w.Write([]byte("Potential Routes Sent"))
}
