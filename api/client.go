package api

type PokeClient interface {
	FetchPokemons(characteristic, value string) ([]byte, error)
}
