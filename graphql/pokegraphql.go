package graphql

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const graphqlUrl string = "https://beta.pokeapi.co/graphql/v1beta"
const clientTimeout time.Duration = time.Second * 60

type PokeGraphQLClient interface {
	FetchPokemons(characteristic, value string) ([]byte, error)
}

type graphqlClient struct {
	client  *http.Client
	baseUrl string
}

func NewDefaultGraphQLClient() *graphqlClient {
	client := http.Client{Timeout: clientTimeout}
	return newGraphQLClient(&client, graphqlUrl)
}

func newGraphQLClient(client *http.Client, url string) *graphqlClient {
	return &graphqlClient{
		client:  client,
		baseUrl: url,
	}
}

func (g graphqlClient) queryForCharacteristic(characteristic, value string) (string, error) {
	switch characteristic {
	case "legendary":
		return queryIsLegendary, nil
	case "baby":
		return queryIsBaby, nil
	case "mythical":
		return queryIsMythical, nil

	}
	return "", fmt.Errorf("characteristic %q not supported", characteristic)
}

func (g graphqlClient) fetch(body io.Reader) ([]byte, error) {
	req, err := g.client.Post(
		g.baseUrl,
		"application/json",
		body,
	)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	respBody, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	if req.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: status code %d, body: %s", req.StatusCode, string(respBody))
	}
	return respBody, nil
}

func (g graphqlClient) FetchPokemons(characteristic, value string) ([]byte, error) {
	query, err := g.queryForCharacteristic(characteristic, value)
	if err != nil {
		return nil, err
	}
	formattedQuery := strings.NewReader(fmt.Sprintf(`{"query": %q}`, query))

	resp, err := g.fetch(formattedQuery)
	return resp, err
}
