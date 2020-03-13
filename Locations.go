package main

import "strconv"

type Location struct {
	Location   Coordinates
	Country    string
	Population int64
	Region     string
	Language   string
}
type Normal interface {
	getLocation() Coordinates
	getFormattedPopulation() string
	getCountry() string
}

func (l *Location) getCountry() string {
	if l == nil {

	}
	return l.Country
}
func (l Location) getLocation() Coordinates {
	return l.Location
}
func (l Location) getFormattedPopulation() string {
	return strconv.FormatFloat(float64(l.Population), 'E', 4, 64)
}
