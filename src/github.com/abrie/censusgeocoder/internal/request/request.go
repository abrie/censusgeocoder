package request

import (
	"context"
	"net/http"
)

import (
	"github.com/abrie/censusgeocoder/internal/service"
)

const FormatJSON = "json"

type Request interface {
	BuildHttpRequest(ctx context.Context, service *service.Service) (*http.Request, error)
}
