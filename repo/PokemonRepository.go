package repo

import (
	"../entities"
	"../entities/base"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func GetPokemonTypeName(TypeId int) (string) {
	Type, _ := client.Get(fmt.Sprintf("PT%d", TypeId)).Result()

	return Type
}

func CreatePokemonFromRow(rows *sql.Rows) (*entities.Pokemon, error) {
	var Pokemon entities.Pokemon

	var PokemonId int
	var PokemonName string
	var PokemonTypeOneId int
	var PokemonTypeTwoId sql.NullInt64
	var PokemonHasPreEvol bool

	err := rows.Scan(
		&PokemonId,
		&PokemonName,
		&PokemonTypeOneId,
		&PokemonTypeTwoId,
		&PokemonHasPreEvol,
	)

	if err == nil {
		Pokemon.Id = PokemonId
		Pokemon.Name = PokemonName
		Pokemon.TypeOneId = PokemonTypeOneId
		Pokemon.TypeOne = &entities.PokemonType{}
		Pokemon.TypeOne.Id = int(Pokemon.TypeOneId)
		Pokemon.HasPreEvol = PokemonHasPreEvol
		Pokemon.TypeOne.Name = GetPokemonTypeName(PokemonTypeOneId)

		if PokemonTypeTwoId.Valid {
			Pokemon.TypeTwo = &entities.PokemonType{}
			Pokemon.TypeTwoId = base.IntNull{Value: int(PokemonTypeTwoId.Int64), Null: false}
			Pokemon.TypeTwo.Id = Pokemon.TypeTwoId.Value
			Pokemon.TypeTwo.Name = GetPokemonTypeName(int(PokemonTypeTwoId.Int64))
		} else {
			Pokemon.TypeTwoId = base.IntNull{Value: 0, Null: true}
		}
	} else {
		panic(err)
	}

	return &Pokemon, nil
}

func AllPokemon() []entities.Pokemon {
	rows, _ := db.Query("CALL junidex.get_all_pokemon()")

	var AllPokemon []entities.Pokemon

	for rows.Next() {
		Pokemon, _ := CreatePokemonFromRow(rows)
		AllPokemon = append(AllPokemon, *Pokemon)
	}

	return AllPokemon
}

func FilterPokemon(TypeOne string, TypeTwo string) ([]entities.Pokemon, error) {
	rows, _ := db.Query("CALL junidex.filter_pokemon(?, ?)",
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

func FindPokemon(Id int) (*entities.Pokemon, error) {
	rows, err := db.Query("CALL junidex.find_pokemon(?)", Id)

	if err != nil {
		panic(err)
	} else {
		var Pokemon *entities.Pokemon

		if rows.Next() {
			Pokemon, _ = CreatePokemonFromRow(rows)

			return Pokemon, nil
		}
	}

	return &entities.Pokemon{}, rows.Err()
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
