package censusgeocoder

import (
	"context"
)

import (
	"github.com/abrie/censusgeocoder/internal/request"
	"github.com/abrie/censusgeocoder/internal/response"
	"github.com/abrie/censusgeocoder/internal/service"
	"github.com/abrie/censusgeocoder/internal/submitter"
)

const endpoint = "https://geocoding.geo.census.gov/geocoder/locations"

func SearchOneLineAddress(ctx context.Context, address string) (*response.Response, error) {
	service := &service.Service{Endpoint: endpoint}
	request := request.OneLineAddress{Address: address, Benchmark: "Public_AR_Current", Format: "json"}
	return submitter.Submit(ctx, service, request)
}

func SearchAddress(ctx context.Context, street, city, state string) (*response.Response, error) {
	service := &service.Service{Endpoint: endpoint}
	request := request.Address{Street: street, City: city, State: state, Benchmark: "Public_AR_Current", Format: "json"}
	return submitter.Submit(ctx, service, request)
}

func SearchCoordinates(ctx context.Context, x, y float64) (*response.Response, error) {
	service := &service.Service{Endpoint: endpoint}
	request := request.Coordinates{X: x, Y: y, Benchmark: "Public_AR_Current", Format: "json"}
	return submitter.Submit(ctx, service, request)
}
