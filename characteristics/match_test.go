package characteristics

import (
	"reflect"
	"testing"

	"github.com/CarusoVitor/dokuex/api"
)

func Test_intersectSets(t *testing.T) {
	type args struct {
		smaller PokemonSet
		bigger  PokemonSet
	}
	tests := []struct {
		name string
		args args
		want PokemonSet
	}{
		{
			name: "sets fully intersect",
			args: args{
				smaller: PokemonSet{
					"x": struct{}{},
					"y": struct{}{},
				},
				bigger: PokemonSet{
					"x": struct{}{},
					"y": struct{}{},
				},
			},
			want: PokemonSet{
				"x": struct{}{},
				"y": struct{}{},
			},
		},
		{
			name: "sets do not intersect",
			args: args{
				smaller: PokemonSet{
					"z": struct{}{},
					"a": struct{}{},
				},
				bigger: PokemonSet{
					"x": struct{}{},
					"y": struct{}{},
				},
			},
			want: PokemonSet{},
		},
		{
			name: "sets partially intersect",
			args: args{
				smaller: PokemonSet{
					"x": struct{}{},
				},
				bigger: PokemonSet{
					"x": struct{}{},
					"y": struct{}{},
				},
			},
			want: PokemonSet{"x": struct{}{}},
		},
		{
			name: "smaller set is empty",
			args: args{
				smaller: PokemonSet{},
				bigger: PokemonSet{
					"x": struct{}{},
					"y": struct{}{},
				},
			},
			want: PokemonSet{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intersectSets(tt.args.smaller, tt.args.bigger); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intersectSets() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_intersectSetsBiggerSmallerSwapped(t *testing.T) {
	smaller := PokemonSet{}
	bigger := PokemonSet{
		"x": struct{}{},
		"y": struct{}{},
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	_ = intersectSets(bigger, smaller)
}

func TestMatchEmAll(t *testing.T) {
	client := newDefaultEndpointClient()
	client.typePokemons["poison"] = PokemonSet{
		"bulbasaur": struct{}{},
		"ivysaur":   struct{}{},
		"venusaur":  struct{}{},
		"koffing":   struct{}{},
	}

	type args struct {
		nameToValue map[string][]string
		client      api.PokeClient
	}
	tests := []struct {
		name    string
		args    args
		want    PokemonSet
		wantErr bool
	}{
		{
			name: "one characteristic only",
			args: args{
				nameToValue: map[string][]string{
					"type": {"grass"},
				},
				client: client,
			},
			want: client.typePokemons["grass"],
		},
		{
			name: "one characteristic multiple values",
			args: args{
				nameToValue: map[string][]string{
					"type": {"grass", "poison"},
				},
				client: client,
			},
			want: PokemonSet{
				"venusaur":  struct{}{},
				"ivysaur":   struct{}{},
				"bulbasaur": struct{}{},
			},
		},
		{
			name: "two characteristics that match",
			args: args{
				nameToValue: map[string][]string{
					"generation": {"generation-I"},
					"move":       {"solar-beam"},
				},
				client: client,
			},
			want: PokemonSet{
				"ivysaur":  struct{}{},
				"venusaur": struct{}{},
			},
		},
		{
			name: "two characteristics that don't match",
			args: args{
				nameToValue: map[string][]string{
					"ability": {"intimidate"},
					"move":    {"solar-beam"},
				},
				client: client,
			},
			want: PokemonSet{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := matchEmAll(tt.args.nameToValue, tt.args.client, nil, nil)
			if (err != nil) != tt.wantErr {
				t.Errorf("matchEmAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("matchEmAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
