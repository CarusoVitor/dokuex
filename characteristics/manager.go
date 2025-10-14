package characteristics

import (
	"fmt"

	"github.com/CarusoVitor/dokuex/api"
	"github.com/CarusoVitor/dokuex/graphql"
	"github.com/CarusoVitor/dokuex/scraper"
)

const (
	typeName        string = "type"
	generationName  string = "generation"
	moveName        string = "move"
	abilityName     string = "ability"
	ultraBeastName  string = "ultra-beast"
	megaName        string = "mega"
	gmaxName        string = "gmax"
	isLegendaryName string = "legendary"
	isBabyName      string = "baby"
	isMythicalName  string = "mythical"
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
	graphQLClient  graphql.PokeGraphQLClient
}

func newCharacteristicManager(
	pokeApiClient api.PokeClient,
	serebiiScraper scraper.SerebiiScraper,
	graphQLClient graphql.PokeGraphQLClient,
) *characteristicManager {
	return &characteristicManager{
		pokeApiClient:  pokeApiClient,
		serebiiScraper: serebiiScraper,
		graphQLClient:  graphQLClient,
	}
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
	case isLegendaryName:
		return newIsLegendaryCharacteristic(cm.graphQLClient), nil
	case isBabyName:
		return newIsBabyCharacteristic(cm.graphQLClient), nil
	case isMythicalName:
		return newIsMythicalCharacteristic(cm.graphQLClient), nil
	}

	return nil, newInvalidCharacteristicsError(name)
}
