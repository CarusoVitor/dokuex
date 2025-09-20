package characteristics

import "github.com/CarusoVitor/dokuex/scraper"

type scraperCharacteristic struct {
	name           string
	serebiiScraper scraper.SerebiiScraper
}

func (sc scraperCharacteristic) getPokemons(value string) (PokemonSet, error) {
	pokemons, err := sc.serebiiScraper.ScrapPokemons(sc.name)
	if err != nil {
		return nil, err
	}

	set := make(PokemonSet, len(pokemons))
	for _, poke := range pokemons {
		set[poke] = struct{}{}
	}
	return set, nil
}

func newMegaCharacteristic(serebiiScraper scraper.SerebiiScraper) scraperCharacteristic {
	return scraperCharacteristic{
		name:           megaName,
		serebiiScraper: serebiiScraper,
	}
}
