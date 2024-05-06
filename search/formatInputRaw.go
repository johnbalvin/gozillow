package search

type SearchRequest struct {
	IsDebugRequest   bool             `json:"isDebugRequest"`
	RequestId        int              `json:"requestId"`
	Wants            Wants            `json:"wants"`
	SearchQueryState SearchQueryState `json:"searchQueryState"`
}

type SearchQueryState struct {
	IsMapVisible  bool       `json:"isMapVisible"`
	IsListVisible bool       `json:"isListVisible"`
	MapZoom       int        `json:"mapZoom"`
	MapBounds     MapBounds  `json:"mapBounds"`
	FilterState   any        `json:"filterState"`
	Pagination    Pagination `json:"pagination"`
}

type MapBounds struct {
	East  float64 `json:"east"`
	North float64 `json:"north"`
	South float64 `json:"south"`
	West  float64 `json:"west"`
}

type FilterInputSale struct {
	SortSelection JustValueString `json:"sortSelection"`
	IsAllHomes    JustValueBool   `json:"isAllHomes"`
}

type FilterInputRent struct {
	SortSelection        JustValueString `json:"sortSelection"`
	IsNewConstruction    JustValueBool   `json:"isNewConstruction"`
	IsForSaleForeclosure JustValueBool   `json:"isForSaleForeclosure"`
	IsForSaleByOwner     JustValueBool   `json:"isForSaleByOwner"`
	IsForSaleByAgent     JustValueBool   `json:"isForSaleByAgent"`
	IsForRent            JustValueBool   `json:"isForRent"`
	IsComingSoon         JustValueBool   `json:"isComingSoon"`
	IsAuction            JustValueBool   `json:"isAuction"`
	IsAllHomes           JustValueBool   `json:"isAllHomes"`
}

type FilterInputSold struct {
	SortSelection        JustValueString `json:"sortSelection"`
	IsNewConstruction    JustValueBool   `json:"isNewConstruction"`
	IsForSaleForeclosure JustValueBool   `json:"isForSaleForeclosure"`
	IsForSaleByOwner     JustValueBool   `json:"isForSaleByOwner"`
	IsForSaleByAgent     JustValueBool   `json:"isForSaleByAgent"`
	IsComingSoon         JustValueBool   `json:"isComingSoon"`
	IsAuction            JustValueBool   `json:"isAuction"`
	IsAllHomes           JustValueBool   `json:"isAllHomes"`
	IsRecentlySold       JustValueBool   `json:"isRecentlySold"`
}
type JustValueBool struct {
	Value bool `json:"value"`
}

type JustValueString struct {
	Value string `json:"value"`
}

type Pagination struct {
	CurrentPage int `json:"currentPage"`
}

type Wants struct {
	Cat1 []string `json:"cat1"`
	Cat2 []string `json:"cat2"`
}
