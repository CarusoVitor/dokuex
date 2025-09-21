package characteristics

import (
	"reflect"
	"testing"
)

func Test_formatMega(t *testing.T) {
	tests := []struct {
		name     string
		pokemons []string
		want     []string
		wantErr  bool
	}{
		{
			name:     "two words mega",
			pokemons: []string{"Mega Lucario"},
			want:     []string{"lucario-mega", "lucario"},
			wantErr:  false,
		},
		{
			name:     "three words mega",
			pokemons: []string{"Mega Charizard X"},
			want:     []string{"charizard-mega-x", "charizard"},
			wantErr:  false,
		},
		{
			name:     "two and three words mega",
			pokemons: []string{"Mega Charizard X", "Mega Altaria"},
			want:     []string{"charizard-mega-x", "charizard", "altaria-mega", "altaria"},
			wantErr:  false,
		},
		{
			name:     "invalid format mega",
			pokemons: []string{"Mega Super Charizard X"},
			want:     nil,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := formatMega(tt.pokemons)
			if (err != nil) != tt.wantErr {
				t.Errorf("formatMega() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("formatMega() = %v, want %v", got, tt.want)
			}
		})
	}
}
