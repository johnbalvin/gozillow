package gozillow

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func test() {
	_, err := FromApartmentURL("https://www.zillow.com/apartments/grand-rapids-mi/glen-oaks-east-apartments/5rYc7j/", nil)
	if err != nil {
		log.Println(err)
		return
	}
	return
	//
	filter := Filter{
		Beds:      MinMax{0, 0},
		Bathrooms: MinMax{0, 0},
		Price:     MinMax{0, 0},
	}
	mapBounds := MapBounds{
		Ne: Coordinates{
			Latitude:  25.943071021982487,
			Longitude: -79.9427828779297,
		},
		Sw: Coordinates{
			Latitude:  25.46295813639662,
			Longitude: -80.74615812207033,
		},
	}
	zoomValue := 11
	pagination := 1
	var allResults []ListResult
	for {
		listResults, _, err := filter.ForRent(pagination, zoomValue, "", mapBounds, nil)
		if err != nil {
			log.Println(err)
			return
		}
		allResults = append(allResults, listResults...)
		log.Printf("Result: %d/%d\n", pagination, len(listResults))
		pagination++
		if len(listResults) == 0 {
			break
		}
	}
	rawJSON, _ := json.MarshalIndent(allResults, "", "  ")
	fmt.Printf("%s", rawJSON) //in case you don't have write permisions
	if err := os.WriteFile("./fullResult43.json", rawJSON, 06444); err != nil {
		log.Println(err)
		return
	}
}
