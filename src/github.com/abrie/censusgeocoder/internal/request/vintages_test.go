package request

import (
	"context"
	"testing"
)

import (
	"github.com/abrie/censusgeocoder/internal/service"
)

func TestVintagesRequest(t *testing.T) {
	params := Vintages{
		Benchmark: "Census_2020",
	}

	service := &service.Service{Endpoint: "https://endpoint/rest"}

	request, err := params.BuildHttpRequest(context.Background(), service)
	if err != nil {
		t.Errorf("Unexpected error from BuildHttpRequest: %v", err)
	}

	expected := "https://endpoint/rest/vintages?benchmark=Census_2020&format=json"

	got := request.URL.String()
	if got != expected {
		t.Errorf("\nWant:\t'%s'\nGot:\t'%s'", expected, got)
	}
}
