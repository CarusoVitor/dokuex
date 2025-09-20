package characteristics

import "github.com/CarusoVitor/dokuex/scraper"

type scraperCharacteristic struct {
	name         string
	bulbaScraper scraper.BulbaScraper
}

func (sc scraperCharacteristic) getPokemons(value string) (PokemonSet, error) {
	pokemons, err := sc.bulbaScraper.ScrapPokemons(sc.name)
	if err != nil {
		return nil, err
	}

	set := make(PokemonSet, len(pokemons))
	for _, poke := range pokemons {
		set[poke] = struct{}{}
	}
	return set, nil
}

func newMegaCharacteristic(bulbaScraper scraper.BulbaScraper) scraperCharacteristic {
	return scraperCharacteristic{
		name:         megaName,
		bulbaScraper: bulbaScraper,
	}
}
