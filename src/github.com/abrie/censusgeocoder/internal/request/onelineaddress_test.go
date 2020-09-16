package request

import (
	"context"
	"testing"
)

import (
	"github.com/abrie/censusgeocoder/internal/service"
)

func TestOneLineAddressRequest(t *testing.T) {
	params := OneLineAddress{
		ReturnType: "locations",
		Benchmark:  "Census_2020",
		Address:    "123 Fake St., Fakeville, Fake Island",
	}

	service := &service.Service{Endpoint: "https://endpoint/rest"}

	request, err := params.BuildHttpRequest(context.Background(), service)
	if err != nil {
		t.Errorf("Unexpected error from BuildHttpRequest: %v", err)
	}

	expected := "https://endpoint/rest/locations/onelineaddress?address=123+Fake+St.%2C+Fakeville%2C+Fake+Island&benchmark=Census_2020&format=json"

	got := request.URL.String()
	if got != expected {
		t.Errorf("\nWant:\t'%s'\nGot:\t'%s'", expected, got)
	}
}
