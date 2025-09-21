package scraper

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gocolly/colly/v2"
)

func Test_serebiiScraper_mega(t *testing.T) {
	html := `
	<table class="trainer" align="center"><tbody><tr><td class="interact" colspan="9">Mega Evolved Pokémon</td></tr>
	<tr>
	<td valign="top"><table><tbody><tr><td><a href="/pokemon/venusaur/"><img src="/pokemonhome/pokemon/small/003-m.png" class="pkmn"></a></td></tr>
	<tr><td align="center"><a href="/pokemon/venusaur/">Mega Venusaur</a></td></tr><tr><td align="center"><a href="/pokemon/type/grass"><img src="/pokedex-bw/type/grass.gif" border="0"></a> <a href="/pokemon/type/poison"><img src="/pokedex-bw/type/poison.gif" border="0"></a></td></tr></tbody></table></td>
	<td valign="top"><table><tbody><tr><td><a href="/pokemon/charizard/"><img src="/pokemonhome/pokemon/small/006-mx.png" class="pkmn"></a></td></tr>
	<tr><td align="center"><a href="/pokemon/charizard/">Mega Charizard X</a></td></tr><tr><td align="center"><a href="/pokemon/type/fire"><img src="/pokedex-bw/type/fire.gif" border="0"></a> <a href="/pokemon/type/dragon"><img src="/pokedex-bw/type/dragon.gif" border="0"></a></td></tr></tbody></table></td>
	<td valign="top"><table><tbody><tr><td><a href="/pokemon/charizard/"><img src="/pokemonhome/pokemon/small/006-my.png" class="pkmn"></a></td></tr>
	<tr><td align="center"><a href="/pokemon/charizard/">Mega Charizard Y</a></td></tr><tr><td align="center"><a href="/pokemon/type/fire"><img src="/pokedex-bw/type/fire.gif" border="0"></a> <a href="/pokemon/type/flying"><img src="/pokedex-bw/type/flying.gif" border="0"></a></td></tr></tbody></table></td>
	<td valign="top"><table><tbody><tr><td><a href="/pokemon/blastoise/"><img src="/pokemonhome/pokemon/small/009-m.png" class="pkmn"></a></td></tr>
	<tr><td align="center"><a href="/pokemon/blastoise/">Mega Blastoise</a></td></tr><tr><td align="center"><a href="/pokemon/type/water"><img src="/pokedex-bw/type/water.gif" border="0"></a></td></tr></tbody></table></td>
	<td valign="top"><table><tbody><tr><td><a href="/pokemon/beedrill/"><img src="/pokemonhome/pokemon/small/015-m.png" class="pkmn"></a></td></tr>
	<tr><td align="center"><a href="/pokemon/beedrill/">Mega Beedrill</a></td></tr><tr><td align="center"><a href="/pokemon/type/bug"><img src="/pokedex-bw/type/bug.gif" border="0"></a> <a href="/pokemon/type/poison"><img src="/pokedex-bw/type/poison.gif" border="0"></a></td></tr></tbody></table></td>
	<td valign="top"><table><tbody><tr><td><a href="/pokemon/pidgeot/"><img src="/pokemonhome/pokemon/small/018-m.png" class="pkmn"></a></td></tr>
	<tr><td align="center"><a href="/pokemon/pidgeot/">Mega Pidgeot</a></td></tr><tr><td align="center"><a href="/pokemon/type/normal"><img src="/pokedex-bw/type/normal.gif" border="0"></a> <a href="/pokemon/type/flying"><img src="/pokedex-bw/type/flying.gif" border="0"></a></td></tr></tbody></table></td>
	</tr><tr><td valign="top"><table><tbody><tr><td><a href="/pokemon/raichu/"><img src="/pokemon/art/026-mx.png" height="120" class="pkmn"></a></td></tr>
	<tr><td align="center"><a href="/pokemon/raichu/">Mega Raichu X</a></td></tr><tr><td align="center"><a href="/pokemon/type/electric"><img src="/pokedex-bw/type/electric.gif" border="0"></a></td></tr></tbody></table></td>
	<td valign="top"><table><tbody><tr><td><a href="/pokemon/raichu/"><img src="/pokemon/art/026-my.png" height="120" class="pkmn"></a></td></tr>
	<tr><td align="center"><a href="/pokemon/raichu/">Mega Raichu Y</a></td></tr><tr><td align="center"><a href="/pokemon/type/electric"><img src="/pokedex-bw/type/electric.gif" border="0"></a></td></tr></tbody></table></td>
	</tr></tbody></table>
	<table class="trainer" align="center"><tbody><tr><td class="interact" colspan="9">Primal Reversion Pokémon</td></tr>
	<tr>
	<td valign="top" align="center"><table><tbody><tr><td><a href="/pokemon/kyogre/"><img src="/pokemonhome/pokemon/small/382-p.png" class="pkmn"></a></td></tr>
	<tr><td align="center"><a href="/pokemon/kyogre/">Primal Kyogre</a></td></tr><tr><td align="center"><a href="/pokemon/type/water"><img src="/pokedex-bw/type/water.gif" border="0"></a></td></tr></tbody></table></td>
	<td valign="top" align="center"><table><tbody><tr><td><a href="/pokemon/groudon/"><img src="/pokemonhome/pokemon/small/383-p.png" class="pkmn"></a></td></tr>
	<tr><td align="center"><a href="/pokemon/groudon/">Primal Groudon</a></td></tr><tr><td align="center"><a href="/pokemon/type/ground"><img src="/pokedex-bw/type/ground.gif" border="0"></a> <a href="/pokemon/type/fire"><img src="/pokedex-bw/type/fire.gif" border="0"></a></td></tr></tbody></table></td>
	</tr>
	</tbody></table>
	`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/pokemon/megaevolution.shtml" {
			fmt.Fprint(w, html)
		} else {
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	scraper := serebiiScraper{baseUrl: server.URL}
	c := colly.NewCollector()
	megas := scraper.mega(c)

	expected := []string{
		"Mega Venusaur",
		"Mega Charizard X",
		"Mega Charizard Y",
		"Mega Blastoise",
		"Mega Beedrill",
		"Mega Pidgeot",
		"Mega Raichu X",
		"Mega Raichu Y",
	}
	if !reflect.DeepEqual(megas, expected) {
		t.Errorf("expected %v, got %v", expected, megas)
	}
}
func Test_serebiiScraper_ScrapPokemonsForbiddenPage(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Forbidden", http.StatusForbidden)
	}))
	defer server.Close()

	scraper := serebiiScraper{baseUrl: server.URL}
	c := scraper.setupCollector()
	_ = scraper.mega(c)
	if scraper.callbackErr != ErrForbidden {
		t.Errorf("serebiiScraper.ScrapPokemons() error = %v, wantErr %v", scraper.callbackErr, ErrForbidden)
	}
}
