package censusgeocoder

import (
	"fmt"
	"log"
)

import (
	"github.com/abrie/censusgeocoder/internal/request"
	"github.com/abrie/censusgeocoder/internal/service"
	"github.com/abrie/censusgeocoder/internal/submitter"
)

func SearchOneLineAddress(address string) {
	service := &service.Service{Endpoint: "https://geocoding.geo.census.gov/geocoder/locations"}
	request := request.OneLineAddress{Address: address, Benchmark: "Public_AR_Current", Format: "json"}
	response, err := submitter.Submit(service, request)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v", response)
}
