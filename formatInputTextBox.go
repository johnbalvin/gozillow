package gozillow

type requestAuto struct {
	OperationName string    `json:"operationName"`
	Variables     variables `json:"variables"`
	Query         string    `json:"query"`
}

type variables struct {
	Query        string       `json:"query"`
	QueryOptions queryOptions `json:"queryOptions"`
	ResultType   []string     `json:"resultType"`
}

type queryOptions struct {
	MaxResults        int          `json:"maxResults"`
	UserIdentity      userIdentity `json:"userIdentity"`
	UserSearchContext string       `json:"userSearchContext"`
	UserLocation      userLocation `json:"userLocation"`
}

type userIdentity struct {
	AbKey string `json:"abKey"`
}

type userLocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
