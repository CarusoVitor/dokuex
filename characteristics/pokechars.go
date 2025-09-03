package characteristics

import (
	"encoding/json"

	"github.com/CarusoVitor/dokuex/pokeapi"
)

type typeCharacteristic struct {
	client pokeapi.PokeClient
}

func newTypeCharacteristic(client pokeapi.PokeClient) typeCharacteristic {
	return typeCharacteristic{client: client}
}

func (tc typeCharacteristic) getPokemons(value string) (PokemonSet, error) {
	rawPokemons, err := tc.client.FetchPokemons(typeName, value)
	if err != nil {
		return nil, err
	}
	pokemons, err := tc.formatResponse(rawPokemons)
	if err != nil {
		return nil, err
	}

	return pokemons, nil
}

func (tc typeCharacteristic) formatResponse(values []byte) (PokemonSet, error) {
	var typeResp pokeapi.TypeResponse
	err := json.Unmarshal(values, &typeResp)
	if err != nil {
		panic("typeCharacteristic.formatResponse: unmarshaling must not produce an error here")
	}

	set := make(PokemonSet, len(typeResp.Pokemon))
	for _, entry := range typeResp.Pokemon {
		set[entry.Pokemon.Name] = struct{}{}
	}
	return set, nil
}
