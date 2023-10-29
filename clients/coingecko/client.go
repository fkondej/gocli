package coingecko

import (
	"net/http"
	"time"

	"go.uber.org/zap"
	"golang.org/x/time/rate"
)

type CoingeckoClient struct {
	httpClient  *http.Client
	ApiURL      string
	rateLimiter *rate.Limiter
	log         *zap.Logger
}

func NewCoingeckoClient(apiURL string, log *zap.Logger) *CoingeckoClient {
	return &CoingeckoClient{
		ApiURL:      apiURL,
		rateLimiter: rate.NewLimiter(rate.Every(200*time.Millisecond), 1),
		httpClient: &http.Client{
			Timeout: 2 * time.Second,
		},
		log: log,
	}
}
