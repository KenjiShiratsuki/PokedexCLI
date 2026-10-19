package pokeapi

import (
	"net/http"
	"time"

	"github.com/KenjiShiratsuki/PokedexCLI/internal/pokecache"
)

const (
	baseURL = "https://pokeapi.co/api/v2"
)

type Client struct {
	httpClient http.Client
	cache      *pokecache.Cache
}

func NewClient(timeout time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(2 * time.Minute),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
