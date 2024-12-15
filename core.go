package gozillow

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/johnbalvin/gozillow/utils"
)

func search[filtersTypes filterInputSale | filterInputRent | filterInputSold](paginationN, zoomValue int, searchValue string, mapBound MapBounds, filterState filtersTypes, proxyURL *url.URL) ([]ListResult, []MapResult, error) {
	searchReq := searchRequest{
		SearchQueryState: searchQueryState{
			IsListVisible:   true,
			IsMapVisible:    true,
			MapBounds:       mapBound.parse(),
			FilterState:     filterState,
			MapZoom:         zoomValue,
			UsersSearchTerm: searchValue,
			Pagination:      pagination{CurrentPage: paginationN},
		},
		Wants:          wants{Cat1: []string{"listResults", "mapResults"}, Cat2: []string{"total"}},
		RequestId:      10,
		IsDebugRequest: false,
	}
	rawData, err := json.Marshal(searchReq)
	if err != nil {
		return nil, nil, err
	}
	req, err := http.NewRequest("PUT", ep, bytes.NewReader(rawData))
	if err != nil {
		return nil, nil, err
	}
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Add("Accept-Language", "en")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Pragma", "no-cache")
	req.Header.Add("Sec-Ch-Ua", `"Not_A Brand";v="8", "Chromium";v="120", "Google Chrome";v="120"`)
	req.Header.Add("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Add("Sec-Ch-Ua-Platform", `"Windows"`)
	req.Header.Add("Sec-Fetch-Dest", "document")
	req.Header.Add("Sec-Fetch-Mode", "navigate")
	req.Header.Add("Sec-Fetch-Site", "none")
	req.Header.Add("Sec-Fetch-User", "?1")
	req.Header.Add("Upgrade-Insecure-Requests", "1")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	transport := &http.Transport{
		MaxIdleConnsPerHost: 30,
		DisableKeepAlives:   true,
	}
	if proxyURL != nil {
		transport.Proxy = http.ProxyURL(proxyURL)
	}
	client := &http.Client{
		Timeout: time.Minute,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Transport: transport,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}
	if resp.StatusCode != 200 {
		errData := fmt.Errorf("status: %d headers: %+v", resp.StatusCode, resp.Header)
		return nil, nil, errData
	}
	body = utils.RemoveSpaceByte(body)
	var output output
	if err := json.Unmarshal(body, &output); err != nil {
		return nil, nil, err
	}
	return output.Cat1.SearchResults.ListResults, output.Cat1.SearchResults.MapResults, nil
}

func FromPropertyURL(roomURL string, proxyURL *url.URL) (PropertyInfo, error) {
	req, err := http.NewRequest("GET", roomURL, nil)
	if err != nil {
		return PropertyInfo{}, err
	}
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Add("Accept-Language", "en")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Pragma", "no-cache")
	req.Header.Add("Sec-Ch-Ua", `"Not_A Brand";v="8", "Chromium";v="120", "Google Chrome";v="120"`)
	req.Header.Add("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Add("Sec-Ch-Ua-Platform", `"Windows"`)
	req.Header.Add("Sec-Fetch-Dest", "document")
	req.Header.Add("Sec-Fetch-Mode", "navigate")
	req.Header.Add("Sec-Fetch-Site", "none")
	req.Header.Add("Sec-Fetch-User", "?1")
	req.Header.Add("Upgrade-Insecure-Requests", "1")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	transport := &http.Transport{
		MaxIdleConnsPerHost: 30,
		DisableKeepAlives:   true,
	}
	if proxyURL != nil {
		transport.Proxy = http.ProxyURL(proxyURL)
	}
	client := &http.Client{
		Timeout: time.Minute,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Transport: transport,
	}
	resp, err := client.Do(req)
	if err != nil {
		return PropertyInfo{}, err
	}
	if resp.StatusCode == 301 {
		redirectPath := resp.Header.Get("location")
		req.URL.Path = redirectPath
		req.URL.RawQuery = ""
		resp, err = client.Do(req)
		if err != nil {
			return PropertyInfo{}, err
		}
		if resp.StatusCode != 200 {
			errData := fmt.Errorf("status: %d headers: %+v", resp.StatusCode, resp.Header)
			return PropertyInfo{}, errData
		}
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return PropertyInfo{}, err
	}
	if resp.StatusCode != 200 {
		errData := fmt.Errorf("status: %d headers: %+v", resp.StatusCode, resp.Header)
		return PropertyInfo{}, errData
	}
	data, err := ParseBodyDetailsHome(body)
	if err != nil {
		return PropertyInfo{}, err
	}
	return data, nil
}

func FromApartmentURL(roomURL string, proxyURL *url.URL) ([]FloorPlan, error) {
	req, err := http.NewRequest("GET", roomURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Add("Accept-Language", "en")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Pragma", "no-cache")
	req.Header.Add("Sec-Ch-Ua", `"Not_A Brand";v="8", "Chromium";v="120", "Google Chrome";v="120"`)
	req.Header.Add("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Add("Sec-Ch-Ua-Platform", `"Windows"`)
	req.Header.Add("Sec-Fetch-Dest", "document")
	req.Header.Add("Sec-Fetch-Mode", "navigate")
	req.Header.Add("Sec-Fetch-Site", "none")
	req.Header.Add("Sec-Fetch-User", "?1")
	req.Header.Add("Upgrade-Insecure-Requests", "1")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	transport := &http.Transport{
		MaxIdleConnsPerHost: 30,
		DisableKeepAlives:   true,
	}
	if proxyURL != nil {
		transport.Proxy = http.ProxyURL(proxyURL)
	}
	client := &http.Client{
		Timeout: time.Minute,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Transport: transport,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == 301 {
		redirectPath := resp.Header.Get("location")
		req.URL.Path = redirectPath
		req.URL.RawQuery = ""
		resp, err = client.Do(req)
		if err != nil {
			return nil, err
		}
		if resp.StatusCode != 200 {
			errData := fmt.Errorf("status: %d headers: %+v", resp.StatusCode, resp.Header)
			return nil, errData
		}
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		errData := fmt.Errorf("status: %d headers: %+v", resp.StatusCode, resp.Header)
		return nil, errData
	}
	data, err := ParseBodyDetailsApartment(body)
	if err != nil {
		return nil, err
	}
	return data, nil
}
