package characteristics

import (
	"fmt"

	"github.com/CarusoVitor/dokuex/pokeapi"
)

type PokemonSet = map[string]struct{}

// TODO: implement set intersection
func intersect(a, b PokemonSet) PokemonSet {
	return a
}

// MatchEmAll takes a map of characteristic names to their desired values and returns a set of pokemon names
// that match all characteristics
func MatchEmAll(nameToValue map[string]string, client pokeapi.PokeClient) (PokemonSet, error) {
	manager := newCharacteristicManager(client)
	pokemons := make(PokemonSet, 0)
	for name, value := range nameToValue {
		char, err := manager.createCharacteristic(name)
		if err != nil {
			return nil, fmt.Errorf("error creating characteristic %s: %w", name, err)
		}
		result, err := char.getPokemons(value)
		if len(result) == 0 {
			return result, fmt.Errorf("error getting pokemons with characteristic %s: %w", name, err)
		}
		if len(pokemons) == 0 {
			pokemons = make(PokemonSet, len(result))
			for name := range result {
				pokemons[name] = struct{}{}
			}
		} else {
			pokemons = intersect(pokemons, result)
			if len(pokemons) == 0 {
				return pokemons, nil
			}
		}
	}
	return pokemons, nil
}
