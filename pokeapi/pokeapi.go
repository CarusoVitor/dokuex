package pokeapi

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"
)

const pokeapiUrl string = "https://pokeapi.co/api/v2"
const clientTimeout time.Duration = time.Second * 60

type PokeClient interface {
	FetchPokemons(characteristic, value string) ([]byte, error)
}

func NewPokeApiClient() *pokeApiClient {
	cache := newCache()
	client := http.Client{Timeout: clientTimeout}
	return &pokeApiClient{
		client:  &client,
		cache:   cache,
		baseUrl: pokeapiUrl,
	}
}

type pokeApiClient struct {
	client  *http.Client
	cache   *cache
	baseUrl string
}

type HttpError struct {
	StatusCode int
	message    []byte
}

func (h HttpError) Error() string {
	return fmt.Sprintf("[%d] - %s", h.StatusCode, h.message)
}

func newHttpError(StatusCode int, message []byte) HttpError {
	return HttpError{StatusCode: StatusCode, message: message}
}

func (c pokeApiClient) formatUrl(characteristic, value string) string {
	return fmt.Sprintf("%s/%s/%s", c.baseUrl, characteristic, value)
}

func (c pokeApiClient) FetchPokemons(characteristic, value string) ([]byte, error) {
	url := c.formatUrl(characteristic, value)
	slog.Info("called", "url", url)

	if value, ok := c.cache.get(url); ok {
		slog.Debug("url was cached")
		return value, nil
	}
	response, err := c.client.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, newHttpError(response.StatusCode, body)
	}
	c.cache.add(url, body)
	return body, nil
}
