package characteristics

import "github.com/CarusoVitor/dokuex/pokeapi"

const ultraBeastAbility = "beast-boost"

type ultraBeastCharacteristic struct {
	client pokeapi.PokeClient
}

func (ubc ultraBeastCharacteristic) getPokemons(value string) (PokemonSet, error) {
	raw, err := ubc.client.FetchPokemons(abilityName, ultraBeastAbility)
	if err != nil {
		return nil, err
	}
	return formatAbilityResponse(raw)
}

func newUltraBeastCharacteristic(client pokeapi.PokeClient) ultraBeastCharacteristic {
	return ultraBeastCharacteristic{client: client}
}
