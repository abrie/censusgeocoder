package request

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

import (
	"github.com/abrie/censusgeocoder/internal/service"
)

type OneLineAddress struct {
	Address    string   `json:"address"`
	Benchmark  string   `json:"benchmark"`
	Vintage    string   `json:"vintage"`
	Layers     []string `json:"layers"`
	ReturnType string   `json:"returnType"`
}

func (params OneLineAddress) BuildHttpRequest(ctx context.Context, service *service.Service) (*http.Request, error) {
	url := strings.Join([]string{service.Endpoint, params.ReturnType, "onelineaddress"}, "/")

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to build HTTP Request: %v", err)
	}

	q := req.URL.Query()
	q.Add("address", params.Address)
	q.Add("benchmark", params.Benchmark)
	q.Add("format", FormatJSON)

	if params.ReturnType == "geographies" {
		if params.Vintage != "" {
			q.Add("vintage", params.Vintage)
		} else {
			return nil, fmt.Errorf("'vintage' parameter must be set if ReturnType is 'geographies'.")
		}

		if len(params.Layers) > 0 {
			q.Add("layers", strings.Join(params.Layers, ","))
		}
	}

	req.URL.RawQuery = q.Encode()

	return req, nil
}
