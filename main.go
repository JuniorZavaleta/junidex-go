package main

import (
	"./repositories"
	"./routers"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	repositories.InitDatabase()

	e.GET("/pokemon", routers.GetFetchPokemon)
	e.GET("/pokemon/filter", routers.GetFilterPokemon)
	e.POST("/pokemon", routers.CreatePokemon)

	e.Logger.Fatal(e.Start(":1323"))
}
