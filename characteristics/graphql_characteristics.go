package characteristics

import (
	"encoding/json"

	"github.com/CarusoVitor/dokuex/graphql"
)

type graphqlCharacteristic struct {
	name          string
	graphQLClient graphql.PokeGraphQLClient
	formatter     func([]byte) (PokemonSet, error)
}

func (gc graphqlCharacteristic) getPokemons(value string) (PokemonSet, error) {
	raw, err := gc.graphQLClient.FetchPokemons(gc.name, value)
	if err != nil {
		return nil, err
	}
	pokemons, err := gc.formatter(raw)
	if err != nil {
		return nil, err
	}

	return pokemons, nil
}

func newGraphQLCharacteristic(
	name string,
	client graphql.PokeGraphQLClient,
	formatter func([]byte) (PokemonSet, error),
) graphqlCharacteristic {
	return graphqlCharacteristic{
		name:          name,
		graphQLClient: client,
		formatter:     formatter,
	}
}

func formatPokemonSpeciesResponse(raw []byte) (PokemonSet, error) {
	var response graphql.PokemonSpeciesResponse
	err := json.Unmarshal(raw, &response)
	if err != nil {
		return nil, err
	}
	pokemons := response.Data.PokemonV2Pokemonspecies
	set := make(PokemonSet, len(pokemons))

	for _, poke := range pokemons {
		set[poke.Name] = struct{}{}
	}
	return set, nil
}

func newIsLegendaryCharacteristic(client graphql.PokeGraphQLClient) graphqlCharacteristic {
	return newGraphQLCharacteristic(LegendaryName, client, formatPokemonSpeciesResponse)
}

func newIsBabyCharacteristic(client graphql.PokeGraphQLClient) graphqlCharacteristic {
	return newGraphQLCharacteristic(BabyName, client, formatPokemonSpeciesResponse)
}

func newIsMythicalCharacteristic(client graphql.PokeGraphQLClient) graphqlCharacteristic {
	return newGraphQLCharacteristic(MythicalName, client, formatPokemonSpeciesResponse)
}
