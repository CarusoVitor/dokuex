package characteristics

import (
	"encoding/json"

	"github.com/CarusoVitor/dokuex/api"
)

// endpointCharacteristic are the ones that query the api directly as
// /characteristic_name/{characteristic}
type endpointCharacteristic struct {
	name      string
	client    api.PokeClient
	formatter func([]byte) (PokemonSet, error)
}

func (ac endpointCharacteristic) getPokemons(value string) (PokemonSet, error) {
	raw, err := ac.client.FetchPokemons(ac.name, value)
	if err != nil {
		return nil, err
	}
	return ac.formatter(raw)
}

func newTypeCharacteristic(client api.PokeClient) endpointCharacteristic {
	return endpointCharacteristic{
		name:      typeName,
		client:    client,
		formatter: formatTypeResponse,
	}
}

func newGenerationCharacteristic(client api.PokeClient) endpointCharacteristic {
	return endpointCharacteristic{
		name:      generationName,
		client:    client,
		formatter: formatGenerationResponse,
	}
}

func newMoveCharacteristic(client api.PokeClient) endpointCharacteristic {
	return endpointCharacteristic{
		name:      moveName,
		client:    client,
		formatter: formatMoveResponse,
	}
}

func newAbilityCharacteristic(client api.PokeClient) endpointCharacteristic {
	return endpointCharacteristic{
		name:      abilityName,
		client:    client,
		formatter: formatAbilityResponse,
	}
}

func formatTypeResponse(values []byte) (PokemonSet, error) {
	var typeResp api.TypeResponse
	if err := json.Unmarshal(values, &typeResp); err != nil {
		return nil, err
	}
	set := make(PokemonSet, len(typeResp.Pokemon))
	for _, entry := range typeResp.Pokemon {
		set[entry.Pokemon.Name] = struct{}{}
	}
	return set, nil
}

func formatGenerationResponse(values []byte) (PokemonSet, error) {
	var generationResp api.GenerationResponse
	if err := json.Unmarshal(values, &generationResp); err != nil {
		return nil, err
	}
	set := make(PokemonSet, len(generationResp.PokemonSpecies))
	for _, pokemon := range generationResp.PokemonSpecies {
		set[pokemon.Name] = struct{}{}
	}
	return set, nil

}

func formatMoveResponse(values []byte) (PokemonSet, error) {
	var moveResp api.MoveResponse
	if err := json.Unmarshal(values, &moveResp); err != nil {
		return nil, err
	}
	set := make(PokemonSet, len(moveResp.LearnedByPokemon))
	for _, pokemon := range moveResp.LearnedByPokemon {
		set[pokemon.Name] = struct{}{}
	}
	return set, nil

}

func formatAbilityResponse(values []byte) (PokemonSet, error) {
	var abilityResp api.AbilityResponse
	if err := json.Unmarshal(values, &abilityResp); err != nil {
		return nil, err
	}
	set := make(PokemonSet, len(abilityResp.Pokemon))
	for _, pokemon := range abilityResp.Pokemon {
		set[pokemon.Pokemon.Name] = struct{}{}
	}
	return set, nil

}
