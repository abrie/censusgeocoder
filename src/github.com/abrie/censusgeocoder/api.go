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

func GetVintages(ctx context.Context, benchmark string) (*[]response.Vintage, error) {
	service := &service.Service{Endpoint: endpoint}
	request := request.Vintages{Benchmark: benchmark}
	response, err := submitter.Submit(ctx, service, request)
	if err != nil {
		return nil, err
	}

	return &response.Vintages, nil
}

func SearchOneLineAddressLocations(ctx context.Context, address, benchmark string) (*response.Result, error) {
	service := &service.Service{Endpoint: endpoint}
	request := request.OneLineAddress{
		Address:    address,
		Benchmark:  benchmark,
		ReturnType: "locations"}

	response, err := submitter.Submit(ctx, service, request)
	if err != nil {
		return nil, err
	}

	return &response.Result, nil
}

func SearchOneLineAddressGeographies(ctx context.Context, address, benchmark, vintage string, layers []string) (*response.Result, error) {
	service := &service.Service{Endpoint: endpoint}
	request := request.OneLineAddress{
		Address:    address,
		Benchmark:  benchmark,
		ReturnType: "geographies",
		Vintage:    vintage,
		Layers:     layers}

	response, err := submitter.Submit(ctx, service, request)
	if err != nil {
		return nil, err
	}

	return &response.Result, nil
}

func SearchAddressLocations(ctx context.Context, street, city, state, benchmark string) (*response.Result, error) {
	service := &service.Service{Endpoint: endpoint}
	request := request.Address{
		Street:     street,
		City:       city,
		State:      state,
		Benchmark:  benchmark,
		ReturnType: "locations"}

	response, err := submitter.Submit(ctx, service, request)
	if err != nil {
		return nil, err
	}

	return &response.Result, nil
}

func SearchAddressGeographies(ctx context.Context, street, city, state, benchmark, vintage string, layers []string) (*response.Result, error) {
	service := &service.Service{Endpoint: endpoint}
	request := request.Address{
		Street:     street,
		City:       city,
		State:      state,
		Benchmark:  benchmark,
		ReturnType: "geographies",
		Vintage:    vintage,
		Layers:     layers}

	response, err := submitter.Submit(ctx, service, request)
	if err != nil {
		return nil, err
	}

	return &response.Result, nil
}

func SearchCoordinateLocations(ctx context.Context, x, y float64, benchmark string) (*response.Result, error) {
	service := &service.Service{Endpoint: endpoint}
	request := request.Coordinates{
		X:          x,
		Y:          y,
		Benchmark:  benchmark,
		ReturnType: "locations"}

	response, err := submitter.Submit(ctx, service, request)
	if err != nil {
		return nil, err
	}

	return &response.Result, nil
}

func SearchCoordinateGeographies(ctx context.Context, x, y float64, benchmark, vintage string, layers []string) (*response.Result, error) {
	service := &service.Service{Endpoint: endpoint}
	request := request.Coordinates{
		X:          x,
		Y:          y,
		Benchmark:  benchmark,
		ReturnType: "geographies",
		Vintage:    vintage,
		Layers:     layers}

	response, err := submitter.Submit(ctx, service, request)
	if err != nil {
		return nil, err
	}

	return &response.Result, nil
}
