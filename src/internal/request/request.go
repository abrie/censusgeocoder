package request

import (
	"fmt"
	"net/http"
)

import (
	"github.com/abrie/censusgeocoder/internal/service"
)

type Request interface {
	BuildHttpRequest(service *service.Service) (*http.Request, error)
}

type OneLineAddress struct {
	Address   string `json:"address"`
	Benchmark string `json:"benchmark"`
	Format    string `json:"format"`
}

// endpoint = https://geocoding.geo.census.gov/geocoder/locations
// benchmark = "Public_AR_Curent"
// format = "json"
func (params OneLineAddress) BuildHttpRequest(service *service.Service) (*http.Request, error) {
	url := fmt.Sprintf("%s/onelineaddress", service.Endpoint)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to build HTTP Request: %v", err)
	}

	q := req.URL.Query()
	q.Add("address", params.Address)
	q.Add("benchmark", params.Benchmark)
	q.Add("format", params.Format)

	req.URL.RawQuery = q.Encode()

	return req, nil
}

func (params Coordinates) BuildHttpRequest(service *service.Service) (*http.Request, error) {
	url := fmt.Sprintf("%s/coordinates", service.Endpoint)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to build HTTP Request: %v", err)
	}

	q := req.URL.Query()
	q.Add("x", fmt.Sprintf("%f", params.X))
	q.Add("y", fmt.Sprintf("%f", params.Y))
	q.Add("benchmark", params.Benchmark)
	q.Add("format", params.Format)

	req.URL.RawQuery = q.Encode()

	return req, nil
}

func (params Address) BuildHttpRequest(service *service.Service) (*http.Request, error) {
	url := fmt.Sprintf("%s/address", service.Endpoint)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to build HTTP Request: %v", err)
	}

	q := req.URL.Query()
	q.Add("street", params.Street)
	q.Add("city", params.City)
	q.Add("state", params.State)
	q.Add("benchmark", params.Benchmark)
	q.Add("format", params.Format)

	req.URL.RawQuery = q.Encode()

	return req, nil
}

type Address struct {
	Street    string `json:"street"`
	City      string `json:"city"`
	State     string `json:"state"`
	Benchmark string `json:"benchmark"`
	Format    string `json:"format"`
}

type Coordinates struct {
	X         float64 `json:"x"`
	Y         float64 `json:"y"`
	Benchmark string  `json:"benchmark"`
	Format    string  `json:"format"`
}

const FormatJSON = "json"

type Layers []string
