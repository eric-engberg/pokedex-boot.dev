package pokeapi

import (
	"net/http"
	"time"

	"github.com/eric-engberg/pokedex-boot.dev/internal/pokecache"
)

// Client -
type Client struct {
	httpClient http.Client
	cache      *pokecache.Cache
}

// NewClient -
func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
