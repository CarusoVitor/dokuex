package characteristics

import (
	"fmt"

	"github.com/CarusoVitor/dokuex/api"
	"github.com/CarusoVitor/dokuex/graphql"
	"github.com/CarusoVitor/dokuex/scraper"
)

const (
	TypeName       string = "type"
	GenerationName string = "generation"
	MoveName       string = "move"
	AbilityName    string = "ability"
	UltraBeastName string = "ultra-beast"
	MegaName       string = "mega"
	GmaxName       string = "gmax"
	LegendaryName  string = "legendary"
	BabyName       string = "baby"
	MythicalName   string = "mythical"
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
	case TypeName:
		return newTypeCharacteristic(cm.pokeApiClient), nil
	case GenerationName:
		return newGenerationCharacteristic(cm.pokeApiClient), nil
	case MoveName:
		return newMoveCharacteristic(cm.pokeApiClient), nil
	case AbilityName:
		return newAbilityCharacteristic(cm.pokeApiClient), nil
	case UltraBeastName:
		return newUltraBeastCharacteristic(cm.pokeApiClient), nil
	case MegaName:
		return newMegaCharacteristic(cm.serebiiScraper), nil
	case GmaxName:
		return newGmaxCharacteristic(cm.serebiiScraper), nil
	case LegendaryName:
		return newIsLegendaryCharacteristic(cm.graphQLClient), nil
	case BabyName:
		return newIsBabyCharacteristic(cm.graphQLClient), nil
	case MythicalName:
		return newIsMythicalCharacteristic(cm.graphQLClient), nil
	}

	return nil, newInvalidCharacteristicsError(name)
}
