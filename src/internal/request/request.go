package request

import (
	"context"
	"fmt"
	"net/http"
)

import (
	"github.com/abrie/censusgeocoder/internal/service"
)

type Request interface {
	BuildHttpRequest(ctx context.Context, service *service.Service) (*http.Request, error)
}

type OneLineAddress struct {
	Address    string `json:"address"`
	Benchmark  string `json:"benchmark"`
	Format     string `json:"format"`
	Vintage    string `json:"vintage"`
	ReturnType string `json:"returnType"`
}

func (params OneLineAddress) BuildHttpRequest(ctx context.Context, service *service.Service) (*http.Request, error) {
	url := fmt.Sprintf("%s/%s/onelineaddress", service.Endpoint, params.ReturnType)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to build HTTP Request: %v", err)
	}

	q := req.URL.Query()

	q.Add("address", params.Address)
	q.Add("benchmark", params.Benchmark)
	q.Add("format", params.Format)

	if params.ReturnType == "geographies" {
		if params.Vintage != "" {
			q.Add("vintage", params.Vintage)
		} else {
			return nil, fmt.Errorf("'vintage' parameter must be set if ReturnType is 'geographies'.")
		}
	}

	req.URL.RawQuery = q.Encode()

	return req, nil
}

func (params Coordinates) BuildHttpRequest(ctx context.Context, service *service.Service) (*http.Request, error) {
	url := fmt.Sprintf("%s/%s/coordinates", service.Endpoint, params.ReturnType)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to build HTTP Request: %v", err)
	}

	q := req.URL.Query()

	q.Add("x", fmt.Sprintf("%f", params.X))
	q.Add("y", fmt.Sprintf("%f", params.Y))
	q.Add("benchmark", params.Benchmark)
	q.Add("format", params.Format)

	if params.ReturnType == "geographies" {
		if params.Vintage != "" {
			q.Add("vintage", params.Vintage)
		} else {
			return nil, fmt.Errorf("'vintage' parameter must be set if ReturnType is 'geographies'.")
		}
	}

	req.URL.RawQuery = q.Encode()

	return req, nil
}

func (params Address) BuildHttpRequest(ctx context.Context, service *service.Service) (*http.Request, error) {
	url := fmt.Sprintf("%s/%s/address", service.Endpoint, params.ReturnType)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to build HTTP Request: %v", err)
	}

	q := req.URL.Query()

	q.Add("street", params.Street)
	q.Add("city", params.City)
	q.Add("state", params.State)
	q.Add("benchmark", params.Benchmark)
	q.Add("format", params.Format)

	if params.ReturnType == "geographies" {
		if params.Vintage != "" {
			q.Add("vintage", params.Vintage)
		} else {
			return nil, fmt.Errorf("'vintage' parameter must be set if ReturnType is 'geographies'.")
		}
	}

	req.URL.RawQuery = q.Encode()

	return req, nil
}

type Address struct {
	Street     string `json:"street"`
	City       string `json:"city"`
	State      string `json:"state"`
	Benchmark  string `json:"benchmark"`
	Format     string `json:"format"`
	Vintage    string `json:"vintage"`
	ReturnType string `json:"returnType"`
}

type Coordinates struct {
	X          float64 `json:"x"`
	Y          float64 `json:"y"`
	Benchmark  string  `json:"benchmark"`
	Format     string  `json:"format"`
	Vintage    string  `json:"vintage"`
	ReturnType string  `json:"returnType"`
}

const FormatJSON = "json"

type Layers []string
