package characteristics

import (
	"encoding/json"

	"github.com/CarusoVitor/dokuex/pokeapi"
)

// endpointCharacteristic are the ones that query the api directly as
// /characteristic_name/{characteristic}
type endpointCharacteristic struct {
	name      string
	client    pokeapi.PokeClient
	formatter func([]byte) (PokemonSet, error)
}

func (ac endpointCharacteristic) getPokemons(value string) (PokemonSet, error) {
	raw, err := ac.client.FetchPokemons(ac.name, value)
	if err != nil {
		return nil, err
	}
	return ac.formatter(raw)
}

func newTypeCharacteristic(client pokeapi.PokeClient) endpointCharacteristic {
	return endpointCharacteristic{
		name:      typeName,
		client:    client,
		formatter: formatTypeResponse,
	}
}

func newGenerationCharacteristic(client pokeapi.PokeClient) endpointCharacteristic {
	return endpointCharacteristic{
		name:      generationName,
		client:    client,
		formatter: formatGenerationResponse,
	}
}

func newMoveCharacteristic(client pokeapi.PokeClient) endpointCharacteristic {
	return endpointCharacteristic{
		name:      moveName,
		client:    client,
		formatter: formatMoveResponse,
	}
}

func newAbilityCharacteristic(client pokeapi.PokeClient) endpointCharacteristic {
	return endpointCharacteristic{
		name:      abilityName,
		client:    client,
		formatter: formatAbilityResponse,
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

func formatGenerationResponse(values []byte) (PokemonSet, error) {
	var generationResp pokeapi.GenerationResponse
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
	var moveResp pokeapi.MoveResponse
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
	var abilityResp pokeapi.AbilityResponse
	if err := json.Unmarshal(values, &abilityResp); err != nil {
		return nil, err
	}
	set := make(PokemonSet, len(abilityResp.Pokemon))
	for _, pokemon := range abilityResp.Pokemon {
		set[pokemon.Pokemon.Name] = struct{}{}
	}
	return set, nil

}
