package details

import (
	"fmt"
	"net/url"

	"github.com/johnbalvin/gozillow/trace"
)

func FromPropertyURL(propertyURL string, proxyURL *url.URL) (PropertyInfo, error) {
	data, err := fromPropertyURL(propertyURL, proxyURL)
	if err != nil {
		return PropertyInfo{}, trace.NewOrAdd(1, "main", "FromPropertyURL", err, "")
	}
	return data, nil
}

func FromPropertyID(propertyID int64, proxyURL *url.URL) (PropertyInfo, error) {
	propertyURL := fmt.Sprintf("https://www.zillow.com/homedetails/any-title/%d_zpid/", propertyID)
	data, err := fromPropertyURL(propertyURL, proxyURL)
	if err != nil {
		return PropertyInfo{}, trace.NewOrAdd(1, "main", "FromPropertyID", err, "")
	}
	return data, nil
}
