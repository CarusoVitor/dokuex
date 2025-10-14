package graphql

type PokemonSpeciesResponse struct {
	Data struct {
		PokemonV2Pokemonspecies []struct {
			Name string `json:"name"`
		} `json:"pokemon_v2_pokemonspecies"`
	} `json:"data"`
}
