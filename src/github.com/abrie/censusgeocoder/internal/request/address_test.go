package request

import (
	"context"
	"testing"
)

import (
	"github.com/abrie/censusgeocoder/internal/service"
)

func TestAddressRequest(t *testing.T) {
	params := Address{
		ReturnType: "locations",
		Street:     "123 Fake St.",
		City:       "Fakeville",
		State:      "Fake Island",
		Benchmark:  "Census_2020",
	}

	service := &service.Service{Endpoint: "https://endpoint/rest"}

	request, err := params.BuildHttpRequest(context.Background(), service)
	if err != nil {
		t.Errorf("Unexpected error from BuildHttpRequest: %v", err)
	}

	expected := "https://endpoint/rest/locations/address?benchmark=Census_2020&city=Fakeville&format=json&state=Fake+Island&street=123+Fake+St."

	got := request.URL.String()
	if got != expected {
		t.Errorf("\nWant:\t'%s'\nGot:\t'%s'", expected, got)
	}
}
