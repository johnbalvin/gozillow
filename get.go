package gozillow

import (
	"fmt"
	"net/url"
	"strings"
)

func FromPropertyURL(maxRetries int, roomURL string, proxyURL *url.URL) (PropertyInfo, int, error) {
	if maxRetries <= 0 {
		return PropertyInfo{}, 0, fmt.Errorf("Not correct retries")
	}
	var err error
	var propertyData PropertyInfo
	for i := 1; ; i++ {
		if i == maxRetries {
			return PropertyInfo{}, 0, err
		}
		propertyData, err = fromPropertyURL(roomURL, proxyURL)
		if err == nil {
			return propertyData, i, nil
		}
		if strings.Contains(err.Error(), "status: 403") {
			continue
		}
		return PropertyInfo{}, i, err
	}
}
func FromApartmentURL(maxRetries int, roomURL string, proxyURL *url.URL) ([]FloorPlan, int, error) {
	if maxRetries <= 0 {
		return nil, 0, fmt.Errorf("Not correct retries")
	}
	var err error
	var floorPlans []FloorPlan
	for i := 1; ; i++ {
		if i == maxRetries {
			return nil, 0, err
		}
		floorPlans, err = fromApartmentURL(roomURL, proxyURL)
		if err == nil {
			return floorPlans, i, nil
		}
		if strings.Contains(err.Error(), "status: 403") {
			continue
		}
		return nil, i, err
	}
}
