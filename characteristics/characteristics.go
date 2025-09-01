package characteristics

import (
	"fmt"
)

type characteristic interface {
	getPokemons(name string) (map[string]struct{}, error)
}

// TODO: implement factory pattern to return the correct characteristic based on the name
func newCharacteristic(name string) (characteristic, error) {
	return nil, nil
}

// TODO: implement set intersection
func intersect(a, b map[string]struct{}) map[string]struct{} {
	return a
}

// MatchEmAll takes a map of characteristic names to their desired values and returns a set of pokemon names
// that match all characteristics
func MatchEmAll(nameToValue map[string]string) (map[string]struct{}, error) {
	pokemons := make(map[string]struct{}, 0)
	for name, value := range nameToValue {
		characteristic, err := newCharacteristic(name)
		if err != nil {
			return nil, fmt.Errorf("error creating characteristic %s: %v", name, err)
		}
		result, err := characteristic.getPokemons(value)
		if err != nil {
			return nil, fmt.Errorf("error getting characteristic %s: %v", name, err)
		}
		if len(result) == 0 {
			return result, nil
		}
		if len(pokemons) == 0 {
			pokemons = make(map[string]struct{}, len(result))
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
