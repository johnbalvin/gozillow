package gozillow

type output struct {
	Cat1 cat1 `json:"cat1"`
}
type cat1 struct {
	SearchResults searchResults `json:"searchResults"`
}
type searchResults struct {
	ListResults []ListResult `json:"listResults"`
	MapResults  []MapResult  `json:"mapResults"`
}
type ListResult struct {
	Zpid                 string          `json:"zpid"`
	ID                   string          `json:"id"`
	HasImage             bool            `json:"hasImage"`
	AvailabilityCount    int             `json:"availabilityCount"`
	ImgSrc               string          `json:"imgSrc"`
	StatusText           string          `json:"statusText"`
	Address              string          `json:"address"`
	AddressStreet        string          `json:"addressStreet"`
	AddressCity          string          `json:"addressCity"`
	AddressState         string          `json:"addressState"`
	AddressZipcode       string          `json:"addressZipcode"`
	Area                 int64           `json:"area"`
	IsUndisclosedAddress bool            `json:"isUndisclosedAddress"`
	StatusType           string          `json:"statusType"`
	DetailUrl            string          `json:"detailUrl"`
	Price                string          `json:"price"`
	CountryCurrency      string          `json:"countryCurrency"`
	UnformattedPrice     int             `json:"unformattedPrice"`
	LatLong              LatLong         `json:"latLong"`
	BuildingName         string          `json:"buildingName"`
	IsZillowOwned        bool            `json:"isZillowOwned"`
	IsUserClaimingOwner  bool            `json:"isUserClaimingOwner"`
	BrokerName           string          `json:"brokerName"`
	IsUserConfirmedClaim bool            `json:"isUserConfirmedClaim"`
	Relaxed              bool            `json:"relaxed"`
	HdpData              HDPData         `json:"hdpData"`
	Units                []Unit          `json:"units"`
	LotId                int64           `json:"lotId"`
	CarouselPhotos       []CarouselPhoto `json:"carouselPhotos"`
}
type Unit struct {
	Price   string `json:"price"`
	Beds    string `json:"beds"`
	ForRent bool   `json:"roomForRent"`
}
type LatLong struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
type CarouselPhoto struct {
	URL string `json:"url"`
}

type HDPData struct {
	HomeInfo HomeInfo `json:"homeInfo"`
}
type HomeInfo struct {
	Zpid                    int64           `json:"zpid"`
	PriceReduction          string          `json:"priceReduction"`
	PriceChange             float32         `json:"priceChange"`
	StreetAddress           string          `json:"streetAddress"`
	Zipcode                 string          `json:"zipcode"`
	City                    string          `json:"city"`
	State                   string          `json:"state"`
	Latitude                float64         `json:"latitude"`
	Longitude               float64         `json:"longitude"`
	Price                   float32         `json:"price"`
	Bathrooms               float32         `json:"bathrooms"`
	Bedrooms                float32         `json:"bedrooms"`
	LivingArea              float32         `json:"livingArea"`
	HomeType                string          `json:"homeType"`
	HomeStatus              string          `json:"homeStatus"`
	DaysOnZillow            int             `json:"daysOnZillow"`
	IsFeatured              bool            `json:"isFeatured"`
	ShouldHighlight         bool            `json:"shouldHighlight"`
	Zestimate               int             `json:"zestimate"`
	TaxAssessedValue        float64         `json:"taxAssessedValue"`
	RentZestimate           int             `json:"rentZestimate"`
	ListingSubType          ListingSubType2 `json:"listing_sub_type"`
	IsUnmappable            bool            `json:"isUnmappable"`
	IsPreforeclosureAuction bool            `json:"isPreforeclosureAuction"`
	HomeStatusForHDP        string          `json:"homeStatusForHDP"`
	PriceForHDP             float32         `json:"priceForHDP"`
	TimeOnZillow            int64           `json:"timeOnZillow"`
	IsNonOwnerOccupied      bool            `json:"isNonOwnerOccupied"`
	IsPremierBuilder        bool            `json:"isPremierBuilder"`
	IsZillowOwned           bool            `json:"isZillowOwned"`
	Currency                string          `json:"currency"`
	Country                 string          `json:"country"`
	LotAreaValue            float32         `json:"lotAreaValue"`
	LotAreaUnit             string          `json:"lotAreaUnit"`
	IsShowcaseListing       bool            `json:"isShowcaseListing"`
}
type MapResult struct {
	Zpid                        string  `json:"zpid"`
	Plid                        string  `json:"plid"`
	BuildingName                string  `json:"buildingName"`
	BuildingId                  string  `json:"buildingId"`
	IsBuilding                  bool    `json:"isBuilding"`
	CanSaveBuilding             bool    `json:"canSaveBuilding"`
	RawHomeStatusCd             string  `json:"rawHomeStatusCd"`
	RentalMarketingSubType      string  `json:"rentalMarketingSubType"`
	MarketingStatusSimplifiedCd string  `json:"marketingStatusSimplifiedCd"`
	ImgSrc                      string  `json:"imgSrc"`
	LotId                       int64   `json:"lotId"`
	UnitCount                   int64   `json:"unitCount"`
	MinBeds                     float32 `json:"minBeds"`
	MinBaths                    float32 `json:"minBaths"`
	MinArea                     float32 `json:"minArea"`
	HasImage                    bool    `json:"hasImage"`
	DetailUrl                   string  `json:"detailUrl"`
	StatusType                  string  `json:"statusType"`
	StatusText                  string  `json:"statusText"`
	Price                       string  `json:"price"`
	PriceLabel                  string  `json:"priceLabel"`
	Address                     string  `json:"address"`
	Beds                        int     `json:"beds"`
	Baths                       float64 `json:"baths"`
	Area                        int     `json:"area"`
	LatLong                     LatLong `json:"latLong"`
	HDPData                     HDPData `json:"hdpData"`
	IsUserClaimingOwner         bool    `json:"isUserClaimingOwner"`
	IsUserConfirmedClaim        bool    `json:"isUserConfirmedClaim"`
	Pgapt                       string  `json:"pgapt"`
	Sgapt                       string  `json:"sgapt"`
	ShouldShowZestimateAsPrice  bool    `json:"shouldShowZestimateAsPrice"`
	Has3DModel                  bool    `json:"has3DModel"`
	HasVideo                    bool    `json:"hasVideo"`
	IsHomeRec                   bool    `json:"isHomeRec"`
	HasAdditionalAttributions   bool    `json:"hasAdditionalAttributions"`
	IsFeaturedListing           bool    `json:"isFeaturedListing"`
	IsShowcaseListing           bool    `json:"isShowcaseListing"`
	ListingType                 string  `json:"listingType"`
	IsFavorite                  bool    `json:"isFavorite"`
	Visited                     bool    `json:"visited"`
	Info3String                 string  `json:"info3String"`
	BrokerName                  string  `json:"brokerName"`
	TimeOnZillow                int64   `json:"timeOnZillow"`
}
