package censusgeocoder

import (
	"github.com/abrie/censusgeocoder/internal/request"
	"github.com/abrie/censusgeocoder/internal/response"
	"github.com/abrie/censusgeocoder/internal/service"
	"github.com/abrie/censusgeocoder/internal/submitter"
)

func SearchOneLineAddress(address string) (*response.Response, error) {
	service := &service.Service{Endpoint: "https://geocoding.geo.census.gov/geocoder/locations"}
	request := request.OneLineAddress{Address: address, Benchmark: "Public_AR_Current", Format: "json"}
	return submitter.Submit(service, request)
}

func SearchAddress(street, city, state string) (*response.Response, error) {
	service := &service.Service{Endpoint: "https://geocoding.geo.census.gov/geocoder/locations"}
	request := request.Address{Street: street, City: city, State: state, Benchmark: "Public_AR_Current", Format: "json"}
	return submitter.Submit(service, request)
}

func SearchCoordinates(x, y float64) (*response.Response, error) {
	service := &service.Service{Endpoint: "https://geocoding.geo.census.gov/geocoder/locations"}
	request := request.Coordinates{X: x, Y: y, Benchmark: "Public_AR_Current", Format: "json"}
	return submitter.Submit(service, request)
}
