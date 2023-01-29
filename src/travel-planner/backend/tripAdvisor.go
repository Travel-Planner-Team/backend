package backend

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	//"strings"
	"travel-planner/constants"
	"travel-planner/model"
)



func SearchDetailFromTrip(sites []model.Site) {
    for _, item := range sites {
        location := GetSearchTripAdvisor(item.Site_name)
		
		var location_id string
		location_id = location.Location_id
		tripDetails := GetDetailTripAdvisor(location_id)
		// if err != nil {
        //     return false,err
		// }
		item.Site_name = tripDetails.Name
		item.Rating = tripDetails.Rating
		item.Phone_number = tripDetails.Phone
		item.Description = tripDetails.Description
		item.Address = tripDetails.Address_string
		fmt.Printf(item.Site_name)
	}

}


func GetSearchTripAdvisor(name string)(model.TripSite){
	key := constants.TRIPADVISOR_API_KEY
	
	//url := "https://api.content.tripadvisor.com/api/v1/location/search?key=62A808FFA5BB43458AA517B597F7C0E1&searchQuery=hi&language=en"
	url := fmt.Sprintf("https://api.content.tripadvisor.com/api/v1/location/%s/search?key=%s&searchQuery=%s&language=en", key, name)
	fmt.Println(url)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
     
    
	fmt.Println(res)
	fmt.Println(string(body))
	// convert response string to json 
	var tripSites []model.TripSite
	json.Unmarshal([]byte(body), &tripSites)
	//??
//    if tripSites == nil {
// 	 return nil, errors.New("unable to find sites in tripadvisor")
//    }
   return tripSites[0]
}

func GetDetailTripAdvisor(location_id string) (model.TripDetails) {
	key := constants.TRIPADVISOR_API_KEY
	url := fmt.Sprintf("https://api.content.tripadvisor.com/api/v1/location/%s/details?language=en&currency=USD&key=%s", location_id, key)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
    
	fmt.Println(res)
	fmt.Println(string(body))

	var tripDetails model.TripDetails
	json.Unmarshal([]byte(body), &tripDetails)
    
    return tripDetails
}

