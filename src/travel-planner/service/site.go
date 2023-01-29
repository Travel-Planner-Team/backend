package service

import (
	"fmt"
	"travel-planner/backend"
	"travel-planner/model"
  
)

func GetSitesList (vacationId uint32) ([]model.Site, error){
	var sites []model.Site
   sites, err := backend.DB.GetSitesInVacation(vacationId)
   // getSitesInVacation (vacationId string)
   return sites, err

}

func SearchSites( interest, city string )([]model.Site, error){
  var sites []model.Site
  var query string
  if interest == "" {
     query = "Find me a list of 10 top destination in " + city
  }else{
	   query = "Find me a list of 10 top " + interest + " in " + city
  }

  sites, err := backend.SearchSitesInChatGPT(query)

  if err != nil {
    fmt.Println("Failed to get information from ChatGPT")
    return nil, err
  }

  backend.SearchDetailFromTrip(sites)
  backend.DB.SaveSites(sites)

  return sites, nil

}
