package request

import (
	"context"
	"fmt"
	"net/http"
)

import (
	"github.com/abrie/censusgeocoder/internal/service"
)

type Vintages struct {
	Benchmark string `json:"benchmark"`
}

func (params Vintages) BuildHttpRequest(ctx context.Context, service *service.Service) (*http.Request, error) {
	url := fmt.Sprintf("%s/vintages", service.Endpoint)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to build HTTP Request: %v", err)
	}

	q := req.URL.Query()
	q.Add("format", FormatJSON)
	q.Add("benchmark", params.Benchmark)

	req.URL.RawQuery = q.Encode()

	return req, nil
}
