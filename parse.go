package gozillow

import (
	"bytes"
	"encoding/json"
	"errors"
	"html"

	"github.com/johnbalvin/gozillow/utils"

	"github.com/PuerkitoBio/goquery"
)

func ParseBodyDetails(body []byte) (PropertyInfo, error) {
	reader := bytes.NewReader(body)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return PropertyInfo{}, err
	}
	htmlData, err := doc.Find("#__NEXT_DATA__").Html()
	if err != nil {
		return PropertyInfo{}, err
	}
	htmlData = utils.RemoveSpace(html.UnescapeString(htmlData))
	var data bodyResponse
	if err := json.Unmarshal([]byte(htmlData), &data); err != nil {
		return PropertyInfo{}, err
	}
	mapData := make(map[string]property)
	if err := json.Unmarshal([]byte(data.Props.PageProps.ComponentProps.GdpClientCache), &mapData); err != nil {
		return PropertyInfo{}, err
	}
	for _, property := range mapData {
		return property.Property, nil
	}
	return PropertyInfo{}, errors.New("empty result")
}
