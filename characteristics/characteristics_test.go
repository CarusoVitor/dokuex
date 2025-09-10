package characteristics

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

type endpointClientOneValue struct{}

func (ec endpointClientOneValue) typePokemons() PokemonSet {
	return PokemonSet{
		"bulbasaur": struct{}{},
		"ivysaur":   struct{}{},
		"venusaur":  struct{}{},
	}
}

func (ec endpointClientOneValue) generationPokemons() PokemonSet {
	return PokemonSet{
		"bulbasaur":  struct{}{},
		"ivysaur":    struct{}{},
		"venusaur":   struct{}{},
		"charmander": struct{}{},
		"charmeleon": struct{}{},
		"charizard":  struct{}{},
		"arcanine":   struct{}{},
	}
}

func (ec endpointClientOneValue) movePokemons() PokemonSet {
	return PokemonSet{
		"ivysaur":  struct{}{},
		"venusaur": struct{}{},
		"togekiss": struct{}{},
	}
}

func (ec endpointClientOneValue) abilityPokemons() PokemonSet {
	return PokemonSet{
		"arcanine":  struct{}{},
		"scraggy":   struct{}{},
		"mightyena": struct{}{},
	}
}

func (ec endpointClientOneValue) ultraBeastPokemons() PokemonSet {
	return PokemonSet{
		"kartana":  struct{}{},
		"nihilego": struct{}{},
	}
}

func (ec endpointClientOneValue) typeValue() string {
	return "grass"
}

func (ec endpointClientOneValue) generationValue() string {
	return "generation-I"
}
func (ec endpointClientOneValue) moveValue() string {
	return "solar-beam"
}
func (ec endpointClientOneValue) abilityValue() string {
	return "intimidate"

}

func (ec endpointClientOneValue) FetchPokemons(characteristic, value string) ([]byte, error) {
	var sb strings.Builder
	i := 0
	if characteristic == typeName && value == ec.typeValue() {
		sb.WriteString(`{"pokemon":[`)
		for p := range ec.typePokemons() {
			sb.WriteString(fmt.Sprintf(`{"pokemon":{"name":"%s"}}`, p))
			if i != len(ec.typePokemons())-1 {
				sb.WriteString(",")
			}
			i++
		}
	}
	if characteristic == generationName && value == ec.generationValue() {
		sb.WriteString(`{"pokemon_species":[`)
		for p := range ec.generationPokemons() {
			sb.WriteString(fmt.Sprintf(`{"name":"%s"}`, p))
			if i != len(ec.generationPokemons())-1 {
				sb.WriteString(",")
			}
			i++
		}
	}
	if characteristic == moveName && value == ec.moveValue() {
		sb.WriteString(`{"learned_by_pokemon":[`)
		for p := range ec.movePokemons() {
			sb.WriteString(fmt.Sprintf(`{"name":"%s"}`, p))
			if i != len(ec.movePokemons())-1 {
				sb.WriteString(",")
			}
			i++
		}
	}
	if (characteristic == abilityName && value == ec.abilityValue()) || (characteristic == ultraBeastName) {
		var pokemons PokemonSet
		if characteristic == abilityName {
			pokemons = ec.abilityPokemons()
		} else {
			pokemons = ec.ultraBeastPokemons()
		}
		sb.WriteString(`{"pokemon":[`)
		for p := range pokemons {
			sb.WriteString(fmt.Sprintf(`{"pokemon":{"name":"%s"}}`, p))
			if i != len(pokemons)-1 {
				sb.WriteString(",")
			}
			i++
		}
	}
	sb.WriteString("]}")
	return []byte(sb.String()), nil
}

func Test_endpointCharacteristic_getPokemons(t *testing.T) {
	client := endpointClientOneValue{}
	tests := []struct {
		name  string
		value string
		char  characteristic
		want  PokemonSet
	}{
		{
			name:  "get TypeCharacteristic response",
			value: client.typeValue(),
			char:  newTypeCharacteristic(client),
			want:  client.typePokemons(),
		},
		{
			name:  "get GenerationCharacteristic response",
			value: client.generationValue(),
			char:  newGenerationCharacteristic(client),
			want:  client.generationPokemons(),
		},
		{
			name:  "get MoveCharacteristic response",
			value: client.moveValue(),
			char:  newMoveCharacteristic(client),
			want:  client.movePokemons(),
		},
		{
			name:  "get AbilityCharacteristic response",
			value: client.abilityValue(),
			char:  newAbilityCharacteristic(client),
			want:  client.abilityPokemons(),
		},
		{
			name:  "get UltraBeastCharacteristic response",
			value: "",
			char:  newUltraBeastCharacteristic(client),
			want:  client.ultraBeastPokemons(),
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
