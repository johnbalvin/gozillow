package gozillow

import (
	"net/url"

	"github.com/johnbalvin/gozillow/details"
	"github.com/johnbalvin/gozillow/search"
)

func DetailsFromPropertyURL(propertyURL string, proxyURL *url.URL) (PropertyInfo, error) {
	propertyDetails, err := details.FromPropertyURL(propertyURL, proxyURL)
	if err != nil {
		return PropertyInfo{}, err
	}
	return PropertyInfo(propertyDetails), nil
}
func DetailsFromPropertyID(peropertyID int64, proxyURL *url.URL) (PropertyInfo, error) {
	propertyDetails, err := details.FromPropertyID(peropertyID, proxyURL)
	if err != nil {
		return PropertyInfo{}, err
	}
	return PropertyInfo(propertyDetails), nil
}

// coordinates should 2 points one from south and one from north(if you think it like a square, this points represent the two points of the diagonal from this square)
// zoom value from 1 - 20, so from the "square" like I said on the coordinates, this represents how much zoom on this squere there is
func SearchForSale(pagination int, zoomValue int, coords CoordinatesInput, proxyURL *url.URL) ([]ListResult, []MapResult, error) {
	return searchFunc(pagination, zoomValue, coords, proxyURL, search.ForSale)
}

// coordinates should 2 points one from south and one from north(if you think it like a square, this points represent the two points of the diagonal from this square)
// zoom value from 1 - 20, so from the "square" like I said on the coordinates, this represents how much zoom on this squere there is
func SearchForRent(pagination int, zoomValue int, coords CoordinatesInput, proxyURL *url.URL) ([]ListResult, []MapResult, error) {
	return searchFunc(pagination, zoomValue, coords, proxyURL, search.ForRent)
}

// coordinates should 2 points one from south and one from north(if you think it like a square, this points represent the two points of the diagonal from this square)
// zoom value from 1 - 20, so from the "square" like I said on the coordinates, this represents how much zoom on this squere there is
func SearchSold(pagination int, zoomValue int, coords CoordinatesInput, proxyURL *url.URL) ([]ListResult, []MapResult, error) {
	return searchFunc(pagination, zoomValue, coords, proxyURL, search.Sold)
}
func searchFunc(pagination int, zoomValue int, coords CoordinatesInput, proxyURL *url.URL, fn func(pagination int, zoomValue int, neLat, neLong, swLat, swLong float64, proxyURL *url.URL) ([]search.ListResult, []search.MapResult, error)) ([]ListResult, []MapResult, error) {
	neLat := coords.Ne.Latitude
	neLong := coords.Ne.Longitud
	swLat := coords.Sw.Latitude
	swLong := coords.Sw.Longitud
	results, mapResults, err := fn(pagination, zoomValue, neLat, neLong, swLat, swLong, proxyURL)
	if err != nil {
		return nil, nil, err
	}
	var convertedResults []ListResult
	for _, result := range results {
		convertedResults = append(convertedResults, ListResult(result))
	}
	var convertedMapResults []MapResult
	for _, result := range mapResults {
		convertedMapResults = append(convertedMapResults, MapResult(result))
	}
	return convertedResults, convertedMapResults, nil
}
