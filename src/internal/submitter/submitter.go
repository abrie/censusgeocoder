package submitter

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

import (
	"github.com/abrie/censusgeocoder/internal/request"
	"github.com/abrie/censusgeocoder/internal/response"
	"github.com/abrie/censusgeocoder/internal/service"
)

func Submit(ctx context.Context, p *service.Service, request request.Request) (*response.Response, error) {
	httpRequest, err := request.BuildHttpRequest(p)
	if err != nil {
		return nil, fmt.Errorf("Failed to submit Geocoder request to remote server: %v", err)
	}

	client := http.Client{}
	resp, err := client.Do(httpRequest)
	if err != nil {
		log.Fatal(err)
	}

	result, err := response.ParseHttpResponse(resp)
	if err != nil {
		return nil, fmt.Errorf("Failed to get result from HTTP response: %v", err)
	}

	return result, nil
}
