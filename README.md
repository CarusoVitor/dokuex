# Dokuex 🔢
Dokuex is a Pokemon characteristic matching tool. 

Dokuex queries all pokemons that match given characteristics, e.g type, generation, moves, etc using data from the [PokeAPI](https://github.com/PokeAPI/pokeapi) project. 

It was inspired by the [pokedoku](https://pokedoku.com/) website.

## Installation
```
git clone https://github.com/CarusoVitor/dokuex
cd dokuex
go build -o dokuex .
```
## Usage
Dokuex match allows matching as many characteristics as needed, following this template:
```
./dokuex match --[characteristic] value 
```

## Characteristics
Currently, the following characteristics are supported by dokuex:
1. Type
2. Generation
3. Move
4. Ability
5. Ultra Beast
6. Mega

The following characteristics will be implemented in the future:
1. Baby
2. Mythical
3. Legendary
4. Gigantamax
5. Mono type
6. Dual type
7. Evolved by (item, stone, friendship, etc)
8. Evolution position
9. Evolution branched
10. Hisui
11. First partner
12. Fossil
13. Paradox