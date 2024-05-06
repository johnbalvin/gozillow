package gozillow

import (
	"github.com/johnbalvin/gozillow/details"
	"github.com/johnbalvin/gozillow/search"
)

type PropertyInfo details.PropertyInfo
type ListResult search.ListResult
type MapResult search.MapResult

type CoordinatesInput struct {
	Ne CoordinatesValues
	Sw CoordinatesValues
}
type CoordinatesValues struct {
	Latitude float64
	Longitud float64
}
type FilterInput struct {
	SortSelection        string
	IsNewConstruction    bool
	IsForSaleForeclosure bool
	IsForSaleByOwner     bool
	IsForSaleByAgent     bool
	IsForRent            bool
	IsComingSoon         bool
	IsAuction            bool
	IsAllHomes           bool
}
