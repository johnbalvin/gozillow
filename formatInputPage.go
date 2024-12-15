package gozillow

type filterPage struct {
	Beds      *MinMax `json:"beds,omitempty"`
	Bathrooms *MinMax `json:"baths,omitempty"`
	Price     *MinMax `json:"price,omitempty"`
}
type filterInputSale struct {
	SortSelection justValueString `json:"sortSelection"`
	IsAllHomes    justValueBool   `json:"isAllHomes"`
	filterPage
}
type filterInputRent struct {
	SortSelection        justValueString `json:"sortSelection"`
	IsNewConstruction    justValueBool   `json:"isNewConstruction"`
	IsForSaleForeclosure justValueBool   `json:"isForSaleForeclosure"`
	IsForSaleByOwner     justValueBool   `json:"isForSaleByOwner"`
	IsForSaleByAgent     justValueBool   `json:"isForSaleByAgent"`
	IsForRent            justValueBool   `json:"isForRent"`
	IsComingSoon         justValueBool   `json:"isComingSoon"`
	IsAuction            justValueBool   `json:"isAuction"`
	IsAllHomes           justValueBool   `json:"isAllHomes"`
	filterPage
}

type filterInputSold struct {
	SortSelection        justValueString `json:"sortSelection"`
	IsNewConstruction    justValueBool   `json:"isNewConstruction"`
	IsForSaleForeclosure justValueBool   `json:"isForSaleForeclosure"`
	IsForSaleByOwner     justValueBool   `json:"isForSaleByOwner"`
	IsForSaleByAgent     justValueBool   `json:"isForSaleByAgent"`
	IsComingSoon         justValueBool   `json:"isComingSoon"`
	IsAuction            justValueBool   `json:"isAuction"`
	IsAllHomes           justValueBool   `json:"isAllHomes"`
	IsRecentlySold       justValueBool   `json:"isRecentlySold"`
	filterPage
}
type searchRequest struct {
	IsDebugRequest   bool             `json:"isDebugRequest"`
	RequestId        int              `json:"requestId"`
	Wants            wants            `json:"wants"`
	SearchQueryState searchQueryState `json:"searchQueryState"`
}

type searchQueryState struct {
	IsMapVisible    bool          `json:"isMapVisible"`
	IsListVisible   bool          `json:"isListVisible"`
	MapZoom         int           `json:"mapZoom"`
	MapBounds       MapBoundsPage `json:"mapBounds"`
	FilterState     any           `json:"filterState"`
	Pagination      pagination    `json:"pagination"`
	UsersSearchTerm string        `json:"usersSearchTerm"`
}

type MapBoundsPage struct {
	East  float64 `json:"east"`
	North float64 `json:"north"`
	South float64 `json:"south"`
	West  float64 `json:"west"`
}

type justValueBool struct {
	Value bool `json:"value"`
}

type justValueString struct {
	Value string `json:"value"`
}

type pagination struct {
	CurrentPage int `json:"currentPage"`
}

type wants struct {
	Cat1 []string `json:"cat1"`
	Cat2 []string `json:"cat2"`
}
