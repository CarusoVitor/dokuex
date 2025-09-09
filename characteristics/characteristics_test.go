package characteristics

import (
	"reflect"
	"testing"
)

type endpointClientOneValue struct{}

const (
	testTypeValue       string = "grass"
	testGenerationValue string = "generation-I"
	testMoveValue       string = "hydro-pump"
	testAbilityValue    string = "beast-boost"
)

func (d endpointClientOneValue) FetchPokemons(characteristic, value string) ([]byte, error) {
	if characteristic == typeName && value == testTypeValue {
		return []byte(`{"pokemon":[{"pokemon":{"name":"bulbasaur"}},{"pokemon":{"name":"venusaur"}}]}`), nil
	}
	if characteristic == generationName && value == testGenerationValue {
		return []byte(`{"pokemon_species":[{"name":"growlithe"},{"name":"aerodactyl"}]}`), nil
	}
	if characteristic == moveName && value == testMoveValue {
		return []byte(`{"learned_by_pokemon":[{"name":"blastoise"},{"name":"feraligatr"}]}`), nil
	}
	if characteristic == abilityName && value == testAbilityValue {
		return []byte(`{"pokemon":[{"pokemon":{"name":"stakataka"}},{"pokemon":{"name":"kartana"}}]}`), nil
	}
	return nil, nil
}

func Test_endpointCharacteristic_getPokemons(t *testing.T) {
	tests := []struct {
		name  string
		value string
		char  endpointCharacteristic
		want  PokemonSet
	}{
		{
			name:  "get TypeCharacteristic response",
			value: testTypeValue,
			char:  newTypeCharacteristic(endpointClientOneValue{}),
			want: PokemonSet{
				"bulbasaur": struct{}{},
				"venusaur":  struct{}{},
			},
		},
		{
			name:  "get GenerationCharacteristic response",
			value: testGenerationValue,
			char:  newGenerationCharacteristic(endpointClientOneValue{}),
			want: PokemonSet{
				"growlithe":  struct{}{},
				"aerodactyl": struct{}{},
			},
		},
		{
			name:  "get MoveCharacteristic response",
			value: testMoveValue,
			char:  newMoveCharacteristic(endpointClientOneValue{}),
			want: PokemonSet{
				"blastoise":  struct{}{},
				"feraligatr": struct{}{},
			},
		},
		{
			name:  "get AbilityCharacteristic response",
			value: testAbilityValue,
			char:  newAbilityCharacteristic(endpointClientOneValue{}),
			want: PokemonSet{
				"stakataka": struct{}{},
				"kartana":   struct{}{},
			},
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
