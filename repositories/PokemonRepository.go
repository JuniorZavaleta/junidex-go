package repositories

import (
	"../entities"
	"database/sql"
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

	err := rows.Scan(
		&PokemonId,
		&PokemonName,
		&PokemonTypeOneId,
		&PokemonTypeTwoId,
		&PokemonTypeOne,
		&PokemonTypeTwo,
	)

	if err == nil {
		Pokemon.Id = PokemonId
		Pokemon.Name = PokemonName
		Pokemon.TypeOneId = PokemonTypeOneId
		Pokemon.TypeOne.Id = int(Pokemon.TypeOneId)
		Pokemon.TypeOne.Name = PokemonTypeOne

		if PokemonTypeTwoId.Valid {
			Pokemon.TypeTwoId = int(PokemonTypeTwoId.Int64)
			Pokemon.TypeTwo.Id = Pokemon.TypeTwoId
			Pokemon.TypeTwo.Name = PokemonTypeTwo.String
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

func CreatePokemon(pokemon entities.Pokemon) (*entities.Pokemon){
	// TODO
	//db.Query("CALL junidex.")

	return nil;
}