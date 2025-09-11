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

The following characteristics will be implemented in the future:
1. Baby
2. Mythical
3. Legendary
4. Mega
5. Gigantamax
6. Mono type
7. Dual type
8. Evolved by (item, stone, friendship, etc)
9. Evolution position
10. Evolution branched
11. Hisui
12. First partner
13. Fossil
14. Paradox