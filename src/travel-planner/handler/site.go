package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	//"regexp"
	//"strconv"
	"travel-planner/model"

	//"time"

	//"travel-planner/backend"
	"travel-planner/service"

	//"github.com/form3tech-oss/jwt-go"
	//"github.com/gorilla/mux"
	//"github.com/pborman/uuid"
)

func exampleHandler(w http.ResponseWriter, r *http.Request) {
	// Parse from body of request to get a json object.
	fmt.Println("Sample API handler")

	// user := r.Context().Value("user")
	// claims := user.(*jwt.Token).Claims
	// username := claims.(jwt.MapClaims)["username"]

	// app := model.App{
	// 	Id:          uuid.New(),
	// 	User:        username.(string),
	// 	Title:       r.FormValue("title"),
	// 	Description: r.FormValue("description"),
	// }

	// price, err := strconv.ParseFloat(r.FormValue("price"), 64)
	// fmt.Printf("%v,%T", price, price)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// app.Price = int(price * 100.0)

	// file, _, err := r.FormFile("media_file")
	// if err != nil {
	// 	http.Error(w, "Media file is not available", http.StatusBadRequest)
	// 	fmt.Printf("Media file is not available %v\n", err)
	// 	return
	// }

	// err = service.SaveApp(&app, file)
	// if err != nil {
	// 	http.Error(w, "Failed to save app to backend", http.StatusInternalServerError)
	// 	fmt.Printf("Failed to save app to backend %v\n", err)
	// 	return
	// }

	// fmt.Println("App is saved successfully.")
}

func GetSitesHandler(w http.ResponseWriter, r *http.Request){
	fmt.Println("Received a get sites request in the vacation")
	w.Header().Set("Content-Type", "application/json")
   
	//line 66 is hardcode for test, we cannot get info from http yet, we should use line65 
	//vacationId := mux.Vars(r)["vacationid"]
   var	vacationId uint32 = 1
   var sites []model.Site
   var err error
   sites, err = service.GetSitesList(vacationId)

   if err != nil || sites == nil {
	   http.Error(w, "Failed to get sites from bd", http.StatusInternalServerError)
	    return
   }

// change to json
     js, err := json.Marshal(sites)
    if err != nil{
    http.Error(w, "Failed to parse sites to JSON format", http.StatusInternalServerError)
     }

    w.Write(js)
}

//Search sites be send on query keywords in current vacation
func SearchSitesHandler(w http.ResponseWriter, r *http.Request){
	fmt.Println("Received a search sites request in vacation")
	w.Header().Set("Content-Type","application/json")

	//line 93 is hardcode for test, we cannot get info from http yet, we should use line65 
    //vacationId := mux.Vars(r)["vacationid"]
	
    interest := r.URL.Query().Get("interest")
	city := "New York"

	var sites []model.Site;
	sites, err := service.SearchSites(interest, city);

	if err != nil {
	http.Error(w, "Failed to search sites", http.StatusInternalServerError)
       return
	}

	js, err:= json.Marshal(sites)
	
	if err != nil {
       http.Error(w, "Failed to parse sites into JSON format", http.StatusInternalServerError)
       return
   }
   //返回
   w.Write(js)
}


	
	
	


