package search

type InputData struct {
	Coordinates CoordinatesInput
	ZoomValue   int
}
type CoordinatesInput struct {
	Ne CoordinatesValues
	Sw CoordinatesValues
}
type CoordinatesValues struct {
	Latitude float64
	Longitud float64
}
