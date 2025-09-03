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
		URL  string `json:"url"`
	} `json:"pokemon_species"`
}
