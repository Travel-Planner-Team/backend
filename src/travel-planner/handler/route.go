package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"travel-planner/model"
	"travel-planner/service"
)


func GetPlanHandler(w http.ResponseWriter, r *http.Request){
	fmt.Println("Received a get sites request in the get all plans handler")
	w.Header().Set("Content-Type", "application/json")
   
	//line 66 is hardcode for test, we cannot get info from http yet, we should use line65 
	//vacationId := mux.Vars(r)["vacationid"]
   var	vacationId uint32 = 1
   var plans []model.Plan
   var err error
   plans, err = service.ShowRoute(vacationId)

   if err != nil || plans == nil {
	   http.Error(w, "Failed to get sites from bd", http.StatusInternalServerError)
	    return
   }

// change to json
     js, err := json.Marshal(plans)
    if err != nil{
    	http.Error(w, "Failed to parse sites to JSON format", http.StatusInternalServerError)
     }

    w.Write(js)
}