package backend

import (
	"context"
	"log"
	"time"
	"travel-planner/constants"
	"travel-planner/model"

	"github.com/kr/pretty"
	"googlemaps.github.io/maps"
)

func GetDistanceMatrix(sites []model.Site, index int) (time.Duration, int, error){
	c, err := maps.NewClient(maps.WithAPIKey(constants.GOOGLE_API_KEY))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
		return time.Microsecond, -1, nil
	}
	origin := []string{sites[index].Address}

	var destination []string
	for i := index + 1; i < len(sites) - 1; i++ {
		s := sites[i].Address
		destination = append(destination, s)
	}

	r := &maps.DistanceMatrixRequest{
		Origins:      origin,
		Destinations: destination,
		Mode: "transit",
	}
	route, err := c.DistanceMatrix(context.Background(), r)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
		return time.Hour, -1, nil
	}
	pretty.Println(route)

	// after this we will get distancematrix response, 
	// we wanna pick the next closest site as our next point.
	// once we choose it, we can create a transportation object.
	// once create transportation, and return back to service level
	
	result, in := convertToTransportation(route)

	return result, in, err

}

func convertToTransportation(res *maps.DistanceMatrixResponse) (time.Duration, int){
	element := res.Rows[0].Elements
	// initialize globalMin
	globalMinDuration := element[0].Duration
	var index int
	for i := 0; i < len(element); i++ {
		du := element[i].Duration.Minutes()
		if (du <= globalMinDuration.Minutes()) {
			globalMinDuration = element[i].Duration
			index = i
		}
	}
	return globalMinDuration, index
}