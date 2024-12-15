package gozillow

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/google/uuid"
)

type response struct {
	Data data `json:"data"`
}

type data struct {
	SearchAssistanceResult searchAssistanceResult `json:"searchAssistanceResult"`
}

type searchAssistanceResult struct {
	Results []Result `json:"results"`
}

type Result struct {
	ID        string   `json:"id,omitempty"`
	RegionId  int      `json:"regionId,omitempty"`
	SubType   string   `json:"subType,omitempty"`
	RegionIds []string `json:"regionIds,omitempty"`
}

type SearchRentalFilters struct {
	MonthlyPayment         *MinMax     `json:"monthlyPayment"`
	PetsAllowed            interface{} `json:"petsAllowed"`
	RentalAvailabilityDate interface{} `json:"rentalAvailabilityDate"`
}

// --------
type response2 struct {
	RegionState         RegionState         `json:"regionState"`
	SearchPageSeoObject SearchPageSeoObject `json:"searchPageSeoObject"`
	RequestID           int                 `json:"requestId"`
}

type RegionState struct {
	RegionInfo   []RegionInfo `json:"regionInfo"`
	RegionBounds mapBounds    `json:"regionBounds"`
}

type RegionInfo struct {
	RegionID      int    `json:"regionId"`
	RegionType    int    `json:"regionType"`
	RegionName    string `json:"regionName"`
	DisplayName   string `json:"displayName"`
	IsPointRegion bool   `json:"isPointRegion"`
}

type SearchPageSeoObject struct {
	BaseURL         string `json:"baseUrl"`
	WindowTitle     string `json:"windowTitle"`
	MetaDescription string `json:"metaDescription"`
}

func GetAutocomplete1(query string, proxyURL *url.URL) ([]Result, error) {
	id := uuid.New()
	searchReq := requestAuto{
		OperationName: "getAutocompleteResults",
		Query:         "query getAutocompleteResults($query: String!, $queryOptions: SearchAssistanceQueryOptions, $resultType: [SearchAssistanceResultType]) {\n  searchAssistanceResult: zgsAutocompleteRequest(\n    query: $query\n    queryOptions: $queryOptions\n    resultType: $resultType\n  ) {\n    requestId\n    results {\n      __typename\n      id\n      ...RegionResultFields\n      ...SemanticResultFields\n      ...RentalCommunityResultFields\n      ...SchoolResultFields\n      ...BuilderCommunityResultFields\n    }\n    __typename\n  }\n}\n\nfragment RegionResultFields on SearchAssistanceRegionResult {\n  regionId\n  subType\n  __typename\n}\n\nfragment SchoolResultFields on SearchAssistanceSchoolResult {\n  id\n  schoolDistrictId\n  schoolId\n  __typename\n}\n\nfragment SemanticResultFields on SearchAssistanceSemanticResult {\n  nearMe\n  regionIds\n  regionTypes\n  regionDisplayIds\n  queryResolutionStatus\n  viewLatitudeDelta\n  filters {\n    basementStatusType\n    baths {\n      min\n      max\n      __typename\n    }\n    beds {\n      min\n      max\n      __typename\n    }\n    excludeTypes\n    hoaFeesPerMonth {\n      min\n      max\n      __typename\n    }\n    homeType\n    keywords\n    listingStatusType\n    livingAreaSqft {\n      min\n      max\n      __typename\n    }\n    lotSizeSqft {\n      min\n      max\n      __typename\n    }\n    parkingSpots {\n      min\n      max\n      __typename\n    }\n    price {\n      min\n      max\n      __typename\n    }\n    searchRentalFilters {\n      monthlyPayment {\n        min\n        max\n        __typename\n      }\n      petsAllowed\n      rentalAvailabilityDate {\n        min\n        max\n        __typename\n      }\n      __typename\n    }\n    searchSaleFilters {\n      daysOnZillow {\n        min\n        max\n        __typename\n      }\n      __typename\n    }\n    showOnlyType\n    view\n    yearBuilt {\n      min\n      max\n      __typename\n    }\n    __typename\n  }\n  __typename\n}\n\nfragment RentalCommunityResultFields on SearchAssistanceRentalCommunityResult {\n  location {\n    latitude\n    longitude\n    __typename\n  }\n  __typename\n}\n\nfragment BuilderCommunityResultFields on SearchAssistanceBuilderCommunityResult {\n  plid\n  __typename\n}",
		Variables: variables{
			Query:      query,
			ResultType: []string{"REGIONS", "FORSALE", "RENTALS", "SOLD", "COMMUNITIES", "BUILDER_COMMUNITIES"},
			QueryOptions: queryOptions{
				UserSearchContext: "FOR_SALE",
				MaxResults:        10,
				UserIdentity: userIdentity{
					AbKey: id.String(),
				},
				UserLocation: userLocation{
					Latitude:  33.83513618052438,
					Longitude: -96.8510195,
				},
			},
		},
	}
	rawData, err := json.Marshal(searchReq)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", "https://www.zillow.com/graphql", bytes.NewReader(rawData))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "*/*")
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
	req.Header.Add("x-caller-id", "hops-homepage")
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
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var respSearch response
	if err := json.Unmarshal(body, &respSearch); err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		errData := fmt.Errorf("status: %d headers: %+v", resp.StatusCode, resp.Header)
		return nil, errData
	}
	return respSearch.Data.SearchAssistanceResult.Results, nil
}

func GetAutocomplete2(regionID int, proxyURL *url.URL) (mapBounds, error) {
	rawInput := `{"searchQueryState":{"pagination":{},"isMapVisible":true,"mapBounds":{"west":-81.28426248920138,"east":-74.85726053607638,"south":33.384359500366614,"north":36.86965680294794},"regionSelection":[{"regionId":%d}],"filterState":{"isForRent":{"value":true},"isForSaleByAgent":{"value":false},"isForSaleByOwner":{"value":false},"isNewConstruction":{"value":false},"isComingSoon":{"value":false},"isAuction":{"value":false},"isForSaleForeclosure":{"value":false},"isTownhouse":{"value":false},"isMultiFamily":{"value":false},"isCondo":{"value":false},"isLotLand":{"value":false},"isApartment":{"value":false},"isManufactured":{"value":false},"isApartmentOrCondo":{"value":false}},"isListVisible":true,"mapZoom":8},"wants":{"regionResults":["regionResults"]},"requestId":9,"isDebugRequest":false}`
	rawData := fmt.Sprintf(rawInput, regionID)
	req, err := http.NewRequest("PUT", "https://www.zillow.com/async-create-search-page-state", strings.NewReader(rawData))
	if err != nil {
		return mapBounds{}, err
	}
	req.Header.Add("Accept", "*/*")
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
	req.Header.Add("x-caller-id", "hops-homepage")
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
		return mapBounds{}, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return mapBounds{}, err
	}
	var respSearch response2
	if err := json.Unmarshal(body, &respSearch); err != nil {
		return mapBounds{}, err
	}
	if resp.StatusCode != 200 {
		errData := fmt.Errorf("status: %d headers: %+v", resp.StatusCode, resp.Header)
		return mapBounds{}, errData
	}
	return respSearch.RegionState.RegionBounds, nil
}
