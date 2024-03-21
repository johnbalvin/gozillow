package search

type output struct {
	Cat1 cat1 `json:"cat1"`
}
type cat1 struct {
	SearchResults searchResults `json:"searchResults"`
}
type searchResults struct {
	ListResults []ListResult `json:"listResults"`
}
type ListResult struct {
	Zpid                 string  `json:"zpid"`
	ID                   string  `json:"id"`
	ImgSrc               string  `json:"imgSrc"`
	Address              string  `json:"address"`
	AddressStreet        string  `json:"addressStreet"`
	AddressCity          string  `json:"addressCity"`
	AddressState         string  `json:"addressState"`
	AddressZipcode       string  `json:"addressZipcode"`
	Area                 int64   `json:"area"`
	IsUndisclosedAddress bool    `json:"isUndisclosedAddress"`
	StatusType           string  `json:"statusType"`
	DetailUrl            string  `json:"detailUrl"`
	Price                string  `json:"price"`
	CountryCurrency      string  `json:"countryCurrency"`
	UnformattedPrice     int     `json:"unformattedPrice"`
	LatLong              LatLong `json:"latLong"`
	IsZillowOwned        bool    `json:"isZillowOwned"`
	IsUserClaimingOwner  bool    `json:"isUserClaimingOwner"`
	BrokerName           string  `json:"brokerName"`
	IsUserConfirmedClaim bool    `json:"isUserConfirmedClaim"`
	Relaxed              bool    `json:"relaxed"`
	HdpData              HdpData `json:"hdpData"`
}
type LatLong struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
type HdpData struct {
	HomeInfo HomeInfo `json:"homeInfo"`
}
type HomeInfo struct {
	Country            string           `json:"country"`
	PriceReduction     string           `json:"priceReduction"`
	PriceChange        float32          `json:"priceChange"`
	TaxAssessedValue   float32          `json:"taxAssessedValue"`
	Bathrooms          float32          `json:"bathrooms"`
	Bedrooms           float32          `json:"bedrooms"`
	LivingArea         float32          `json:"livingArea"`
	HomeStatusForHDP   string           `json:"homeStatusForHDP"`
	PriceForHDP        float32          `json:"priceForHDP"`
	HomeType           string           `json:"homeType"`
	DaysOnZillow       int              `json:"daysOnZillow"`
	IsNonOwnerOccupied bool             `json:"isNonOwnerOccupied"`
	IsShowcaseListing  bool             `json:"isShowcaseListing"`
	IsPremierBuilder   bool             `json:"isPremierBuilder"`
	LotAreaValue       float32          `json:"lotAreaValue"`
	LotAreaUnit        string           `json:"lotAreaUnit"`
	Zestimate          int64            `json:"zestimate"`
	RentZestimate      int64            `json:"rentZestimate"`
	Listing_sub_type   Listing_sub_type `json:"listing_sub_type"`
	CarouselPhotos     []CarouselPhoto  `json:"carouselPhotos"`
}
type CarouselPhoto struct {
	URL string `json:"url"`
}

type Listing_sub_type struct {
	IsFSBA        *bool `json:"is_FSBA"`
	IsFSBO        *bool `json:"is_FSBO"`
	IsPending     *bool `json:"is_pending"`
	IsNewHome     *bool `json:"is_newHome"`
	IsForeclosure *bool `json:"is_foreclosure"`
	IsBankOwned   *bool `json:"is_bankOwned"`
	IsForAuction  *bool `json:"is_forAuction"`
	IsOpenHouse   *bool `json:"is_openHouse"`
	IsComingSoon  *bool `json:"is_comingSoon"`
}
