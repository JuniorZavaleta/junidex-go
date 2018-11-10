package repo

import (
	"../entities"
	"../entities/base"
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
)

func CreatePokemonFromRow(rows *sql.Rows) (*entities.Pokemon, error) {
	var Pokemon entities.Pokemon

	var PokemonId int
	var PokemonName string
	var PokemonTypeOneId int
	var PokemonTypeTwoId sql.NullInt64
	var PokemonTypeOne string
	var PokemonTypeTwo sql.NullString
	var PokemonHasPreEvol bool

	err := rows.Scan(
		&PokemonId,
		&PokemonName,
		&PokemonTypeOneId,
		&PokemonTypeTwoId,
		&PokemonHasPreEvol,
		&PokemonTypeOne,
		&PokemonTypeTwo,
	)

	if err == nil {
		Pokemon.Id = PokemonId
		Pokemon.Name = PokemonName
		Pokemon.TypeOneId = PokemonTypeOneId
		Pokemon.TypeOne = &entities.PokemonType{}
		Pokemon.TypeOne.Id = int(Pokemon.TypeOneId)
		Pokemon.TypeOne.Name = PokemonTypeOne
		Pokemon.HasPreEvol = PokemonHasPreEvol

		if PokemonTypeTwoId.Valid {
			Pokemon.TypeTwo = &entities.PokemonType{}
			Pokemon.TypeTwoId = base.IntNull{Value: int(PokemonTypeTwoId.Int64), Null: false}
			Pokemon.TypeTwo.Id = Pokemon.TypeTwoId.Value
			Pokemon.TypeTwo.Name = PokemonTypeTwo.String
		} else {
			Pokemon.TypeTwoId = base.IntNull{Value: 0, Null: true}
		}
	} else {
		panic(err)
	}

	return &Pokemon, nil
}

func AllPokemon() []entities.Pokemon {
	rows, _ := db.Query("CALL junidex.get_all_Pokemon()")

	var AllPokemon []entities.Pokemon

	for rows.Next() {
		Pokemon, _ := CreatePokemonFromRow(rows)
		AllPokemon = append(AllPokemon, *Pokemon)
	}

	return AllPokemon
}

func FilterPokemon(TypeOne string, TypeTwo string) ([]entities.Pokemon, error) {
	rows, _ := db.Query(
		"CALL junidex.filter_Pokemon(?, ?)",
		TypeOne,
		TypeTwo,
	)

	var FilteredPokemon []entities.Pokemon

	for rows.Next() {
		Pokemon, _ := CreatePokemonFromRow(rows)
		FilteredPokemon = append(FilteredPokemon, *Pokemon)
	}

	return FilteredPokemon, nil
}

func CreatePokemon(pokemon entities.Pokemon) (entities.Pokemon, error) {
	pokemonJson, err := json.Marshal(pokemon)

	if err != nil {
		panic(err)
	}

	rows, err := db.Query("CALL junidex.create_pokemon(?)", pokemonJson)

	if err != nil {
		panic(err)
	}

	if rows.Next() {
		pokemon, _ := CreatePokemonFromRow(rows)

		return *pokemon, nil
	} else {
		return entities.Pokemon{}, rows.Err()
	}
}
