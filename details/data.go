package details

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/johnbalvin/gozillow/trace"
)

func fromPropertyURL(roomURL string, proxyURL *url.URL) (PropertyInfo, error) {
	req, err := http.NewRequest("GET", roomURL, nil)
	if err != nil {
		return PropertyInfo{}, trace.NewOrAdd(1, "details", "fromPropertyURL", err, "")
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
		return PropertyInfo{}, trace.NewOrAdd(2, "details", "fromPropertyURL", err, "")
	}
	if resp.StatusCode == 301 {
		redirectPath := resp.Header.Get("location")
		req.URL.Path = redirectPath
		req.URL.RawQuery = ""
		resp, err = client.Do(req)
		if err != nil {
			return PropertyInfo{}, trace.NewOrAdd(3, "details", "fromPropertyURL", err, "")
		}
		if resp.StatusCode != 200 {
			errData := fmt.Sprintf("status: %d headers: %+v", resp.StatusCode, resp.Header)
			return PropertyInfo{}, trace.NewOrAdd(4, "details", "fromPropertyURL", trace.ErrStatusCode, errData)
		}
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return PropertyInfo{}, trace.NewOrAdd(5, "details", "fromPropertyURL", err, "")
	}
	if resp.StatusCode != 200 {
		errData := fmt.Sprintf("status: %d headers: %+v", resp.StatusCode, resp.Header)
		return PropertyInfo{}, trace.NewOrAdd(6, "details", "fromPropertyURL", trace.ErrStatusCode, errData)
	}
	data, err := ParseBodyDetails(body)
	if err != nil {
		return PropertyInfo{}, trace.NewOrAdd(7, "details", "fromPropertyURL", err, "")
	}
	return data, nil
}
