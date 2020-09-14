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

const endpoint = "https://geocoding.geo.census.gov/geocoder"

func GetBenchmarks(ctx context.Context) (*[]response.Benchmark, error) {
	service := &service.Service{Endpoint: endpoint}
	request := request.Benchmarks{}
	response, err := submitter.Submit(ctx, service, request)
	if err != nil {
		return nil, err
	}

	return &response.Benchmarks, nil
}

func GetVintages(ctx context.Context, benchmark string) (*response.Response, error) {
	service := &service.Service{Endpoint: endpoint}
	request := request.Vintages{Benchmark: benchmark}
	return submitter.Submit(ctx, service, request)
}

func SearchOneLineAddressLocations(ctx context.Context, address, benchmark string) (*response.Response, error) {
	service := &service.Service{Endpoint: endpoint}
	request := request.OneLineAddress{
		Address:    address,
		Benchmark:  benchmark,
		Format:     "json",
		ReturnType: "locations"}
	return submitter.Submit(ctx, service, request)
}

func SearchOneLineAddressGeographies(ctx context.Context, address, benchmark, vintage string) (*response.Response, error) {
	service := &service.Service{Endpoint: endpoint}
	request := request.OneLineAddress{
		Address:    address,
		Benchmark:  benchmark,
		Format:     "json",
		ReturnType: "geographies",
		Vintage:    vintage}
	return submitter.Submit(ctx, service, request)
}

func SearchAddressLocations(ctx context.Context, street, city, state, benchmark string) (*response.Response, error) {
	service := &service.Service{Endpoint: endpoint}
	request := request.Address{
		Street:     street,
		City:       city,
		State:      state,
		Benchmark:  benchmark,
		Format:     "json",
		ReturnType: "locations"}
	return submitter.Submit(ctx, service, request)
}

func SearchAddressGeographies(ctx context.Context, street, city, state, benchmark, vintage string) (*response.Response, error) {
	service := &service.Service{Endpoint: endpoint}
	request := request.Address{
		Street:     street,
		City:       city,
		State:      state,
		Benchmark:  benchmark,
		Format:     "json",
		ReturnType: "geographies",
		Vintage:    vintage}
	return submitter.Submit(ctx, service, request)
}

func SearchCoordinateLocations(ctx context.Context, x, y float64, benchmark string) (*response.Response, error) {
	service := &service.Service{Endpoint: endpoint}
	request := request.Coordinates{
		X:          x,
		Y:          y,
		Benchmark:  benchmark,
		Format:     "json",
		ReturnType: "locations"}
	return submitter.Submit(ctx, service, request)
}

func SearchCoordinateGeographies(ctx context.Context, x, y float64, benchmark, vintage string) (*response.Response, error) {
	service := &service.Service{Endpoint: endpoint}
	request := request.Coordinates{
		X:          x,
		Y:          y,
		Benchmark:  benchmark,
		Format:     "json",
		ReturnType: "geographies",
		Vintage:    vintage}
	return submitter.Submit(ctx, service, request)
}
