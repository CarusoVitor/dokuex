package characteristics

import (
	"fmt"

	"github.com/CarusoVitor/dokuex/api"
	"github.com/CarusoVitor/dokuex/scraper"
)

const (
	typeName       string = "type"
	generationName string = "generation"
	moveName       string = "move"
	abilityName    string = "ability"
	ultraBeastName string = "ultra-beast"
	megaName       string = "mega"
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
	pokeApiClient api.PokeClient
	bulbaScraper  scraper.BulbaScraper
}

func newCharacteristicManager(
	pokeApiClient api.PokeClient,
	bulbaScraper scraper.BulbaScraper,
) *characteristicManager {
	return &characteristicManager{pokeApiClient: pokeApiClient, bulbaScraper: bulbaScraper}
}

func (cm characteristicManager) createCharacteristic(name string) (characteristic, error) {
	switch name {
	case typeName:
		return newTypeCharacteristic(cm.pokeApiClient), nil
	case generationName:
		return newGenerationCharacteristic(cm.pokeApiClient), nil
	case moveName:
		return newMoveCharacteristic(cm.pokeApiClient), nil
	case abilityName:
		return newAbilityCharacteristic(cm.pokeApiClient), nil
	case ultraBeastName:
		return newUltraBeastCharacteristic(cm.pokeApiClient), nil
	case megaName:
		return newMegaCharacteristic(cm.bulbaScraper), nil
	}
	return nil, newInvalidCharacteristicsError(name)
}
