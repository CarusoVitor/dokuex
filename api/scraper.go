package api

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/gocolly/colly/v2"
)

const bulbapediaUrl = "https://bulbapedia.bulbagarden.net/wiki/"
const megaTableHeader = "Height and weight comparisons"

type BulbaClient interface {
	ScrapPokemons(characteristic string) ([]string, error)
}

type bulbapediaScraper struct {
	baseUrl string
}

func NewBulbapediaScraper(collector colly.Collector) *bulbapediaScraper {
	return &bulbapediaScraper{
		baseUrl: bulbapediaUrl,
	}
}

type UnexpectedHtmlError struct {
	message string
}

func (uh UnexpectedHtmlError) Error() string {
	return fmt.Sprintf("html parsed is not as expected: %s", uh.message)
}

// mega scraps Bulbapedia's Mega Evolution page to obtain all mega pokemons.
// Since there are multiple tables with different comparisons, we query only
// one which have all mega pokemons.

// The html table has the following header:
// Dex (1) | Pokemon (2) | Heigth (before mega) (3) | Weigth (before mega) (4) | \
// Mega-Evolved Pokemon (5) | Heigth (after mega) (6) | Weigth (after mega) (7) | \
// Heigth (increased/decreased) (8) | Weigth (increased/decreased) (9)
func (bs bulbapediaScraper) mega() ([]string, error) {
	c := colly.NewCollector()

	megas := make([]string, 0)
	var err error

	c.OnHTML("table.roundy.sortable", func(e *colly.HTMLElement) {
		prev := e.DOM.Prev().Text()
		if prev != megaTableHeader {
			return
		}
		e.ForEach("tr", func(_ int, row *colly.HTMLElement) {
			value := row.ChildText("td:nth-child(1)")
			if len(value) == 0 {
				return
			}
			if strings.HasPrefix(value, "#") {
				mega := row.ChildText("td:nth-child(5)")
				if len(mega) == 0 {
					err = UnexpectedHtmlError{"mega pokemon name is empty"}
				}
				megas = append(megas, mega)
			} else {
				err = UnexpectedHtmlError{"mega first element is not the dex number"}
			}

		})
	})
	c.OnRequest(func(r *colly.Request) {
		slog.Debug("Visiting", "url", r.URL.String())
	})

	c.Visit(fmt.Sprintf("%s/Mega_Evolution", bs.baseUrl))

	return megas, err
}

func (bs *bulbapediaScraper) ScrapPokemons(characteristic string) ([]string, error) {
	var pokemons []string
	var err error = nil

	switch characteristic {
	case "mega":
		pokemons, err = bs.mega()
	default:
		return nil, fmt.Errorf("characteristic %s was not implemented", characteristic)
	}

	return pokemons, err
}
