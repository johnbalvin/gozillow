package gozillow

import (
	"net/url"

	"github.com/johnbalvin/gozillow/details"
	"github.com/johnbalvin/gozillow/search"
	"github.com/johnbalvin/gozillow/trace"
)

func GetFromPropertyURL(propertyURL string, proxyURL *url.URL) (PropertyInfo, error) {
	propertyDetails, err := details.GetFromPropertyURL(propertyURL, proxyURL)
	if err != nil {
		return PropertyInfo{}, trace.NewOrAdd(1, "main", "GetFromPropertyURL", err, "")
	}
	return PropertyInfo(propertyDetails), nil
}
func GetFromPropertyID(peropertyID int64, proxyURL *url.URL) (PropertyInfo, error) {
	propertyDetails, err := details.GetFromPropertyID(peropertyID, proxyURL)
	if err != nil {
		return PropertyInfo{}, trace.NewOrAdd(1, "main", "GetFromPropertyID", err, "")
	}
	return PropertyInfo(propertyDetails), nil
}

// coordinates should 2 points one from south and one from north(if you think it like a square, this points represent the two points of the diagonal from this square)
// zoom value from 1 - 20, so from the "square" like I said on the coordinates, this represents how much zoom on this squere there is
func SearchAll(zoomValue int, coordinates CoordinatesInput, proxyURL *url.URL) ([]ListResult, error) {
	input := search.InputData{
		ZoomValue:   zoomValue,
		Coordinates: search.CoordinatesInput(coordinates),
	}
	results, err := input.SearchAll(proxyURL)
	if err != nil {
		return nil, trace.NewOrAdd(1, "main", "SearchAll", err, "")
	}
	var convertedResults []ListResult
	for _, result := range results {
		convertedResults = append(convertedResults, ListResult(result))
	}
	return convertedResults, nil
}

// coordinates should 2 points one from south and one from north(if you think it like a square, this points represent the two points of the diagonal from this square)
// zoom value from 1 - 20, so from the "square" like I said on the coorinates, this represents how much zoom on this squere there is
func SearchFirstPage(zoomValue int, coordinates CoordinatesInput, proxyURL *url.URL) ([]ListResult, error) {
	input := search.InputData{
		ZoomValue:   zoomValue,
		Coordinates: search.CoordinatesInput(coordinates),
	}
	results, err := input.SearchFirstPage(proxyURL)
	if err != nil {
		return nil, trace.NewOrAdd(1, "main", "SearchFirstPage", err, "")
	}
	var convertedResults []ListResult
	for _, result := range results {
		convertedResults = append(convertedResults, ListResult(result))
	}
	return convertedResults, nil
}
