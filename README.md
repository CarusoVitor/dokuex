# Dokuex 🔢
Dokuex is a Pokémon characteristic matching tool. 

Dokuex queries all Pokémon that match given characteristics, e.g type, generation, moves, etc using data from the [PokeAPI](https://github.com/PokeAPI/pokeapi) project and from the [Serebii](http://serebii.net/) website. 

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
7. Gigantamax
8. Baby
9. Mythical
10. Legendary

The following characteristics will be implemented in the future:
1. Mono type
2. Dual type
3. Evolved by (item, stone, friendship, etc)
4. Evolution position
5. Evolution branched
6. Hisui
7. First partner
8. Fossil
9. Paradox