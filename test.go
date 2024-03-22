package gozillow

import (
	"encoding/json"
	"log"
	"os"
)

func Test() {
	test3()
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
