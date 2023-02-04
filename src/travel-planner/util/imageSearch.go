package util

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Response struct {
	Results []struct {
		PlaceID string `json:"place_id"`
		Photos  []struct {
			PhotoReference string `json:"photo_reference"`
		}
	} `json:"results"`
	Status string `json:"status"`
}

func GetImageURL(placeName string) string {
	//apiKey := "AIzaSyAh7k8l1eosOYKu9PwWfQ1NIZ3cuCLXFuM"
	apiKey := "AIzaSyBeDsdP8vyzXFVPg75Tmb7fU51KvPGOGbk"

	//placeName := "Empire State Building"
	placeName = strings.ReplaceAll(placeName, " ", "%")
	url := "https://maps.googleapis.com/maps/api/place/textsearch/json?query=" +
		placeName +
		"%in%New%York&&key=AIzaSyAh7k8l1eosOYKu9PwWfQ1NIZ3cuCLXFuM"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return ""
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer res.Body.Close()

	var place Response
	decode := json.NewDecoder(res.Body)
	err = decode.Decode(&place)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	if len(place.Results) < 1 {
		return ""
	}

	photoReference := place.Results[0].Photos[0].PhotoReference
	fmt.Println("Image url: ")
	imageURL := fmt.Sprintf("https://maps.googleapis.com/maps/api/place/photo?maxwidth=1600&photoreference=%s&key=%s", photoReference, apiKey)
	fmt.Println(imageURL)
	return imageURL
}
