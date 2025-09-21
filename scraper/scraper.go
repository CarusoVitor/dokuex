package scraper

import (
	"errors"
	"fmt"
	"log/slog"

	"github.com/gocolly/colly/v2"
)

const serebiiUrl = "https://www.serebii.net"
const megaPage = "pokemon/megaevolution"

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

var ErrForbidden error = errors.New("unable to acess page, probably behind proxy")

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

// mega scraps Serebii's Mega Evolution page to obtain all mega pokémons.
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

func (bs *serebiiScraper) ScrapPokemons(characteristic string) ([]string, error) {
	var pokemons []string

	collector := bs.setupCollector()
	switch characteristic {
	case "mega":
		pokemons = bs.mega(collector)
	default:
		return nil, fmt.Errorf("characteristic %s was not implemented", characteristic)
	}
	slog.Info("Total pokemons found scrapping", "num", len(pokemons))

	if bs.callbackErr != nil {
		return nil, bs.callbackErr
	}

	if len(pokemons) == 0 {
		return nil, UnexpectedHtmlError{"scrap resulted in an empty list (internal error)"}
	}

	return pokemons, nil
}
