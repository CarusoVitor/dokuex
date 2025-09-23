package scraper

import (
	"errors"
	"fmt"
	"log/slog"
	"strings"

	"github.com/gocolly/colly/v2"
)

const serebiiUrl = "https://www.serebii.net"
const megaPage = "pokemon/megaevolution"
const gmaxPage = "swordshield/gigantamax"

type SerebiiScraper interface {
	ScrapPokemons(characteristic string) ([]string, error)
}

type serebiiScraper struct {
	baseUrl     string
	callbackErr error
}

func NewSerebiiScraper() *serebiiScraper {
	return &serebiiScraper{
		baseUrl: serebiiUrl,
	}
}

type UnexpectedHtmlError struct {
	message string
}

func (uh UnexpectedHtmlError) Error() string {
	return fmt.Sprintf("html parsed is not as expected: %s", uh.message)
}

var ErrForbidden error = errors.New("unable to access page, probably behind proxy")

func (bs *serebiiScraper) setupCollector() *colly.Collector {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		slog.Info("called", "url", r.URL.String())
	})

	c.OnError(func(r *colly.Response, err error) {
		bs.callbackErr = err
		if err.Error() == "Forbidden" {
			bs.callbackErr = ErrForbidden
		}
	})
	return c
}

func (bs serebiiScraper) formatUrl(page string) string {
	return fmt.Sprintf("%s/%s.shtml", bs.baseUrl, page)
}

// mega scraps Mega Evolution page to obtain all mega pokémons.
func (bs serebiiScraper) mega(c *colly.Collector) []string {
	megas := make([]string, 0)

	c.OnHTML("table.trainer", func(e *colly.HTMLElement) {
		// skip other tables with same class
		if e.ChildText("tr:nth-child(1)") != "Mega Evolved Pokémon" {
			return
		}
		e.ForEach("td", func(_ int, td *colly.HTMLElement) {
			pokemon := td.ChildText("table tr:nth-child(2)")
			if len(pokemon) == 0 {
				return
			}
			megas = append(megas, pokemon)
		})
	})

	c.Visit(bs.formatUrl(megaPage))
	return megas
}

// gmax scraps Gigantamax page to obtain all gigantamax pokémons.
// The table has the following header:
// | No | Pic | Name | Type | Abilities | Location |
// The pokemon's name element tag also contains the japanese name, which
// comes after a <br>, e.g "Venusaur<br/>フシギバナ"
func (bs serebiiScraper) gmax(c *colly.Collector) ([]string, error) {
	gmaxs := make([]string, 0)
	var err error
	c.OnHTML("table.tab", func(e *colly.HTMLElement) {
		// # explicitly use table.tab to not query inner tables
		e.ForEach("table.tab > tbody > tr", func(i int, tr *colly.HTMLElement) {
			// skip header
			if i == 0 {
				return
			}
			a, htmlErr := tr.DOM.Find("td:nth-child(3) a").Html()
			if htmlErr != nil || len(a) == 0 {
				err = UnexpectedHtmlError{message: fmt.Sprintf("failed to get pokémon name HTML: %v", htmlErr)}
				return
			}
			name := strings.Split(a, "<br/>")[0]
			if len(name) == 0 {
				err = UnexpectedHtmlError{message: "empty pokémon name"}
				return
			}
			gmaxs = append(gmaxs, name)
		})

	})
	c.Visit(bs.formatUrl(gmaxPage))
	return gmaxs, err
}

func (bs *serebiiScraper) ScrapPokemons(characteristic string) ([]string, error) {
	var pokemons []string
	var err error

	collector := bs.setupCollector()
	switch characteristic {
	case "mega":
		pokemons = bs.mega(collector)
	case "gmax":
		pokemons, err = bs.gmax(collector)
	default:
		return nil, fmt.Errorf("characteristic %s was not implemented", characteristic)
	}
	slog.Debug("Total pokemons found scrapping", "num", len(pokemons))

	if bs.callbackErr != nil {
		return nil, bs.callbackErr
	}

	if err != nil {
		return nil, err
	}

	if len(pokemons) == 0 {
		return nil, UnexpectedHtmlError{"scrape resulted in an empty list (internal error)"}
	}

	return pokemons, nil
}
