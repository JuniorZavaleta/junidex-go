package routers

import (
	"../entities"
	"../repo"
	"encoding/json"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func GetFetchPokemon(c echo.Context) error {
	pokemonList := repo.AllPokemon()

	return c.JSON(http.StatusOK, pokemonList)
}

func GetFilterPokemon(c echo.Context) error {
	typeOne := c.QueryParam("typeOne")
	typeTwo := c.QueryParam("typeTwo")
	pokemonList, _ := repo.FilterPokemon(typeOne, typeTwo)

	return c.JSON(http.StatusOK, pokemonList)
}

func CreatePokemon(c echo.Context) error {
	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&json_map)
	if err != nil {
		return err
	}

	TypeOne, _ := strconv.Atoi(json_map["type_one_id"].(string))
	var TypeTwo entities.IntNull

	if json_map["type_two_id"] != nil {
		aux, _ := strconv.Atoi(json_map["type_two_id"].(string))
		TypeTwo = entities.IntNull{Value: aux, Null: false}
	} else {
		TypeTwo = entities.IntNull{Value:0, Null: true}
	}

	pokemon := entities.Pokemon {
		Name: json_map["name"].(string),
		TypeOneId: TypeOne,
		TypeTwoId: TypeTwo,
		HasPreEvol: json_map["has_preevolution"].(bool),
	}

	pokemon, err = repo.CreatePokemon(pokemon)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusCreated, pokemon)
}