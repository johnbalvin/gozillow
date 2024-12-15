package gozillow

import (
	"fmt"
	"net/url"
)

func FromPropertyID(propertyID int64, proxyURL *url.URL) (PropertyInfo, error) {
	propertyURL := fmt.Sprintf("https://www.zillow.com/homedetails/any-title/%d_zpid/", propertyID)
	data, err := FromPropertyURL(propertyURL, proxyURL)
	if err != nil {
		return PropertyInfo{}, err
	}
	return data, nil
}

func (filter Filter) ForSale(paginationN, zoomValue int, searchValue string, mapBound MapBounds, proxyURL *url.URL) ([]ListResult, []MapResult, error) {
	sale := filterInputSale{
		SortSelection: justValueString{Value: "globalrelevanceex"},
		IsAllHomes:    justValueBool{Value: true},
		filterPage:    filter.parse(),
	}
	results, mapResults, err := search(paginationN, zoomValue, searchValue, mapBound, sale, proxyURL)
	if err != nil {
		return nil, nil, err
	}
	return results, mapResults, nil
}
func (filter Filter) ForRent(paginationN, zoomValue int, searchValue string, mapBound MapBounds, proxyURL *url.URL) ([]ListResult, []MapResult, error) {
	rent := filterInputRent{
		SortSelection: justValueString{Value: "priorityscore"},
		IsAllHomes:    justValueBool{Value: true},
		IsForRent:     justValueBool{Value: true},
		filterPage:    filter.parse(),
	}
	results, mapResults, err := search(paginationN, zoomValue, searchValue, mapBound, rent, proxyURL)
	if err != nil {
		return nil, nil, err
	}
	return results, mapResults, nil
}

func (filter Filter) Sold(paginationN, zoomValue int, searchValue string, mapBound MapBounds, proxyURL *url.URL) ([]ListResult, []MapResult, error) {
	sold := filterInputSold{
		SortSelection:  justValueString{Value: "globalrelevanceex"},
		IsAllHomes:     justValueBool{Value: true},
		IsRecentlySold: justValueBool{Value: true},
		filterPage:     filter.parse(),
	}
	results, mapResults, err := search(paginationN, zoomValue, searchValue, mapBound, sold, proxyURL)
	if err != nil {
		return nil, nil, err
	}
	return results, mapResults, nil
}
