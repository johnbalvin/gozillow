package gozillow

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Test() {
	test4()
}
func test1() {
	coords := CoordinatesInput{
		Ne: CoordinatesValues{
			Latitude: 43.389689973008714,
			Longitud: -72.41662360750962,
		},
		Sw: CoordinatesValues{
			Latitude: 40.07030481778054,
			Longitud: -78.84362556063462,
		},
	}
	zomValue := 3
	results, err := SearchAll(zomValue, coords, nil)
	if err != nil {
		log.Println("err: ", err)
		return
	}
	f, _ := os.Create("./results.json")
	json.NewEncoder(f).Encode(results)
}

func test2() {
	propertyURL := "https://www.zillow.com/homedetails/858-Shady-Grove-Ln-Harrah-OK-73045/339897685_zpid/"
	results, err := DetailsFromPropertyURL(propertyURL, nil)
	if err != nil {
		log.Println("err: ", err)
	}
	f, _ := os.Create("./details.json")
	json.NewEncoder(f).Encode(results)
}

func test3() {
	propertyID := int64(2056016566)
	results, err := DetailsFromPropertyID(propertyID, nil)
	if err != nil {
		log.Println("err: ", err)
	}
	f, _ := os.Create("./details.json")
	json.NewEncoder(f).Encode(results)
}

func test4() {
	coords := CoordinatesInput{
		Ne: CoordinatesValues{
			Latitude: 47.76725314073866,
			Longitud: -122.15539952490977,
		},
		Sw: CoordinatesValues{
			Latitude: 47.67128302452179,
			Longitud: -122.3442270395582,
		},
	}
	zoomValue := 2
	results, err := SearchFirstPage(zoomValue, coords, nil)
	if err != nil {
		log.Println("searchHanlder 6 err: ", err)
		return
	}

	rawJSON, _ := json.MarshalIndent(results, "", "  ")
	if err := os.WriteFile("./details.json", rawJSON, 0644); err != nil {
		log.Println(err)
		return
	}

	for i, result := range results {
		propertyID, err := strconv.ParseInt(result.Zpid, 10, 64)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Printf("propertyID: %d\n", propertyID)
		_, err = DetailsFromPropertyID(propertyID, nil)
		if err != nil {
			fmt.Println("Error retrieving property details:", err)
			return
		}
		log.Printf("Progress: %d/%d\n", i+1, len(results))
	}
}
