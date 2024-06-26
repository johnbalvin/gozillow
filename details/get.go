package details

import (
	"fmt"
	"net/url"
)

func FromPropertyURL(propertyURL string, proxyURL *url.URL) (PropertyInfo, error) {
	data, err := fromPropertyURL(propertyURL, proxyURL)
	if err != nil {
		return PropertyInfo{}, err
	}
	return data, nil
}

func FromPropertyID(propertyID int64, proxyURL *url.URL) (PropertyInfo, error) {
	propertyURL := fmt.Sprintf("https://www.zillow.com/homedetails/any-title/%d_zpid/", propertyID)
	data, err := fromPropertyURL(propertyURL, proxyURL)
	if err != nil {
		return PropertyInfo{}, err
	}
	return data, nil
}
