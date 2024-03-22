package gozillow

import (
	"github.com/johnbalvin/gozillow/details"
	"github.com/johnbalvin/gozillow/search"
)

type PropertyInfo details.PropertyInfo
type ListResult search.ListResult

type CoordinatesInput struct {
	Ne CoordinatesValues
	Sw CoordinatesValues
}
type CoordinatesValues struct {
	Latitude float64
	Longitud float64
}
