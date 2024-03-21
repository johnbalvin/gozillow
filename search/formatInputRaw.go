package search

type SearchRequest struct {
	SearchQueryState SearchQueryState `json:"searchQueryState"`
	Wants            Wants            `json:"wants"`
	RequestId        int              `json:"requestId"`
	IsDebugRequest   bool             `json:"isDebugRequest"`
}

type SearchQueryState struct {
	IsMapVisible         bool        `json:"isMapVisible"`
	MapBounds            MapBounds   `json:"mapBounds"`
	FilterState          FilterState `json:"filterState"`
	IsEntirePlaceForRent bool        `json:"isEntirePlaceForRent"`
	IsRoomForRent        bool        `json:"isRoomForRent"`
	IsListVisible        bool        `json:"isListVisible"`
	MapZoom              int         `json:"mapZoom"`
	Pagination           Pagination  `json:"pagination"`
}

type MapBounds struct {
	West  float64 `json:"west"`
	East  float64 `json:"east"`
	South float64 `json:"south"`
	North float64 `json:"north"`
}

type FilterState struct {
	SortSelection SortSelection `json:"sortSelection"`
}

type SortSelection struct {
	Value string `json:"value"`
}

type Pagination struct {
	CurrentPage int `json:"currentPage"`
}

type Wants struct {
	Cat1 []string `json:"cat1"`
	Cat2 []string `json:"cat2"`
}
