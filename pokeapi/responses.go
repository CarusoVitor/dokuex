// currently only the pokemon list fields are queried
package pokeapi

type TypeResponse struct {
	Pokemon []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon"`
}
type GenerationResponse struct {
	PokemonSpecies []struct {
		Name string `json:"name"`
	} `json:"pokemon_species"`
}

type MoveResponse struct {
	LearnedByPokemon []struct {
		Name string `json:"name"`
	} `json:"learned_by_pokemon"`
}

type AbilityResponse struct {
	Pokemon []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon"`
}
