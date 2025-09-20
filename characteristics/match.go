package characteristics

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/CarusoVitor/dokuex/api"
	"github.com/CarusoVitor/dokuex/scraper"
)

type PokemonSet = map[string]struct{}

func intersect(a, b PokemonSet) PokemonSet {
	if len(a) < len(b) {
		return intersectSets(a, b)
	}
	return intersectSets(b, a)
}

func intersectSets(smaller, bigger PokemonSet) PokemonSet {
	if len(smaller) > len(bigger) {
		panic("intersectSets must be called with smaller being smaller than bigger")
	}
	intersection := make(PokemonSet, len(smaller))
	for name := range smaller {
		if _, ok := bigger[name]; ok {
			intersection[name] = struct{}{}
		}
	}
	return intersection
}

// MatchEmAll takes a map of characteristic names to their desired values
// and returns a set of pokemon names that match all characteristics
func MatchEmAll(
	nameToValues map[string][]string,
	pokeApiClient api.PokeClient,
	serebiiScraper scraper.SerebiiScraper,
) (PokemonSet, error) {
	manager := newCharacteristicManager(pokeApiClient, serebiiScraper)
	pokemons := make(PokemonSet, 0)
	for name, values := range nameToValues {
		char, err := manager.createCharacteristic(name)
		if err != nil {
			return nil, fmt.Errorf("error creating characteristic %s: %w", name, err)
		}
		for _, value := range values {
			result, err := char.getPokemons(value)

			var httpErr api.HttpError
			if errors.As(err, &httpErr) {
				if httpErr.StatusCode == http.StatusNotFound {
					return nil, fmt.Errorf("%s is not a valid value for %s characteristic", value, name)
				}
			}

			if err != nil {
				return nil, fmt.Errorf("error getting pokemons with characteristic %s: %w", name, err)
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
	}
	return pokemons, nil
}
