package characteristics

import (
	"fmt"

	"github.com/CarusoVitor/dokuex/pokeapi"
)

const (
	typeName       string = "type"
	generationName string = "generation"
)

type invalidCharacteristicError struct {
	name string
}

func newInvalidCharacteristicsError(name string) invalidCharacteristicError {
	return invalidCharacteristicError{name: name}
}

func (e invalidCharacteristicError) Error() string {
	return fmt.Sprintf("characteristic %s was not implemented", e.name)
}

type characteristic interface {
	getPokemons(string) (PokemonSet, error)
}

type characteristicManager struct {
	client pokeapi.PokeClient
}

func newCharacteristicManager(client pokeapi.PokeClient) characteristicManager {
	return characteristicManager{client: client}
}

func (cm characteristicManager) createCharacteristic(name string) (characteristic, error) {
	switch name {
	case typeName:
		return newTypeCharacteristic(cm.client), nil
	}
	return nil, newInvalidCharacteristicsError(name)
}
