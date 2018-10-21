package routers

import (
	"../repositories"
	"github.com/labstack/echo"
	"net/http"
)

func GetFetchPokemon(c echo.Context) error {
	pokemonList := repositories.AllPokemon()

	return c.JSON(http.StatusOK, pokemonList)
}

func GetFilterPokemon(c echo.Context) error {
	typeOne := c.QueryParam("typeOne")
	typeTwo := c.QueryParam("typeTwo")
	pokemonList, _ := repositories.FilterPokemon(typeOne, typeTwo)

	return c.JSON(http.StatusOK, pokemonList)
}
