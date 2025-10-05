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
	gmaxName       string = "gmax"
)

type InvalidCharacteristicError struct {
	Name string
}

func newInvalidCharacteristicsError(name string) InvalidCharacteristicError {
	return InvalidCharacteristicError{Name: name}
}

func (e InvalidCharacteristicError) Error() string {
	return fmt.Sprintf("characteristic %s was not implemented", e.Name)
}

type characteristic interface {
	getPokemons(string) (PokemonSet, error)
}

type characteristicManager struct {
	pokeApiClient  api.PokeClient
	serebiiScraper scraper.SerebiiScraper
}

func newCharacteristicManager(
	pokeApiClient api.PokeClient,
	serebiiScraper scraper.SerebiiScraper,
) *characteristicManager {
	return &characteristicManager{pokeApiClient: pokeApiClient, serebiiScraper: serebiiScraper}
}

func (cm *characteristicManager) createCharacteristic(name string) (characteristic, error) {
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
		return newMegaCharacteristic(cm.serebiiScraper), nil
	case gmaxName:
		return newGmaxCharacteristic(cm.serebiiScraper), nil
	}
	return nil, newInvalidCharacteristicsError(name)
}
