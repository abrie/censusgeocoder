package request

import (
	"context"
	"testing"
)

import (
	"github.com/abrie/censusgeocoder/internal/service"
)

func TestCoordinatesRequest(t *testing.T) {
	params := Coordinates{
		ReturnType: "locations",
		X:          10.5,
		Y:          -30.2,
		Benchmark:  "Census_2020",
	}

	service := &service.Service{Endpoint: "https://endpoint/rest"}

	request, err := params.BuildHttpRequest(context.Background(), service)
	if err != nil {
		t.Errorf("Unexpected error from BuildHttpRequest: %v", err)
	}

	expected := "https://endpoint/rest/locations/coordinates?benchmark=Census_2020&format=json&x=10.500000&y=-30.200000"

	got := request.URL.String()
	if got != expected {
		t.Errorf("\nWant:\t'%s'\nGot:\t'%s'", expected, got)
	}
}
