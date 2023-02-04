package backend

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"travel-planner/constants"
	"travel-planner/model"
	"travel-planner/util"
)

func SearchDetailFromTrip(sites []model.Site) {
	for key, item := range sites {
		fmt.Printf("Sitename:%v\n", item.SiteName)
		location := GetSearchTripAdvisor(item.SiteName)
		if location == nil {
			continue
		}
		location_id := location.LocationId

		res := GetDetailsWithLocationId(location_id)

		if res == "" {
			continue
		}

		resBytes := []byte(res)                // Converting the string "res" into byte array
		var jsonRes map[string]interface{}     // declaring a map for key names as string and values as interface
		_ = json.Unmarshal(resBytes, &jsonRes) // Unmarshalling

		if jsonRes == nil {
			continue
		}

		if jsonRes["description"] != nil {
			item.Description = jsonRes["description"].(string)
		}

		if jsonRes["phone"] != nil {
			item.PhoneNumber = jsonRes["phone"].(string)
		}

		if jsonRes["rating"] != nil {
			item.Rating = jsonRes["rating"].(string)
		}

		if jsonRes["address_obj"] != nil {
			details_Address := jsonRes["address_obj"].(map[string]interface{})
			item.Address = details_Address["address_string"].(string)
		}

		if jsonRes["latitude"] != nil {
			var l = jsonRes["latitude"].(string)
			value, _ := strconv.ParseFloat(l, 32)
			item.Latitude = float32(value)
		}

		if jsonRes["longitude"] != nil {
			var l = jsonRes["longitude"].(string)
			value, _ := strconv.ParseFloat(l, 32)
			item.Longitude = float32(value)
		}

		item.ImageUrl = util.GetImageURL(item.SiteName)

		fmt.Println(item)
		//DB.SaveSingleSite(item)
		sites[key] = item
	}
}
func GetSearchTripAdvisor(name string) *model.TripSite {
	url := getUrl(name)
	fmt.Println(url)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)

	resBytes := []byte(body)               // Converting the string "res" into byte array
	var jsonRes map[string]interface{}     // declaring a map for key names as string and values as interface
	_ = json.Unmarshal(resBytes, &jsonRes) // Unmarshalling

	data := jsonRes["data"].([]interface{})
	if data == nil {
		return nil
	}

	firstData := data[0]
	if firstData == nil {
		return nil
	}

	fmt.Print(firstData)
	firstDataJson, _ := json.Marshal(firstData)
	var tripSites model.TripSite
	json.Unmarshal(firstDataJson, &tripSites)

	return &tripSites
}

func GetDetailTripAdvisor(location_id string) model.TripDetails {
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

func GetDetailsWithLocationId(location_id string) string {
	key := constants.TRIPADVISOR_API_KEY
	url := fmt.Sprintf("https://api.content.tripadvisor.com/api/v1/location/%s/details?language=en&currency=USD&key=%s", location_id, key)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")

	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode == 404 {
		return ""
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
	return string(body)

}

func getUrl(name string) string {
	key := constants.TRIPADVISOR_API_KEY
	baseUrl := "https://api.content.tripadvisor.com/api/v1/location/search?"
	baseUrl += "key=" + key
	baseUrl += "&searchQuery=" + url.QueryEscape(name)
	baseUrl += "&language=en"
	return baseUrl
}
