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

The following characteristics will be implemented in the future:
1. Baby
2. Mythical
3. Legendary
4. Mono type
5. Dual type
6. Evolved by (item, stone, friendship, etc)
7. Evolution position
8. Evolution branched
9. Hisui
10. First partner
11. Fossil
12. Paradox