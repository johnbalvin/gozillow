package gozillow

import (
	"net/url"

	"github.com/johnbalvin/gozillow/details"
	"github.com/johnbalvin/gozillow/search"
	"github.com/johnbalvin/gozillow/trace"
)

func DetailsFromPropertyURL(propertyURL string, proxyURL *url.URL) (PropertyInfo, error) {
	propertyDetails, err := details.FromPropertyURL(propertyURL, proxyURL)
	if err != nil {
		return PropertyInfo{}, trace.NewOrAdd(1, "main", "DetailsFromPropertyURL", err, "")
	}
	return PropertyInfo(propertyDetails), nil
}
func DetailsFromPropertyID(peropertyID int64, proxyURL *url.URL) (PropertyInfo, error) {
	propertyDetails, err := details.FromPropertyID(peropertyID, proxyURL)
	if err != nil {
		return PropertyInfo{}, trace.NewOrAdd(1, "main", "GetDetailsFromPropID", err, "")
	}
	return PropertyInfo(propertyDetails), nil
}

// coordinates should 2 points one from south and one from north(if you think it like a square, this points represent the two points of the diagonal from this square)
// zoom value from 1 - 20, so from the "square" like I said on the coordinates, this represents how much zoom on this squere there is
func SearchAll(zoomValue int, coords CoordinatesInput, proxyURL *url.URL) ([]ListResult, error) {
	neLat := coords.Ne.Latitude
	neLong := coords.Ne.Longitud
	swLat := coords.Sw.Latitude
	swLong := coords.Sw.Longitud
	results, err := search.All(zoomValue, neLat, neLong, swLat, swLong, proxyURL)
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
func SearchFirstPage(zoomValue int, coords CoordinatesInput, proxyURL *url.URL) ([]ListResult, error) {
	neLat := coords.Ne.Latitude
	neLong := coords.Ne.Longitud
	swLat := coords.Sw.Latitude
	swLong := coords.Sw.Longitud
	results, err := search.FirstPage(zoomValue, neLat, neLong, swLat, swLong, proxyURL)
	if err != nil {
		return nil, trace.NewOrAdd(1, "main", "SearchFirstPage", err, "")
	}
	var convertedResults []ListResult
	for _, result := range results {
		convertedResults = append(convertedResults, ListResult(result))
	}
	return convertedResults, nil
}
