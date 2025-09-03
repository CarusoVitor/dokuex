package characteristics

import (
	"encoding/json"

	"github.com/CarusoVitor/dokuex/pokeapi"
)

type apiCharacteristic struct {
	name      string
	client    pokeapi.PokeClient
	formatter func([]byte) (PokemonSet, error)
}

func (ac apiCharacteristic) getPokemons(value string) (PokemonSet, error) {
	raw, err := ac.client.FetchPokemons(ac.name, value)
	if err != nil {
		return nil, err
	}
	return ac.formatter(raw)
}

func newTypeCharacteristic(client pokeapi.PokeClient) apiCharacteristic {
	return apiCharacteristic{
		name:      typeName,
		client:    client,
		formatter: formatTypeResponse,
	}
}

func formatTypeResponse(values []byte) (PokemonSet, error) {
	var typeResp pokeapi.TypeResponse
	if err := json.Unmarshal(values, &typeResp); err != nil {
		return nil, err
	}
	set := make(PokemonSet, len(typeResp.Pokemon))
	for _, entry := range typeResp.Pokemon {
		set[entry.Pokemon.Name] = struct{}{}
	}
	return set, nil
}
