package gozillow

type bodyResponse struct {
	Props prod `json:"props"`
}
type prod struct {
	PageProps pageProps `json:"pageProps"`
}
type pageProps struct {
	Zpid           int64          `json:"zpid"`
	ComponentProps componentProps `json:"componentProps"`
}
type componentProps struct {
	GdpClientCache    string            `json:"gdpClientCache"`
	InitialReduxState InitialReduxState `json:"initialReduxState"`
}
type InitialReduxState struct {
	GDP GDP `json:"gdp"`
}
type GDP struct {
	Building Building `json:"building"`
}
type Building struct {
	FloorPlans []FloorPlan `json:"floorPlans"`
}
type FloorPlan struct {
	Zpid   string   `json:"zpid"`
	Units  []Unit2  `json:"units"`
	Videos []string `json:"videos"`
	//FloorPlanUnitPhotos []string `json:"floorPlanUnitPhotos"`
	FloorplanVRModel  *string  `json:"floorplanVRModel"`
	UnitSpecialOffers *string  `json:"unitSpecialOffers"`
	MinPrice          int      `json:"minPrice"`
	MaxPrice          int      `json:"maxPrice"`
	Beds              int      `json:"beds"`
	MaloneId          *string  `json:"maloneId"`
	AvailableFrom     string   `json:"availableFrom"`
	Baths             int      `json:"baths"`
	Name              string   `json:"name"`
	Photos            []Photo  `json:"photos"`
	Sqft              int      `json:"sqft"`
	VrModels          []string `json:"vrModels"`
	AmenityDetails    []string `json:"amenityDetails"`
	LeaseTerm         string   `json:"leaseTerm"`
	DepositsAndFees   int      `json:"depositsAndFees"`
	Description       *string  `json:"description"`
}
type Unit2 struct {
	UnitNumber                       string  `json:"unitNumber"`
	Zpid                             string  `json:"zpid"`
	HousingConnector                 bool    `json:"housingConnector"`
	HousingConnectorExclusive        bool    `json:"housingConnectorExclusive"`
	Beds                             int     `json:"beds"`
	VrModel                          *string `json:"vrModel"`
	AvailableFrom                    string  `json:"availableFrom"`
	HasApprovedThirdPartyVirtualTour bool    `json:"hasApprovedThirdPartyVirtualTour"`
	Price                            int     `json:"price"`
	MinPrice                         *int    `json:"minPrice"`
	MaxPrice                         *int    `json:"maxPrice"`
	//ThirdPartyVirtualTour            *string `json:"thirdPartyVirtualTour"`
	UnitVRModel *string `json:"unitVRModel"`
	Sqft        int     `json:"sqft"`
}
type property struct {
	Property PropertyInfo `json:"property"`
}
type PropertyInfo struct {
	ZipID                                 int64                           `json:"zpid"`
	Country                               string                          `json:"country"`
	HdpUrl                                string                          `json:"hdpUrl"`
	Price                                 int64                           `json:"price"`
	Currency                              string                          `json:"currency"`
	Latitude                              float64                         `json:"latitude"`
	Longitude                             float64                         `json:"longitude"`
	Address                               Address                         `json:"address"`
	IsListingClaimedByCurrentSignedInUser *bool                           `json:"isListingClaimedByCurrentSignedInUser"`
	IsCurrentSignedInAgentResponsible     *bool                           `json:"isCurrentSignedInAgentResponsible"`
	HomeStatus                            string                          `json:"homeStatus"`
	Bedrooms                              int                             `json:"bedrooms"`
	Bathrooms                             float32                         `json:"bathrooms"`
	YearBuilt                             int                             `json:"yearBuilt"`
	RentZestimate                         int                             `json:"rentZestimate"`
	Zestimate                             int                             `json:"zestimate"`
	LastSoldPrice                         int                             `json:"lastSoldPrice"`
	AnnualHomeownersInsurance             float32                         `json:"annualHomeownersInsurance"`
	DaysOnZillow                          int                             `json:"daysOnZillow"`
	FavoriteCount                         int                             `json:"favoriteCount"`
	MonthlyHoaFee                         float32                         `json:"monthlyHoaFee"`
	LotSize                               int64                           `json:"lotSize"`
	LotAreaValue                          float32                         `json:"lotAreaValue"`
	LotAreaUnits                          string                          `json:"lotAreaUnits"`
	PageViewCount                         int                             `json:"pageViewCount"`
	TimeOnZillow                          string                          `json:"timeOnZillow"`
	ParcelId                              string                          `json:"parcelId"`
	PropertyTaxRate                       float32                         `json:"propertyTaxRate"`
	BrokerageName                         string                          `json:"brokerageName"`
	Description                           string                          `json:"description"`
	LivingAreaUnitsShort                  string                          `json:"livingAreaUnitsShort"`
	VirtualTourUrl                        string                          `json:"virtualTourUrl"`
	DatePostedString                      string                          `json:"datePostedString"`
	PropertyTypeDimension                 string                          `json:"propertyTypeDimension"`
	IsZillowOwned                         *bool                           `json:"isZillowOwned"`
	ForeclosureJudicialType               string                          `json:"foreclosureJudicialType"`
	AttributionInfo                       AttributionInfo                 `json:"attributionInfo"`
	ResoFacts                             ResoFacts                       `json:"resoFacts"`
	MortgageRates                         MortgageRates                   `json:"mortgageRates"`
	PostingContact                        PostingContact                  `json:"postingContact"`
	ListingSubType                        ListingSubType                  `json:"listingSubType"`
	Listing_sub_type                      ListingSubType2                 `json:"listing_sub_type"`
	ForeclosureTypes                      ForeclosureTypes                `json:"foreclosureTypes"`
	HomeInsights                          []HomeInsight                   `json:"homeInsights"`
	ListedBy                              []listedBy                      `json:"listedBy"`
	PriceHistory                          []PriceHistory                  `json:"priceHistory"`
	TaxHistory                            []TaxHistory                    `json:"taxHistory"`
	Schools                               []School                        `json:"schools"`
	ResponsivePhotosOriginalRatio         []ResponsivePhotosOriginalRatio `json:"responsivePhotosOriginalRatio"`
	ResponsivePhotos                      []ResponsivePhotosOriginalRatio `json:"responsivePhotos"`
}
type Address struct {
	StreetAddress string `json:"streetAddress"`
	City          string `json:"city"`
	State         string `json:"state"`
	Zipcode       string `json:"zipcode"`
}
type ResoFacts struct {
	AccessibilityFeatures             []string        `json:"accessibilityFeatures"`
	AdditionalFeeInfo                 string          `json:"additionalFeeInfo"`
	Associations                      []Association   `json:"associations"`
	AssociationFee                    string          `json:"associationFee"`
	AssociationAmenities              []string        `json:"associationAmenities"`
	AssociationFee2                   any             `json:"associationFee2"`
	AssociationFeeIncludes            []string        `json:"associationFeeIncludes"`
	AssociationName                   string          `json:"associationName"`
	AssociationName2                  string          `json:"associationName2"`
	AssociationPhone                  string          `json:"associationPhone"`
	AssociationPhone2                 string          `json:"associationPhone2"`
	BasementYN                        *bool           `json:"basementYN"`
	BuildingName                      string          `json:"buildingName"`
	BuyerAgencyCompensation           string          `json:"buyerAgencyCompensation"`
	BuyerAgencyCompensationType       string          `json:"buyerAgencyCompensationType"`
	Appliances                        []string        `json:"appliances"`
	AtAGlanceFacts                    []AtAGlanceFact `json:"atAGlanceFacts"`
	Attic                             string          `json:"attic"`
	AvailabilityDate                  string          `json:"availabilityDate"`
	Basement                          string          `json:"basement"`
	Bathrooms                         float32         `json:"bathrooms"`
	BathroomsFull                     float32         `json:"bathroomsFull"`
	BathroomsHalf                     float32         `json:"bathroomsHalf"`
	BathroomsOneQuarter               *float32        `json:"bathroomsOneQuarter"`
	BathroomsPartial                  *float32        `json:"bathroomsPartial"`
	BathroomsFloat                    float64         `json:"bathroomsFloat"`
	BathroomsThreeQuarter             *float32        `json:"bathroomsThreeQuarter"`
	Bedrooms                          int             `json:"bedrooms"`
	BodyType                          string          `json:"bodyType"`
	CanRaiseHorses                    *bool           `json:"canRaiseHorses"`
	CarportParkingCapacity            *int            `json:"carportParkingCapacity"`
	CityRegion                        string          `json:"cityRegion"`
	CommonWalls                       string          `json:"commonWalls"`
	CommunityFeatures                 []string        `json:"communityFeatures"`
	CompensationBasedOn               string          `json:"compensationBasedOn"`
	Contingency                       string          `json:"contingency"`
	Cooling                           []string        `json:"cooling"`
	CoveredParkingCapacity            int             `json:"coveredParkingCapacity"`
	CropsIncludedYN                   *bool           `json:"cropsIncludedYN"`
	CumulativeDaysOnMarket            string          `json:"cumulativeDaysOnMarket"`
	DevelopmentStatus                 string          `json:"developmentStatus"`
	DoorFeatures                      []string        `json:"doorFeatures"`
	Electric                          []string        `json:"electric"`
	Elevation                         string          `json:"elevation"`
	ElevationUnits                    string          `json:"elevationUnits"`
	EntryLevel                        string          `json:"entryLevel"`
	EntryLocation                     string          `json:"entryLocation"`
	Exclusions                        string          `json:"exclusions"`
	FeesAndDues                       []FeesAndDues   `json:"feesAndDues"`
	Fencing                           string          `json:"fencing"`
	FireplaceFeatures                 []string        `json:"fireplaceFeatures"`
	Fireplaces                        int             `json:"fireplaces"`
	Flooring                          []string        `json:"flooring"`
	FoundationArea                    string          `json:"foundationArea"`
	Furnished                         *bool           `json:"furnished"`
	GarageParkingCapacity             int             `json:"garageParkingCapacity"`
	Gas                               string          `json:"gas"`
	GreenBuildingVerificationType     string          `json:"greenBuildingVerificationType"`
	GreenEnergyEfficient              []string        `json:"greenEnergyEfficient"`
	GreenEnergyGeneration             string          `json:"greenEnergyGeneration"`
	GreenIndoorAirQuality             string          `json:"greenIndoorAirQuality"`
	GreenSustainability               string          `json:"greenSustainability"`
	GreenWaterConservation            []string        `json:"greenWaterConservation"`
	HasAssociation                    *bool           `json:"hasAssociation"`
	HasAttachedGarage                 *bool           `json:"hasAttachedGarage"`
	HasAttachedProperty               *bool           `json:"hasAttachedProperty"`
	HasCooling                        *bool           `json:"hasCooling"`
	HasCarport                        *bool           `json:"hasCarport"`
	HasElectricOnProperty             *bool           `json:"hasElectricOnProperty"`
	HasFireplace                      *bool           `json:"hasFireplace"`
	HasGarage                         *bool           `json:"hasGarage"`
	HasHeating                        *bool           `json:"hasHeating"`
	HasLandLease                      *bool           `json:"hasLandLease"`
	HasOpenParking                    *bool           `json:"hasOpenParking"`
	HasSpa                            *bool           `json:"hasSpa"`
	HasPrivatePool                    *bool           `json:"hasPrivatePool"`
	HasView                           *bool           `json:"hasView"`
	HasWaterfrontView                 *bool           `json:"hasWaterfrontView"`
	Heating                           []string        `json:"heating"`
	HighSchool                        string          `json:"highSchool"`
	HighSchoolDistrict                string          `json:"highSchoolDistrict"`
	HoaFee                            string          `json:"hoaFee"`
	HoaFeeTotal                       string          `json:"hoaFeeTotal"`
	HomeType                          string          `json:"homeType"`
	HorseAmenities                    string          `json:"horseAmenities"`
	HorseYN                           *bool           `json:"horseYN"`
	InteriorFeatures                  []string        `json:"interiorFeatures"`
	IrrigationWaterRightsAcres        *float64        `json:"irrigationWaterRightsAcres"`
	IrrigationWaterRightsYN           *bool           `json:"irrigationWaterRightsYN"`
	IsSeniorCommunity                 *bool           `json:"isSeniorCommunity"`
	LandLeaseAmount                   *float64        `json:"landLeaseAmount"`
	LandLeaseExpirationDate           string          `json:"landLeaseExpirationDate"`
	LaundryFeatures                   []string        `json:"laundryFeatures"`
	Levels                            string          `json:"levels"`
	ListingId                         string          `json:"listingId"`
	LotFeatures                       []string        `json:"lotFeatures"`
	LotSize                           string          `json:"lotSize"`
	LivingQuarters                    []string        `json:"livingQuarters"`
	MainLevelBathrooms                *float32        `json:"mainLevelBathrooms"`
	MainLevelBedrooms                 *int            `json:"mainLevelBedrooms"`
	MarketingType                     string          `json:"marketingType"`
	MiddleOrJuniorSchool              string          `json:"middleOrJuniorSchool"`
	MiddleOrJuniorSchoolDistrict      string          `json:"middleOrJuniorSchoolDistrict"`
	Municipality                      string          `json:"municipality"`
	NumberOfUnitsInCommunity          *int            `json:"numberOfUnitsInCommunity"`
	OfferReviewDate                   string          `json:"offerReviewDate"`
	OnMarketDate                      int64           `json:"onMarketDate"`
	OpenParkingCapacity               *int            `json:"openParkingCapacity"`
	OtherEquipment                    []string        `json:"otherEquipment"`
	OtherFacts                        []string        `json:"otherFacts"`
	OtherParking                      string          `json:"otherParking"`
	OwnershipType                     string          `json:"ownershipType"`
	ParkingCapacity                   int             `json:"parkingCapacity"`
	ParkingFeatures                   []string        `json:"parkingFeatures"`
	PatioAndPorchFeatures             []string        `json:"patioAndPorchFeatures"`
	PoolFeatures                      []string        `json:"poolFeatures"`
	PricePerSquareFoot                int             `json:"pricePerSquareFoot"`
	RoadSurfaceType                   []string        `json:"roadSurfaceType"`
	RoofType                          string          `json:"roofType"`
	Rooms                             []Room          `json:"rooms"`
	SecurityFeatures                  []string        `json:"securityFeatures"`
	Sewer                             []string        `json:"sewer"`
	SpaFeatures                       []string        `json:"spaFeatures"`
	SpecialListingConditions          string          `json:"specialListingConditions"`
	Stories                           *int            `json:"stories"`
	StoriesTotal                      *int            `json:"storiesTotal"`
	SubAgencyCompensation             string          `json:"subAgencyCompensation"`
	SubAgencyCompensationType         string          `json:"subAgencyCompensationType"`
	SubdivisionName                   string          `json:"subdivisionName"`
	TotalActualRent                   *float64        `json:"totalActualRent"`
	TransactionBrokerCompensation     string          `json:"transactionBrokerCompensation"`
	TransactionBrokerCompensationType string          `json:"transactionBrokerCompensationType"`
	Utilities                         []string        `json:"utilities"`
	View                              []string        `json:"view"`
	WaterSource                       []string        `json:"waterSource"`
	WaterBodyName                     string          `json:"waterBodyName"`
	WaterfrontFeatures                []string        `json:"waterfrontFeatures"`
	WaterView                         string          `json:"waterView"`
	WaterViewYN                       *bool           `json:"waterViewYN"`
	WindowFeatures                    []string        `json:"windowFeatures"`
	YearBuilt                         int             `json:"yearBuilt"`
	Zoning                            string          `json:"zoning"`
	ZoningDescription                 string          `json:"zoningDescription"`
	AboveGradeFinishedArea            string          `json:"aboveGradeFinishedArea"`
	AdditionalParcelsDescription      string          `json:"additionalParcelsDescription"`
	ArchitecturalStyle                string          `json:"architecturalStyle"`
	BelowGradeFinishedArea            string          `json:"belowGradeFinishedArea"`
	BuilderModel                      string          `json:"builderModel"`
	BuilderName                       string          `json:"builderName"`
	BuildingArea                      string          `json:"buildingArea"`
	BuildingAreaSource                string          `json:"buildingAreaSource"`
	BuildingFeatures                  string          `json:"buildingFeatures"`
	ConstructionMaterials             []string        `json:"constructionMaterials"`
	ExteriorFeatures                  []string        `json:"exteriorFeatures"`
	FoundationDetails                 []string        `json:"foundationDetails"`
	FrontageLength                    string          `json:"frontageLength"`
	FrontageType                      string          `json:"frontageType"`
	HasAdditionalParcels              *bool           `json:"hasAdditionalParcels"`
	HasPetsAllowed                    *bool           `json:"hasPetsAllowed"`
	HasRentControl                    *bool           `json:"hasRentControl"`
	HasHomeWarranty                   *bool           `json:"hasHomeWarranty"`
	Inclusions                        []string        `json:"inclusions"`
	IncomeIncludes                    string          `json:"incomeIncludes"`
	IsNewConstruction                 *bool           `json:"isNewConstruction"`
	ListingTerms                      string          `json:"listingTerms"`
	LivingAreaRange                   string          `json:"livingAreaRange"`
	LivingAreaRangeUnits              string          `json:"livingAreaRangeUnits"`
	LivingArea                        string          `json:"livingArea"`
	LotSizeDimensions                 string          `json:"lotSizeDimensions"`
	NumberOfUnitsVacant               *int            `json:"numberOfUnitsVacant"`
	OtherStructures                   []string        `json:"otherStructures"`
	Ownership                         string          `json:"ownership"`
	ParcelNumber                      string          `json:"parcelNumber"`
	PropertyCondition                 string          `json:"propertyCondition"`
	PropertySubType                   []string        `json:"propertySubType"`
	StructureType                     string          `json:"structureType"`
	Topography                        string          `json:"topography"`
	Vegetation                        []string        `json:"vegetation"`
	WoodedArea                        string          `json:"woodedArea"`
	YearBuiltEffective                *int            `json:"yearBuiltEffective"`
	VirtualTour                       string          `json:"virtualTour"`
	ElementarySchool                  string          `json:"elementarySchool"`
	ElementarySchoolDistrict          string          `json:"elementarySchoolDistrict"`
	ListAOR                           string          `json:"listAOR"`
}

type AtAGlanceFact struct {
	FactLabel string `json:"factLabel"`
	FactValue string `json:"factValue"` // Using pointer to handle null values
}
type AttributionInfo struct {
	ListingAgreement             string          `json:"listingAgreement"`
	MlsName                      string          `json:"mlsName"`
	AgentEmail                   string          `json:"agentEmail"`
	AgentLicenseNumber           string          `json:"agentLicenseNumber"`
	AgentName                    string          `json:"agentName"`
	AgentPhoneNumber             string          `json:"agentPhoneNumber"`
	AttributionTitle             string          `json:"attributionTitle"`
	BrokerName                   string          `json:"brokerName"`
	BrokerPhoneNumber            string          `json:"brokerPhoneNumber"`
	BuyerAgentMemberStateLicense string          `json:"buyerAgentMemberStateLicense"`
	BuyerAgentName               string          `json:"buyerAgentName"`
	BuyerBrokerageName           string          `json:"buyerBrokerageName"`
	CoAgentLicenseNumber         string          `json:"coAgentLicenseNumber"`
	CoAgentName                  string          `json:"coAgentName"`
	CoAgentNumber                string          `json:"coAgentNumber"`
	LastChecked                  string          `json:"lastChecked"`
	LastUpdated                  string          `json:"lastUpdated"`
	ListingOffices               []ListingOffice `json:"listingOffices"`
	ListingAgents                []ListingAgent  `json:"listingAgents"`
	MlsDisclaimer                string          `json:"mlsDisclaimer"`
	MlsId                        string          `json:"mlsId"`
	ProviderLogo                 string          `json:"providerLogo"`
	InfoString3                  string          `json:"infoString3"`
	InfoString5                  string          `json:"infoString5"`
	InfoString10                 string          `json:"infoString10"`
	InfoString16                 string          `json:"infoString16"`
	TrueStatus                   string          `json:"trueStatus"`
}
type ListingOffice struct {
	AssociatedOfficeType string `json:"associatedOfficeType"`
	OfficeName           string `json:"officeName"`
}

type ListingAgent struct {
	AssociatedAgentType string `json:"associatedAgentType"`
	MemberFullName      string `json:"memberFullName"`
	MemberStateLicense  string `json:"memberStateLicense"`
}
type School struct {
	Distance           float32 `json:"distance"`
	Name               string  `json:"name"`
	Rating             int     `json:"rating"`
	Level              string  `json:"level"`
	StudentsPerTeacher string  `json:"studentsPerTeacher"`
	Assigned           string  `json:"assigned"`
	Grades             string  `json:"grades"`
	Link               string  `json:"link"`
	Type               string  `json:"type"`
	Size               string  `json:"size"`
	TotalCount         string  `json:"totalCount"`
	IsAssigned         string  `json:"isAssigned"`
}
type TaxHistory struct {
	Time              int64   `json:"time"`
	TaxPaid           float32 `json:"taxPaid"`
	TaxIncreaseRate   float32 `json:"taxIncreaseRate"`
	Value             float32 `json:"value"`
	ValueIncreaseRate float32 `json:"valueIncreaseRate"`
}
type PriceHistory struct {
	Date               string  `json:"date"`
	Time               int64   `json:"time"`
	Price              float32 `json:"price"`
	PricePerSquareFoot float32 `json:"pricePerSquareFoot"`
	PriceChangeRate    float32 `json:"priceChangeRate"`
	Event              string  `json:"event"`
	Source             string  `json:"source"`
	BuyerAgent         Agent   `json:"buyerAgent"`
	SellerAgent        Agent   `json:"sellerAgent"`
}
type Agent struct {
	ProfileUrl string `json:"profileUrl"`
	Name       string `json:"name"`
	Photo      Photo  `json:"photo"`
}
type Photo struct {
	URL string `json:"url"`
}
type listedBy struct {
	ID       string    `json:"id"`
	Elements []Element `json:"elements"`
}
type Element struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}
type HomeInsight struct {
	HomeInsights []HomeInsightValues `json:"insights"`
}
type HomeInsightValues struct {
	Phrases []string `json:"phrases"`
}
type ForeclosureTypes struct {
	IsBankOwned         *bool `json:"isBankOwned"`
	IsForeclosedNFS     *bool `json:"isForeclosedNFS"`
	IsPreforeclosure    *bool `json:"isPreforeclosure"`
	IsAnyForeclosure    *bool `json:"isAnyForeclosure"`
	WasNonRetailAuction *bool `json:"wasNonRetailAuction"`
	WasForeclosed       *bool `json:"wasForeclosed"`
	WasREO              *bool `json:"wasREO"`
	WasDefault          *bool `json:"wasDefault"`
}
type ListingSubType struct {
	IsFSBA        *bool `json:"isFSBA"`
	IsFSBO        *bool `json:"isFSBO"`
	IsPending     *bool `json:"isPending"`
	IsNewHome     *bool `json:"isNewHome"`
	IsForeclosure *bool `json:"isForeclosure"`
	IsBankOwned   *bool `json:"isBankOwned"`
	IsForAuction  *bool `json:"isForAuction"`
	IsOpenHouse   *bool `json:"isOpenHouse"`
	IsComingSoon  *bool `json:"isComingSoon"`
}
type ListingSubType2 struct {
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
type PostingContact struct {
	Name  string `json:"name"`
	Photo string `json:"photo"`
}
type ResponsivePhotosOriginalRatio struct {
	MixedSources MixedSources `json:"mixedSources"`
}
type MixedSources struct {
	Webp []Img `json:"webp"`
	JPEG []Img `json:"jpeg"`
}
type Img struct {
	URL   string `json:"url"`
	Width int    `json:"width"`
}
type MortgageRates struct {
	ThirtyYearFixedRate float32 `json:"thirtyYearFixedRate"`
}
type FeesAndDues struct {
	Type  string `json:"type"`
	Fee   string `json:"fee"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}
type Room struct {
	Aea                   string `json:"area"`
	Description           string `json:"description"`
	Dimensions            string `json:"dimensions"`
	Level                 string `json:"level"`
	Features              string `json:"features"`
	RoomArea              string `json:"roomArea"`
	RoomAreaSource        string `json:"roomAreaSource"`
	RoomAreaUnits         string `json:"roomAreaUnits"`
	RoomDescription       string `json:"roomDescription"`
	RoomDimensions        string `json:"roomDimensions"`
	RoomFeatures          string `json:"roomFeatures"`
	RoomLength            string `json:"roomLength"`
	RoomLengthWidthSource string `json:"roomLengthWidthSource"`
	RoomLengthWidthUnits  string `json:"roomLengthWidthUnits"`
	RoomLevel             string `json:"roomLevel"`
	RoomType              string `json:"roomType"`
	RoomWidth             string `json:"roomWidth"`
}
type Association struct {
	FeeFrequency string `json:"feeFrequency"`
	Name         string `json:"name"`
	Phone        string `json:"phone"`
}
