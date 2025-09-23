package scraper

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gocolly/colly/v2"
)

const megaHtml = `
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

const gmaxHtml = `
<table class="tab" align="center">
   	<tbody>
      <tr>
         <td align="center" class="fooevo" width="60">
            No.
         </td>
         <td align="center" class="fooevo" width="40">
            Pic	
         </td>
         <td align="center" class="fooevo">
            Name	
         </td>
         <td align="center" class="fooevo">
            Type
         </td>
         <td align="center" class="fooevo">
            Abilities
         </td>
         <td align="center" class="fooevo">
            Location
         </td>
      </tr>
      <tr>
         <td align="center" class="fooinfo">
            #003
         </td>
         <td align="center" class="fooinfo">
            <table>
               <tbody>
                  <tr>
                     <td class="pkmn"><a href="/pokedex-swsh/venusaur"><img src="/swordshield/pokemon/003-gi.png" loading="lazy" class="listsprite"></a></td>
                  </tr>
               </tbody>
            </table>
         </td>
         <td align="center" class="fooinfo">
            <a href="/pokedex-swsh/venusaur">Venusaur<br>フシギバナ</a>
         </td>
         <td align="center" class="fooinfo"><a href="/pokedex-swsh/grass.shtml"><img src="/pokedex-bw/type/grass.gif" border="0"></a> <a href="/pokedex-swsh/poison.shtml"><img src="/pokedex-bw/type/poison.gif" border="0"></a></td>
         <td align="center" class="fooinfo">
            <a href="/abilitydex/overgrow.shtml">Overgrow</a> <br><a href="/abilitydex/chlorophyll.shtml">Chlorophyll</a>
         </td>
         <td align="center" class="fooinfo"><a href="/pokearth/galar/forestoffocus.shtml">Forest of Focus</a><br><a href="/pokearth/galar/traininglowlands.shtml">Training Lowlands</a><br><i><i>Isle of Armor Only</i></i><br></td>
      </tr>
      <tr>
         <td align="center" class="fooinfo">
            #006
         </td>
         <td align="center" class="fooinfo">
            <table>
               <tbody>
                  <tr>
                     <td class="pkmn"><a href="/pokedex-swsh/charizard"><img src="/swordshield/pokemon/006-gi.png" loading="lazy" class="listsprite"></a></td>
                  </tr>
               </tbody>
            </table>
         </td>
         <td align="center" class="fooinfo">
            <a href="/pokedex-swsh/charizard">Charizard<br>リザードン</a>
         </td>
         <td align="center" class="fooinfo"><a href="/pokedex-swsh/fire.shtml"><img src="/pokedex-bw/type/fire.gif" border="0"></a> <a href="/pokedex-swsh/flying.shtml"><img src="/pokedex-bw/type/flying.gif" border="0"></a></td>
         <td align="center" class="fooinfo">
            <a href="/abilitydex/blaze.shtml">Blaze</a> <br><a href="/abilitydex/solarpower.shtml">Solar Power</a>
         </td>
         <td align="center" class="fooinfo"><a href="/pokearth/galar/lakeofoutrage.shtml">Lake of Outrage</a><br></td>
      </tr>
      <tr>
         <td align="center" class="fooinfo">
            #009
         </td>
         <td align="center" class="fooinfo">
            <table>
               <tbody>
                  <tr>
                     <td class="pkmn"><a href="/pokedex-swsh/blastoise"><img src="/swordshield/pokemon/009-gi.png" loading="lazy" class="listsprite"></a></td>
                  </tr>
               </tbody>
            </table>
         </td>
         <td align="center" class="fooinfo">
            <a href="/pokedex-swsh/blastoise">Blastoise<br>カメックス</a>
         </td>
         <td align="center" class="fooinfo"><a href="/pokedex-swsh/water.shtml"><img src="/pokedex-bw/type/water.gif" border="0"></a></td>
         <td align="center" class="fooinfo">
            <a href="/abilitydex/torrent.shtml">Torrent</a> <br><a href="/abilitydex/raindish.shtml">Rain Dish</a>
         </td>
         <td align="center" class="fooinfo"><a href="/pokearth/galar/workoutsea.shtml">Workout Sea</a><br><a href="/pokearth/galar/stepping-stonesea.shtml">Stepping-Stone Sea</a><br><i><i>Isle of Armor Only</i></i><br></td>
      </tr>
   </tbody>
</table>
`

func Test_serebiiScraper_mega(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/pokemon/megaevolution.shtml" {
			fmt.Fprint(w, megaHtml)
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

func Test_serebiiScraper_gmax(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/swordshield/gigantamax.shtml" {
			fmt.Fprint(w, gmaxHtml)
		} else {
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	scraper := serebiiScraper{baseUrl: server.URL}
	c := colly.NewCollector()
	gmax, err := scraper.gmax(c)

	if err != nil {
		t.Errorf("gmax() error = %v, expected nil", err)
	}

	expected := []string{
		"Venusaur",
		"Charizard",
		"Blastoise",
	}

	if !reflect.DeepEqual(gmax, expected) {
		t.Errorf("expected %v, got %v", expected, gmax)
	}
}
