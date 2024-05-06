package gozillow

import (
	"encoding/json"
	"log"
	"os"
)

func Test() {
	test1()
}
func test1() {
	coords := CoordinatesInput{
		Ne: CoordinatesValues{
			Latitude: 44.80185651702412,
			Longitud: -96.70265550613955,
		},
		Sw: CoordinatesValues{
			Latitude: 38.14378460775949,
			Longitud: -109.55665941238955,
		},
	}
	zomValue := 7
	pagination := 1
	listResults, mapResults, err := SearchSold(pagination, zomValue, coords, nil)
	if err != nil {
		log.Println("err: ", err)
		return
	}
	log.Println("len: ", len(listResults), len(mapResults))
	f, _ := os.Create("./listResults.json")
	json.NewEncoder(f).Encode(listResults)
	f2, _ := os.Create("./mapResult.json")
	json.NewEncoder(f2).Encode(mapResults)
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

/*
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
	listResults, mapResults, err := SearchForRent(1, zoomValue, coords, nil)
	if err != nil {
		log.Println("searchHanlder 6 err: ", err)
		return
	}

	rawJSON, _ := json.MarshalIndent(listResults, "", "  ")
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
}*/
