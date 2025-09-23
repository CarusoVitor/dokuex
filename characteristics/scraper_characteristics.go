package characteristics

import (
	"fmt"
	"strings"

	"github.com/CarusoVitor/dokuex/scraper"
)

type scraperCharacteristic struct {
	name           string
	serebiiScraper scraper.SerebiiScraper
	formatter      func([]string) ([]string, error)
}

func (sc scraperCharacteristic) getPokemons(value string) (PokemonSet, error) {
	pokemons, err := sc.serebiiScraper.ScrapPokemons(sc.name)
	if err != nil {
		return nil, err
	}

	formattedPokemons, err := sc.formatter(pokemons)
	if err != nil {
		return nil, err
	}

	set := make(PokemonSet, len(formattedPokemons))
	for _, poke := range formattedPokemons {
		set[poke] = struct{}{}
	}
	return set, nil
}

func newMegaCharacteristic(serebiiScraper scraper.SerebiiScraper) scraperCharacteristic {
	return scraperCharacteristic{
		name:           megaName,
		serebiiScraper: serebiiScraper,
		formatter:      formatMega,
	}
}
func newGmaxCharacteristic(serebiiScraper scraper.SerebiiScraper) scraperCharacteristic {
	return scraperCharacteristic{
		name:           gmaxName,
		serebiiScraper: serebiiScraper,
		formatter:      formatGmax,
	}
}

// formatMega format pokemon mega names to be in the standard lowercase
// hyphen separated form with the pokemon name coming first
// There are two options:
// 1. Two word mega names e.g Mega Lucario
// 2. Three word mega names e.g Mega Charizard X
// Base name must also be returned since some characteristics only use them
func formatMega(pokemons []string) ([]string, error) {
	formatted := make([]string, 0, len(pokemons)*2)

	for idx := range pokemons {
		parts := strings.Split(strings.ToLower(pokemons[idx]), " ")
		if len(parts) != 2 && len(parts) != 3 {
			return nil, fmt.Errorf("mega pokemon name is not in an expected format: %s", pokemons[idx])
		}
		var sb strings.Builder

		name := parts[1]
		sb.WriteString(fmt.Sprintf("%s-mega", name))
		if len(parts) == 3 {
			sb.WriteString(fmt.Sprintf("-%s", parts[2]))
		}
		formatted = append(formatted, sb.String())
		formatted = append(formatted, name)
	}
	return formatted, nil
}

// formatGmax formats pokemon gmax (Gigantamax) names to be in the standard lowercase
// hyphen separated form with the pokemon name coming first
// Base name must also be returned since some characteristics only use them
func formatGmax(pokemons []string) ([]string, error) {
	formatted := make([]string, 0, len(pokemons)*2)

	for idx := range pokemons {
		name := strings.ToLower(pokemons[idx])
		formatted = append(formatted, fmt.Sprintf("%s-gmax", name))
		formatted = append(formatted, name)
	}
	return formatted, nil
}
