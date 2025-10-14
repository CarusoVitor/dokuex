package characteristics

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

type endpointClient struct {
	typePokemons       map[string]PokemonSet
	generationPokemons map[string]PokemonSet
	movePokemons       map[string]PokemonSet
	abilityPokemons    map[string]PokemonSet
	ultraBeastPokemons PokemonSet
}

func newDefaultEndpointClient() endpointClient {
	return endpointClient{
		typePokemons: map[string]PokemonSet{
			"grass": {
				"bulbasaur": struct{}{},
				"ivysaur":   struct{}{},
				"venusaur":  struct{}{},
			},
		},
		generationPokemons: map[string]PokemonSet{
			"generation-I": {
				"bulbasaur":  struct{}{},
				"ivysaur":    struct{}{},
				"venusaur":   struct{}{},
				"charmander": struct{}{},
				"charmeleon": struct{}{},
				"charizard":  struct{}{},
				"arcanine":   struct{}{},
			},
		},
		movePokemons: map[string]PokemonSet{
			"solar-beam": {
				"ivysaur":  struct{}{},
				"venusaur": struct{}{},
				"togekiss": struct{}{},
			},
		},
		abilityPokemons: map[string]PokemonSet{
			"intimidate": {
				"arcanine":  struct{}{},
				"scraggy":   struct{}{},
				"mightyena": struct{}{},
			},
		},
		ultraBeastPokemons: PokemonSet{
			"kartana":  struct{}{},
			"nihilego": struct{}{},
		},
	}
}

func createPokeJson(outerField, innerField string, pokes PokemonSet) []byte {
	var sb strings.Builder
	i := 0
	sb.WriteString(outerField)
	for p := range pokes {
		sb.WriteString(fmt.Sprintf(innerField, p))
		if i != len(pokes)-1 {
			sb.WriteString(",")
		}
		i++
	}
	sb.WriteString(`]}`)
	return []byte(sb.String())
}

func (ec endpointClient) FetchPokemons(characteristic, value string) ([]byte, error) {
	if characteristic == TypeName {
		return createPokeJson(`{"pokemon":[`, `{"pokemon":{"name":"%s"}}`, ec.typePokemons[value]), nil
	}
	if characteristic == GenerationName {
		return createPokeJson(`{"pokemon_species":[`, `{"name":"%s"}`, ec.generationPokemons[value]), nil
	}
	if characteristic == MoveName {
		return createPokeJson(`{"learned_by_pokemon":[`, `{"name":"%s"}`, ec.movePokemons[value]), nil
	}
	if value == ultraBeastAbility {
		return createPokeJson(`{"pokemon":[`, `{"pokemon":{"name":"%s"}}`, ec.ultraBeastPokemons), nil
	}
	if characteristic == AbilityName {
		return createPokeJson(`{"pokemon":[`, `{"pokemon":{"name":"%s"}}`, ec.abilityPokemons[value]), nil
	}
	return nil, nil
}

func Test_endpointCharacteristic_getPokemons(t *testing.T) {
	client := newDefaultEndpointClient()
	tests := []struct {
		name  string
		value string
		char  characteristic
		want  PokemonSet
	}{
		{
			name:  "get TypeCharacteristic response",
			value: "grass",
			char:  newTypeCharacteristic(client),
			want:  client.typePokemons["grass"],
		},
		{
			name:  "get GenerationCharacteristic response",
			value: "generation-I",
			char:  newGenerationCharacteristic(client),
			want:  client.generationPokemons["generation-I"],
		},
		{
			name:  "get MoveCharacteristic response",
			value: "solar-beam",
			char:  newMoveCharacteristic(client),
			want:  client.movePokemons["solar-beam"],
		},
		{
			name:  "get AbilityCharacteristic response",
			value: "intimidate",
			char:  newAbilityCharacteristic(client),
			want:  client.abilityPokemons["intimidate"],
		},
		{
			name:  "get UltraBeastCharacteristic response",
			value: "",
			char:  newUltraBeastCharacteristic(client),
			want:  client.ultraBeastPokemons,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := tt.char.getPokemons(tt.value)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("endpointCharacteristic.getPokemons() = %v, want %v", got, tt.want)
			}
		})
	}
}
