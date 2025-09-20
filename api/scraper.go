package api

import (
	"fmt"
	"os"
	"strings"

	"github.com/gocolly/colly/v2"
)

const bulbapediaUrl = "https://bulbapedia.bulbagarden.net/wiki/"

type bulbapediaScraper struct {
	cache   *cache
	baseUrl string
}

func NewBulbapediaScraper() *bulbapediaScraper {
	cache := newCache()
	return &bulbapediaScraper{
		cache:   cache,
		baseUrl: pokeapiUrl,
	}
}

func (bs bulbapediaScraper) mega() []byte {
	c := colly.NewCollector()

	megas := make([]byte, 0)

	c.OnHTML("table.roundy.sortable", func(e *colly.HTMLElement) {
		prev := e.DOM.Prev().Text()
		if prev != "Height and weight comparisons" {
			return
		}
		e.ForEach("tr", func(_ int, row *colly.HTMLElement) {
			value := row.ChildText("td:nth-child(1)")
			if len(value) == 0 {
				return
			}
			if strings.HasPrefix(value, "#") {
				// single row pokedex number
				mega := row.ChildText("td:nth-child(5)")
				megas = append(megas, mega)
				fmt.Printf("%s\n", mega)
			} else if strings.HasPrefix(value, "Mega") {
				megas = append(megas, value)
				fmt.Printf("%s\n", value)
			} else {
				fmt.Fprintf(os.Stderr, "invalid form %s", value)
			}

		})
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://bulbapedia.bulbagarden.net/wiki/Mega_Evolution")
}

func (bs *bulbapediaScraper) FetchPokemons(characteristic, value string) ([]byte, error) {
	if pokemons, ok := bs.cache.get(characteristic); ok {
		return pokemons, nil
	}

	var pokemons []byte
	switch characteristic {
	case "mega":
		pokemons = bs.mega()
	}
	bs.cache.add(characteristic, pokemons)
	return pokemons, nil
}
