package details

import (
	"bytes"
	"encoding/json"
	"html"

	"github.com/johnbalvin/gozillow/trace"
	"github.com/johnbalvin/gozillow/utils"

	"github.com/PuerkitoBio/goquery"
)

func ParseBodyDetails(body []byte) (PropertyInfo, error) {
	dataRaw, err := parseBodyDetails(body)
	if err != nil {
		return PropertyInfo{}, trace.NewOrAdd(1, "main", "ParseBodyDetails", err, "")
	}
	return dataRaw, nil
}
func parseBodyDetails(body []byte) (PropertyInfo, error) {
	reader := bytes.NewReader(body)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return PropertyInfo{}, trace.NewOrAdd(1, "main", "parseBodyDetails", err, "")
	}
	htmlData, err := doc.Find("#__NEXT_DATA__").Html()
	if err != nil {
		return PropertyInfo{}, trace.NewOrAdd(2, "main", "parseBodyDetails", err, "")
	}
	htmlData = utils.RemoveSpace(html.UnescapeString(htmlData))
	var data bodyResponse
	if err := json.Unmarshal([]byte(htmlData), &data); err != nil {
		return PropertyInfo{}, trace.NewOrAdd(3, "main", "parseBodyDetails", err, "")
	}
	mapData := make(map[string]property)
	if err := json.Unmarshal([]byte(data.Props.PageProps.ComponentProps.GdpClientCache), &mapData); err != nil {
		return PropertyInfo{}, trace.NewOrAdd(4, "main", "parseBodyDetails", err, "")
	}
	for _, property := range mapData {
		return property.Property, nil
	}
	return PropertyInfo{}, trace.ErrEmpty
}