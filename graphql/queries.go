package graphql

const queryIsLegendary string = `
  {
  pokemon_v2_pokemonspecies(
    where: {is_legendary: {_eq: true}}
  ) {
    name
  }
}
`

const queryIsBaby string = `
  {
  pokemon_v2_pokemonspecies(
    where: {is_baby: {_eq: true}}
  ) {
    name
  }
}
`

const queryIsMythical string = `
  {
  pokemon_v2_pokemonspecies(
    where: {is_mythical: {_eq: true}}
  ) {
    name
  }
}
`
