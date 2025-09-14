package characteristics

import "github.com/CarusoVitor/dokuex/api"

const ultraBeastAbility = "beast-boost"

type ultraBeastCharacteristic struct {
	client api.PokeClient
}

func (ubc ultraBeastCharacteristic) getPokemons(value string) (PokemonSet, error) {
	raw, err := ubc.client.FetchPokemons(abilityName, ultraBeastAbility)
	if err != nil {
		return nil, err
	}
	return formatAbilityResponse(raw)
}

func newUltraBeastCharacteristic(client api.PokeClient) ultraBeastCharacteristic {
	return ultraBeastCharacteristic{client: client}
}
