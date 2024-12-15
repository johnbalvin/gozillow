package gozillow

type Filter struct {
	Beds      MinMax
	Bathrooms MinMax
	Price     MinMax
}
type MinMax struct {
	Min int `json:"min,omitempty"`
	Max int `json:"max,omitempty"`
}
type MapBounds struct {
	Ne Coordinates
	Sw Coordinates
}
type Coordinates struct {
	Latitude  float64
	Longitude float64
}

func (mapBound MapBounds) parse() MapBoundsPage {
	neLat := mapBound.Ne.Latitude
	neLong := mapBound.Ne.Longitude
	swLat := mapBound.Sw.Latitude
	swLong := mapBound.Sw.Longitude
	mb := MapBoundsPage{
		North: neLat,
		East:  neLong,
		South: swLat,
		West:  swLong,
	}
	return mb
}
func (filter Filter) parse() filterPage {
	var res filterPage
	if filter.Bathrooms.Min != 0 || filter.Bathrooms.Max != 0 {
		res.Bathrooms = &filter.Bathrooms
	}
	if filter.Price.Min != 0 || filter.Price.Max != 0 {
		res.Price = &filter.Price
	}
	if filter.Beds.Min != 0 || filter.Beds.Max != 0 {
		res.Beds = &filter.Beds
	}
	return res
}
