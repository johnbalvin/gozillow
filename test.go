package gozillow

import (
	"encoding/json"
	"log"
	"os"

	"github.com/johnbalvin/gozillow/search"
)

func Test() {
	test1()
}
func test1() {
	coords := CoordinatesInput{
		Ne: search.CoordinatesValues{
			Latitude: 43.389689973008714,
			Longitud: -72.41662360750962,
		},
		Sw: search.CoordinatesValues{
			Latitude: 40.07030481778054,
			Longitud: -78.84362556063462,
		},
	}
	zomValue := 3
	results, err := SearchAll(zomValue, coords, nil)
	if err != nil {
		log.Println("err: ", err)
	}
	f, _ := os.Create("./results.json")
	json.NewEncoder(f).Encode(results)
}

/*
func test2() {
	propertyURL := ""
	results, err := GetFromPropertyURL(propertyURL, nil)
	if err != nil {
		log.Println("err: ", err)
	}
	f, _ := os.Create("./results.json")
	json.NewEncoder(f).Encode(results)
}

func test3() {
	propertyURL := ""
	results, err := GetFromPropertyID(propertyURL, nil)
	if err != nil {
		log.Println("err: ", err)
	}
	f, _ := os.Create("./results.json")
	json.NewEncoder(f).Encode(results)
}
*/
