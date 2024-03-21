package details

import (
	"fmt"
	"net/url"

	"github.com/johnbalvin/gozillow/trace"
)

func GetFromPropertyURL(propertyURL string, proxyURL *url.URL) (PropertyInfo, error) {
	data, err := getFromPropertyURL(propertyURL, proxyURL)
	if err != nil {
		return PropertyInfo{}, trace.NewOrAdd(1, "main", "GetFromPropertyURL", err, "")
	}
	return data, nil
}

func GetFromPropertyID(propertyID int64, proxyURL *url.URL) (PropertyInfo, error) {
	propertyURL := fmt.Sprintf("https://www.zillow.com/homedetails/any-title/%d_zpid/", propertyID)
	data, err := getFromPropertyURL(propertyURL, proxyURL)
	if err != nil {
		return PropertyInfo{}, trace.NewOrAdd(1, "main", "GetFromPropertyURL", err, "")
	}
	return data, nil
}
